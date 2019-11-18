// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"encoding/base64"
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/process"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/metainfo"
)

const maxNumOfSegments = byte(64)

var (
	detectCmd = &cobra.Command{
		Use:   "detect",
		Short: "Detects zombie segments in DB",
		Args:  cobra.OnlyValidArgs,
		RunE:  cmdDetect,
	}

	detectCfg struct {
		DatabaseURL string `help:"the database connection string to use" default:"postgres://"`
		From        string `help:"begin of date range for detecting zombie segments" default:""`
		To          string `help:"end of date range for detecting zombie segments" default:""`
		File        string `help:"location of file with report" default:"detect_result.csv"`
	}
)

// Cluster key for objects map.
type Cluster struct {
	projectID string
	bucket    string
}

// Object represents object with segments.
type Object struct {
	// TODO verify if we have more than 64 segments for object in network
	segments uint64

	expectedNumberOfSegments byte

	hasLastSegment bool
	// if skip is true then segments from this object shouldn't be treated as zombie segments
	// and printed out, e.g. when one of segments is out of specified date rage
	skip bool
}

// ObjectsMap map that keeps objects representation.
type ObjectsMap map[Cluster]map[storj.Path]*Object

// Observer metainfo.Loop observer for zombie reaper.
type Observer struct {
	db      metainfo.PointerDB
	objects ObjectsMap
	writer  *csv.Writer

	from *time.Time
	to   *time.Time

	lastProjectID string

	inlineSegments     int
	lastInlineSegments int
	remoteSegments     int
}

// RemoteSegment processes a segment to collect data needed to detect zombie segment.
func (observer *Observer) RemoteSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return observer.processSegment(ctx, path, pointer)
}

// InlineSegment processes a segment to collect data needed to detect zombie segment.
func (observer *Observer) InlineSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return observer.processSegment(ctx, path, pointer)
}

// Object not used in this implementation.
func (observer *Observer) Object(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return nil
}

func (observer *Observer) processSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) error {
	if observer.lastProjectID != "" && observer.lastProjectID != path.ProjectIDString {
		err := analyzeProject(ctx, observer.db, observer.objects, observer.writer)
		if err != nil {
			return err
		}

		// cleanup map to free memory
		observer.objects = make(ObjectsMap)
	}

	cluster := Cluster{
		projectID: path.ProjectIDString,
		bucket:    path.BucketName,
	}
	object := findOrCreate(cluster, path.EncryptedObjectPath, observer.objects)
	if observer.from != nil && pointer.CreationDate.Before(*observer.from) {
		object.skip = true
		return nil
	} else if observer.to != nil && pointer.CreationDate.After(*observer.to) {
		object.skip = true
		return nil
	}

	isLastSegment := path.Segment == "l"
	if isLastSegment {
		object.hasLastSegment = true

		streamMeta := pb.StreamMeta{}
		err := proto.Unmarshal(pointer.Metadata, &streamMeta)
		if err != nil {
			return errs.New("unexpected error unmarshalling pointer metadata %s", err)
		}

		if streamMeta.NumberOfSegments > 0 {
			if streamMeta.NumberOfSegments > int64(maxNumOfSegments) {
				object.skip = true
				zap.S().Warn("unsupported number of segments", zap.Int64("index", streamMeta.NumberOfSegments))
				return nil
			}
			object.expectedNumberOfSegments = byte(streamMeta.NumberOfSegments)
		}
	} else {
		segmentIndex, err := strconv.Atoi(path.Segment[1:])
		if err != nil {
			return err
		}
		if segmentIndex >= int(maxNumOfSegments) {
			object.skip = true
			zap.S().Warn("unsupported segment index", zap.Int("index", segmentIndex))
			return nil
		}

		if object.segments&(1<<uint64(segmentIndex)) != 0 {
			encodedPath := storj.JoinPaths(path.ProjectIDString, path.Segment, path.BucketName, base64.StdEncoding.EncodeToString([]byte(path.EncryptedObjectPath)))
			return errs.New("fatal error this segment is duplicated: %s", encodedPath)
		}

		object.segments |= 1 << uint64(segmentIndex)
	}

	// collect number of pointers for report
	if pointer.Type == pb.Pointer_INLINE {
		observer.inlineSegments++
		if isLastSegment {
			observer.lastInlineSegments++
		}
	} else {
		observer.remoteSegments++
	}
	observer.lastProjectID = cluster.projectID
	return nil
}

func init() {
	rootCmd.AddCommand(detectCmd)

	defaults := cfgstruct.DefaultsFlag(rootCmd)
	process.Bind(detectCmd, &detectCfg, defaults)
}

