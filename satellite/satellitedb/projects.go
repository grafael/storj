// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"time"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/satellite/console"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// ensures that projects implements console.Projects.
var _ console.Projects = (*projects)(nil)

// implementation of Projects interface repository using spacemonkeygo/dbx orm
type projects struct {
	db dbx.Methods
}

// GetAll is a method for querying all projects from the database.
func (projects *projects) GetAll(ctx context.Context) (_ []console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	projectsDbx, err := projects.db.All_Project(ctx)
	if err != nil {
		return nil, err
	}

	return projectsFromDbxSlice(ctx, projectsDbx)
}

// GetOwn is a method for querying all projects created by current user from the database.
func (projects *projects) GetOwn(ctx context.Context, userID uuid.UUID) (_ []console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	projectsDbx, err := projects.db.All_Project_By_OwnerId_OrderBy_Asc_CreatedAt(ctx, dbx.Project_OwnerId(userID[:]))
	if err != nil {
		return nil, err
	}

	return projectsFromDbxSlice(ctx, projectsDbx)
}

// GetCreatedBefore retrieves all projects created before provided date
func (projects *projects) GetCreatedBefore(ctx context.Context, before time.Time) (_ []console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	projectsDbx, err := projects.db.All_Project_By_CreatedAt_Less_OrderBy_Asc_CreatedAt(ctx, dbx.Project_CreatedAt(before))
	if err != nil {
		return nil, err
	}

	return projectsFromDbxSlice(ctx, projectsDbx)
}

// GetByUserID is a method for querying all projects from the database by userID.
func (projects *projects) GetByUserID(ctx context.Context, userID uuid.UUID) (_ []console.Project, err error) {
	defer mon.Task()(&ctx)(&err)
	projectsDbx, err := projects.db.All_Project_By_ProjectMember_MemberId_OrderBy_Asc_Project_Name(ctx, dbx.ProjectMember_MemberId(userID[:]))
	if err != nil {
		return nil, err
	}

	return projectsFromDbxSlice(ctx, projectsDbx)
}

// Get is a method for querying project from the database by id.
func (projects *projects) Get(ctx context.Context, id uuid.UUID) (_ *console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	project, err := projects.db.Get_Project_By_Id(ctx, dbx.Project_Id(id[:]))
	if err != nil {
		return nil, err
	}

	return projectFromDBX(ctx, project)
}

// Insert is a method for inserting project into the database.
func (projects *projects) Insert(ctx context.Context, project *console.Project) (_ *console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	projectID, err := uuid.New()
	if err != nil {
		return nil, err
	}

	createFields := dbx.Project_Create_Fields{}
	if !project.PartnerID.IsZero() {
		createFields.PartnerId = dbx.Project_PartnerId(project.PartnerID[:])
	}

	createdProject, err := projects.db.Create_Project(ctx,
		dbx.Project_Id(projectID[:]),
		dbx.Project_Name(project.Name),
		dbx.Project_Description(project.Description),
		dbx.Project_UsageLimit(0),
		dbx.Project_OwnerId(project.OwnerID[:]),
		createFields,
	)

	if err != nil {
		return nil, err
	}

	return projectFromDBX(ctx, createdProject)
}

// Delete is a method for deleting project by Id from the database.
func (projects *projects) Delete(ctx context.Context, id uuid.UUID) (err error) {
	defer mon.Task()(&ctx)(&err)

	_, err = projects.db.Delete_Project_By_Id(ctx, dbx.Project_Id(id[:]))

	return err
}

// Update is a method for updating project entity
func (projects *projects) Update(ctx context.Context, project *console.Project) (err error) {
	defer mon.Task()(&ctx)(&err)

	updateFields := dbx.Project_Update_Fields{
		Description: dbx.Project_Description(project.Description),
		UsageLimit:  dbx.Project_UsageLimit(project.UsageLimit),
	}

	_, err = projects.db.Update_Project_By_Id(ctx,
		dbx.Project_Id(project.ID[:]),
		updateFields)

	return err
}

// List returns paginated projects, created before provided timestamp.
func (projects *projects) List(ctx context.Context, offset int64, limit int, before time.Time) (_ console.ProjectsPage, err error) {
	defer mon.Task()(&ctx)(&err)

	var page console.ProjectsPage

	dbxProjects, err := projects.db.Limited_Project_By_CreatedAt_Less_OrderBy_Asc_CreatedAt(ctx,
		dbx.Project_CreatedAt(before.UTC()),
		limit+1,
		offset,
	)
	if err != nil {
		return console.ProjectsPage{}, err
	}

	if len(dbxProjects) == limit+1 {
		page.Next = true
		page.NextOffset = offset + int64(limit) + 1

		dbxProjects = dbxProjects[:len(dbxProjects)-1]
	}

	projs, err := projectsFromDbxSlice(ctx, dbxProjects)
	if err != nil {
		return console.ProjectsPage{}, err
	}

	page.Projects = projs
	return page, nil
}

// projectFromDBX is used for creating Project entity from autogenerated dbx.Project struct
func projectFromDBX(ctx context.Context, project *dbx.Project) (_ *console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	if project == nil {
		return nil, errs.New("project parameter is nil")
	}

	id, err := bytesToUUID(project.Id)
	if err != nil {
		return nil, err
	}

	var partnerID uuid.UUID
	if len(project.PartnerId) > 0 {
		partnerID, err = bytesToUUID(project.PartnerId)
		if err != nil {
			return nil, err
		}
	}

	ownerID, err := bytesToUUID(project.OwnerId)
	if err != nil {
		return nil, err
	}

	return &console.Project{
		ID:          id,
		Name:        project.Name,
		Description: project.Description,
		PartnerID:   partnerID,
		OwnerID:     ownerID,
		CreatedAt:   project.CreatedAt,
	}, nil
}

// projectsFromDbxSlice is used for creating []Project entities from autogenerated []*dbx.Project struct
func projectsFromDbxSlice(ctx context.Context, projectsDbx []*dbx.Project) (_ []console.Project, err error) {
	defer mon.Task()(&ctx)(&err)

	var projects []console.Project
	var errors []error

	// Generating []dbo from []dbx and collecting all errors
	for _, projectDbx := range projectsDbx {
		project, err := projectFromDBX(ctx, projectDbx)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		projects = append(projects, *project)
	}

	return projects, errs.Combine(errors...)
}
