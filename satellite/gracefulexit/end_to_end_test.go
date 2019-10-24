// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"

	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/uplink"
)

func TestEndToEnd(t *testing.T) {
	nodes := 20
	objects := 100
	exits := 3

	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: nodes,
		UplinkCount:      1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		rs := &uplink.RSConfig{
			MinThreshold:     4,
			RepairThreshold:  6,
			SuccessThreshold: 8,
			MaxThreshold:     10,
		}

		for i := 0; i < objects; i++ {
			err := uplinkPeer.UploadWithConfig(ctx, satellite, rs, "testbucket", "test/path"+strconv.Itoa(i), testrand.Bytes(5*memory.KiB))
			require.NoError(t, err)
		}

		logPieces(t, ctx, planet, satellite, objects, true)

		for i, node := range planet.StorageNodes {
			if i >= exits {
				break
			}
			t.Logf("GETEST: exiting node %v\n", i)
			err := node.DB.Satellites().InitiateGracefulExit(ctx, satellite.ID(), time.Now().UTC(), 1)
			require.NoError(t, err)
		}

		for {
			time.Sleep(time.Second * 1)
			complete := 0

			for i, node := range planet.StorageNodes {
				if i >= exits {
					break
				}
				exitStatus, err := satellite.DB.OverlayCache().GetExitStatus(ctx, node.ID())
				require.NoError(t, err)
				if exitStatus.ExitFinishedAt != nil {
					complete++
				}
			}
			if complete == exits {
				t.Logf("GETEST: completed exit processes: %v\n", complete)
				break
			}
		}

		logPieces(t, ctx, planet, satellite, objects, false)

		// try download
		t.Logf("GETEST: trying to download all files\n")
		succeeded := 0
		failed := 0
		for i := 0; i < objects; i++ {
			_, err := uplinkPeer.Download(ctx, satellite, "testbucket", "test/path"+strconv.Itoa(i))
			//require.NoError(t, err)
			if err != nil {
				failed++
				t.Logf("GETEST: test/path%v failed. %+v\n", strconv.Itoa(i), errs.Wrap(err))
			} else {
				//t.Logf("GETEST: test/path%v succeeded.\n", strconv.Itoa(i))
				succeeded++
			}
		}

		t.Logf("GETEST: succeeded: %v, failed: %v\n", succeeded, failed)
	})
}

func logPieces(t *testing.T, ctx context.Context, planet *testplanet.Planet, satellite *testplanet.SatelliteSystem, objects int, diff bool) {
	satPieceCountMap, err := getAllNodePieceCounts(ctx, satellite, objects)
	require.NoError(t, err)
	t.Logf("GETEST: piece tallies\n")
	for i, node := range planet.StorageNodes {
		count, err := getNodePieceCounts(ctx, node)
		require.NoError(t, err)
		metaCount, _ := satPieceCountMap[node.ID()]
		if diff {
			t.Logf("GETEST: node %v\tsn pieces: %v\tsat pieces: %v\tdiff: %3v\n", i, count, metaCount, count-metaCount)
		} else {
			t.Logf("GETEST: node %v\tsn pieces: %v\tsat pieces: %v\n", i, count, metaCount)
		}
	}
	segmentPieceCounts, err := getSegmentPieceCounts(t, ctx, satellite, objects)
	require.NoError(t, err)
	t.Logf("GETEST: segment piece counts < 8\n")
	for k, v := range segmentPieceCounts {
		if v < 8 {
			t.Logf("GETEST: %v has less than 8 pieces. count: %v\n", k, v)
		}
	}
}

func getSegmentPieceCounts(t *testing.T, ctx context.Context, satellite *testplanet.SatelliteSystem, objects int) (_ map[string]int, err error) {
	keys, err := satellite.Metainfo.Database.List(ctx, nil, objects)
	if err != nil {
		return nil, err
	}
	segmentPieceCounts := make(map[string]int)
	for _, key := range keys {
		pointer, err := satellite.Metainfo.Service.Get(ctx, string(key))
		if err != nil {
			return nil, err
		}

		dupTestMap := make(map[storj.NodeID]int)

		pieces := pointer.GetRemote().GetRemotePieces()
		segmentPieceCounts[key.String()] = len(pieces)
		for _, piece := range pieces {
			v, _ := dupTestMap[piece.NodeId]
			dupTestMap[piece.NodeId] = v + 1
		}
		for k, v := range dupTestMap {
			if v > 1 {
				t.Logf("GETEST: sn %v has %v pieces for a segment\n", k, v)
			}
		}
	}

	return segmentPieceCounts, nil
}