func cmdDetect(cmd *cobra.Command, args []string) (err error) {
	ctx, _ := process.Ctx(cmd)

	log := zap.L()
	db, err := metainfo.NewStore(log.Named("pointerdb"), detectCfg.DatabaseURL)
	if err != nil {
		return errs.New("error connecting database: %+v", err)
	}
	defer func() {
		err = errs.Combine(err, db.Close())
	}()

	file, err := os.Create(detectCfg.File)
	if err != nil {
		return errs.New("error creating result file: %+v", err)
	}
	defer func() {
		err = errs.Combine(err, file.Close())
	}()

	writer := csv.NewWriter(file)
	defer func() {
		writer.Flush()
		err = errs.Combine(err, writer.Error())
	}()

	headers := []string{
		"ProjectID",
		"SegmentIndex",
		"Bucket",
		"EncodedEncryptedPath",
		"CreationDate",
	}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	observer := &Observer{
		objects: make(ObjectsMap),
		db:      db,
		writer:  writer,
	}

	if detectCfg.From != "" {
		fromDate, err := time.Parse(time.RFC3339, detectCfg.From)
		if err != nil {
			return err
		}
		observer.from = &fromDate
	}

	if detectCfg.To != "" {
		toDate, err := time.Parse(time.RFC3339, detectCfg.From)
		if err != nil {
			return err
		}
		observer.to = &toDate
	}

	err = metainfo.IterateDatabase(ctx, db, observer)
	if err != nil {
		return err
	}

	log.Info("number of inline segments", zap.Int("segments", observer.inlineSegments))
	log.Info("number of last inline segments", zap.Int("segments", observer.lastInlineSegments))
	log.Info("number of remote segments", zap.Int("segments", observer.remoteSegments))
	log.Info("number of all segments", zap.Int("segments", observer.remoteSegments+observer.inlineSegments))
	return nil
}

func analyzeProject(ctx context.Context, db metainfo.PointerDB, objectsMap ObjectsMap, csvWriter *csv.Writer) error {
	for cluster, objects := range objectsMap {
		for path, object := range objects {
			if object.skip {
				continue
			}

			segments := make([]byte, 0)
			for i := byte(0); i < maxNumOfSegments; i++ {
				found := object.segments&(1<<uint64(i)) != 0
				if found {
					segments = append(segments, i)
				}
			}

			if object.hasLastSegment {
				brokenObject := false
				if object.expectedNumberOfSegments == 0 {
					for i := 0; i < len(segments); i++ {
						if int(segments[i]) != i {
							brokenObject = true
							break
						}
					}
				} else if len(segments) != int(object.expectedNumberOfSegments)-1 {
					// expectedNumberOfSegments-1 because 'segments' doesn't contain last segment
					brokenObject = true
				}

				if !brokenObject {
					segments = []byte{}
				} else {
					err := printSegment(ctx, db, cluster, "l", path, csvWriter)
					if err != nil {
						return err
					}
				}
			}

			for _, segmentIndex := range segments {
				err := printSegment(ctx, db, cluster, "s"+strconv.Itoa(int(segmentIndex)), path, csvWriter)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func printSegment(ctx context.Context, db metainfo.PointerDB, cluster Cluster, segmentIndex string, path string, csvWriter *csv.Writer) error {
	creationDate, err := pointerCreationDate(ctx, db, cluster, segmentIndex, path)
	if err != nil {
		return err
	}
	encodedPath := base64.StdEncoding.EncodeToString([]byte(path))
	return csvWriter.Write([]string{
		cluster.projectID,
		segmentIndex,
		cluster.bucket,
		encodedPath,
		creationDate,
	})
}

func pointerCreationDate(ctx context.Context, db metainfo.PointerDB, cluster Cluster, segmentIndex string, path string) (string, error) {
	key := []byte(storj.JoinPaths(cluster.projectID, segmentIndex, cluster.bucket, path))
	pointerBytes, err := db.Get(ctx, key)
	if err != nil {
		return "", err
	}

	pointer := &pb.Pointer{}
	err = proto.Unmarshal(pointerBytes, pointer)
	if err != nil {
		return "", err
	}
	return pointer.CreationDate.String(), nil
}
func findOrCreate(cluster Cluster, path string, objects ObjectsMap) *Object {
	objectsMap, ok := objects[cluster]
	if !ok {
		objectsMap = make(map[storj.Path]*Object)
		objects[cluster] = objectsMap
	}

	object, ok := objectsMap[path]
	if !ok {
		object = &Object{}
		objectsMap[path] = object
	}

	return object
}
