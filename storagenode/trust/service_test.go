// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package trust_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/private/errs2"
	"storj.io/storj/private/testcontext"
	"storj.io/storj/private/testplanet"
)

func TestGetSignee(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	planet, err := testplanet.New(t, 1, 1, 0)
	require.NoError(t, err)
	defer ctx.Check(planet.Shutdown)

	planet.Start(ctx)

	trust := planet.StorageNodes[0].Storage2.Trust

	canceledContext, cancel := context.WithCancel(ctx)
	cancel()

	var group errgroup.Group
	group.Go(func() error {
		_, err := trust.GetSignee(canceledContext, planet.Satellites[0].ID())
		if errs2.IsCanceled(err) {
			return nil
		}
		// if the other goroutine races us,
		// then we might get the certificate from the cache, however we shouldn't get an error
		return err
	})

	group.Go(func() error {
		cert, err := trust.GetSignee(ctx, planet.Satellites[0].ID())
		if err != nil {
			return err
		}
		if cert == nil {
			return errors.New("didn't get certificate")
		}
		return nil
	})

	assert.NoError(t, group.Wait())
}

func TestGetAddress(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 5, StorageNodeCount: 1, UplinkCount: 0,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {

		// test address is stored correctly
		for _, sat := range planet.Satellites {
			address, err := planet.StorageNodes[0].Storage2.Trust.GetAddress(ctx, sat.ID())
			require.NoError(t, err)
			assert.Equal(t, sat.Addr(), address)
		}

		var group errgroup.Group

		// test parallel reads
		for i := 0; i < 10; i++ {
			group.Go(func() error {
				address, err := planet.StorageNodes[0].Storage2.Trust.GetAddress(ctx, planet.Satellites[0].ID())
				if err != nil {
					return err
				}
				assert.Equal(t, planet.Satellites[0].Addr(), address)
				return nil
			})
		}

		assert.NoError(t, group.Wait())
	})
}
