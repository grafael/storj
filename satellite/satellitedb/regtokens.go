// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/satellite/console"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// ensures that registrationTokens implements console.RegistrationTokens.
var _ console.RegistrationTokens = (*registrationTokens)(nil)

// registrationTokens is an implementation of RegistrationTokens interface using spacemonkeygo/dbx orm
type registrationTokens struct {
	db dbx.Methods
}

// Create creates new registration token
func (rt *registrationTokens) Create(ctx context.Context, projectLimit int) (_ *console.RegistrationToken, err error) {
	defer mon.Task()(&ctx)(&err)
	secret, err := console.NewRegistrationSecret()
	if err != nil {
		return nil, err
	}

	regToken, err := rt.db.Create_RegistrationToken(
		ctx,
		dbx.RegistrationToken_Secret(secret[:]),
		dbx.RegistrationToken_ProjectLimit(projectLimit),
		dbx.RegistrationToken_Create_Fields{},
	)
	if err != nil {
		return nil, err
	}

	return registrationTokenFromDBX(ctx, regToken)
}

// GetBySecret retrieves RegTokenInfo with given Secret
func (rt *registrationTokens) GetBySecret(ctx context.Context, secret console.RegistrationSecret) (_ *console.RegistrationToken, err error) {
	defer mon.Task()(&ctx)(&err)
	regToken, err := rt.db.Get_RegistrationToken_By_Secret(ctx, dbx.RegistrationToken_Secret(secret[:]))
	if err != nil {
		return nil, err
	}

	return registrationTokenFromDBX(ctx, regToken)
}

// GetByOwnerID retrieves RegTokenInfo by ownerID
func (rt *registrationTokens) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) (_ *console.RegistrationToken, err error) {
	defer mon.Task()(&ctx)(&err)
	regToken, err := rt.db.Get_RegistrationToken_By_OwnerId(ctx, dbx.RegistrationToken_OwnerId(ownerID[:]))
	if err != nil {
		return nil, err
	}

	return registrationTokenFromDBX(ctx, regToken)
}

// UpdateOwner updates registration token's owner
func (rt *registrationTokens) UpdateOwner(ctx context.Context, secret console.RegistrationSecret, ownerID uuid.UUID) (err error) {
	defer mon.Task()(&ctx)(&err)
	_, err = rt.db.Update_RegistrationToken_By_Secret(
		ctx,
		dbx.RegistrationToken_Secret(secret[:]),
		dbx.RegistrationToken_Update_Fields{
			OwnerId: dbx.RegistrationToken_OwnerId(ownerID[:]),
		},
	)

	return err
}

// registrationTokenFromDBX is used for creating RegistrationToken entity from autogenerated dbx.RegistrationToken struct
func registrationTokenFromDBX(ctx context.Context, regToken *dbx.RegistrationToken) (_ *console.RegistrationToken, err error) {
	if regToken == nil {
		return nil, errs.New("token parameter is nil")
	}

	var secret [32]byte

	copy(secret[:], regToken.Secret)

	result := &console.RegistrationToken{
		Secret:       secret,
		OwnerID:      nil,
		ProjectLimit: regToken.ProjectLimit,
		CreatedAt:    regToken.CreatedAt,
	}

	if regToken.OwnerId != nil {
		ownerID, err := bytesToUUID(regToken.OwnerId)
		if err != nil {
			return nil, err
		}

		result.OwnerID = &ownerID
	}

	return result, nil
}
