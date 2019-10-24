// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit_test

import (
	"context"
	"io"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"

	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/signing"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storage"
	"storj.io/storj/storagenode"
	"storj.io/storj/uplink"
)

const numObjects = 6

// exitProcessClient is used so we can pass the graceful exit process clients regardless of implementation.
type exitProcessClient interface {
	Send(*pb.StorageNodeMessage) error
	Recv() (*pb.SatelliteMessage, error)
}

func TestSuccess(t *testing.T) {
	testTransfers(t, numObjects, func(ctx *testcontext.Context, storageNodes map[storj.NodeID]*storagenode.Peer, satellite *testplanet.SatelliteSystem, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
		var pieceID storj.PieceID
		failedCount := 0
		startingNodePieceCounts, err := getAllNodePieceCounts(ctx, satellite, numObjects)
		pieceTransferMap := make(map[storj.NodeID]int)
		for {
			response, err := processClient.Recv()
			if errs.Is(err, io.EOF) {
				// Done
				break
			}
			require.NoError(t, err)

			switch m := response.GetMessage().(type) {
			case *pb.SatelliteMessage_TransferPiece:
				require.NotNil(t, m)

				// pick the first one to fail
				if pieceID.IsZero() {
					pieceID = m.TransferPiece.OriginalPieceId
				}

				if failedCount > 0 || pieceID != m.TransferPiece.OriginalPieceId {

					pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
					require.NoError(t, err)

					header, err := pieceReader.GetPieceHeader()
					require.NoError(t, err)

					orderLimit := header.OrderLimit
					originalPieceHash := &pb.PieceHash{
						PieceId:   orderLimit.PieceId,
						Hash:      header.GetHash(),
						PieceSize: pieceReader.Size(),
						Timestamp: header.GetCreationTime(),
						Signature: header.GetSignature(),
					}

					newPieceHash := &pb.PieceHash{
						PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
						Hash:      originalPieceHash.Hash,
						PieceSize: originalPieceHash.PieceSize,
						Timestamp: time.Now(),
					}

					receivingNode := storageNodes[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
					require.NotNil(t, receivingNode)

					// increment transfer count for receiving node
					value, _ := pieceTransferMap[receivingNode.ID()]
					pieceTransferMap[receivingNode.ID()] = value + 1

					signer := signing.SignerFromFullIdentity(receivingNode.Identity)

					signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
					require.NoError(t, err)

					success := &pb.StorageNodeMessage{
						Message: &pb.StorageNodeMessage_Succeeded{
							Succeeded: &pb.TransferSucceeded{
								OriginalPieceId:      m.TransferPiece.OriginalPieceId,
								OriginalPieceHash:    originalPieceHash,
								OriginalOrderLimit:   &orderLimit,
								ReplacementPieceHash: signedNewPieceHash,
							},
						},
					}
					err = processClient.Send(success)
					require.NoError(t, err)
				} else {
					failedCount++
					failed := &pb.StorageNodeMessage{
						Message: &pb.StorageNodeMessage_Failed{
							Failed: &pb.TransferFailed{
								OriginalPieceId: m.TransferPiece.OriginalPieceId,
								Error:           pb.TransferFailed_UNKNOWN,
							},
						},
					}
					err = processClient.Send(failed)
					require.NoError(t, err)
				}
			case *pb.SatelliteMessage_ExitCompleted:
				// TODO test completed signature stuff
				break
			default:
				// ignore
			}
		}

		// check that the exit has completed and we have the correct transferred/failed values
		progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
		require.NoError(t, err)

		require.EqualValues(t, numPieces, progress.PiecesTransferred)
		// even though we failed 1, it eventually succeeded, so the count should be 0
		require.EqualValues(t, 0, progress.PiecesFailed)

		// check that the receiving node counts are increased to the expected amount
		endingNodePieceCounts, err := getAllNodePieceCounts(ctx, satellite, numObjects)
		for k, v := range pieceTransferMap {
			starting, _ := startingNodePieceCounts[k]
			ending, _ := endingNodePieceCounts[k]
			require.Equal(t, starting+v, ending)
		}
		// check that the exiting node has no pieces in metainfo
		exitingEnding, _ := endingNodePieceCounts[exitingNode.ID()]
		require.Equal(t, 0, exitingEnding)

		exitingNode.GracefulExit.Chore.Loop.TriggerWait()

		// check that the exiting node has no pieces in its DB
		/* TODO enable this once SN chore is complete
		counts, err := getNodePieceCounts(ctx, exitingNode)
		require.NoError(t, err)
		require.Equal(t, 0, counts)
		*/
	})
}

func TestFailure(t *testing.T) {
	for _, tt := range []struct {
		name                string
		transferFailUnknown bool
		hashesDontMatch     bool
		badUplinkSig        bool
		badNodeSig          bool
		success             bool
		message             *pb.StorageNodeMessage
	}{
		{
			name:                "unknown error",
			hashesDontMatch:     false,
			transferFailUnknown: true,
			badNodeSig:          false,
			badUplinkSig:        false,
			success:             false,
		},
		{
			name:                "signatures valid, hashes don't match",
			transferFailUnknown: false,
			hashesDontMatch:     true,
			badNodeSig:          false,
			badUplinkSig:        false,
			success:             false,
		},
		{
			name:                "bad uplink signature",
			transferFailUnknown: false,
			hashesDontMatch:     false,
			badUplinkSig:        true,
			badNodeSig:          false,
			success:             false,
		},
		{
			name:                "bad storage node signature",
			transferFailUnknown: false,
			badUplinkSig:        false,
			hashesDontMatch:     false,
			badNodeSig:          true,
			success:             false,
		},
		{
			name:                "successful transfer",
			transferFailUnknown: false,
			badUplinkSig:        false,
			hashesDontMatch:     false,
			badNodeSig:          false,
			success:             true,
		},
	} {
		tt := tt
		firstIteration := true
		testTransfers(t, 1, func(ctx *testcontext.Context, storageNodes map[storj.NodeID]*storagenode.Peer, satellite *testplanet.SatelliteSystem, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int) {
			for {
				response, err := processClient.Recv()
				if errs.Is(err, io.EOF) {
					// Done
					break
				}
				if firstIteration {
					require.NoError(t, err)
					firstIteration = false
				} else {
					if tt.hashesDontMatch {
						// TODO check rpc error code and message
						require.Error(t, err, tt.name)
						break
					} else if tt.badUplinkSig {
						// TODO nat check the error
						require.Error(t, err, tt.name)
						break
					} else if tt.badNodeSig {
						// TODO nat check the error
						require.Error(t, err, tt.name)
						break
					} else {
						require.NoError(t, err)
					}
				}

				switch m := response.GetMessage().(type) {
				case *pb.SatelliteMessage_TransferPiece:
					require.NotNil(t, m)
					if tt.transferFailUnknown {
						tt.message = &pb.StorageNodeMessage{
							Message: &pb.StorageNodeMessage_Failed{
								Failed: &pb.TransferFailed{
									Error:           pb.TransferFailed_UNKNOWN,
									OriginalPieceId: m.TransferPiece.OriginalPieceId,
								},
							},
						}
					} else if tt.hashesDontMatch {
						pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
						require.NoError(t, err)

						header, err := pieceReader.GetPieceHeader()
						require.NoError(t, err)

						orderLimit := header.OrderLimit
						originalPieceHash := &pb.PieceHash{
							PieceId:   orderLimit.PieceId,
							Hash:      header.GetHash(),
							PieceSize: pieceReader.Size(),
							Timestamp: header.GetCreationTime(),
							Signature: header.GetSignature(),
						}

						newPieceHash := &pb.PieceHash{
							PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
							Hash:      originalPieceHash.Hash[:1],
							PieceSize: originalPieceHash.PieceSize,
							Timestamp: time.Now(),
						}

						receivingNode := storageNodes[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
						require.NotNil(t, receivingNode)
						signer := signing.SignerFromFullIdentity(receivingNode.Identity)

						signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
						require.NoError(t, err)

						tt.message = &pb.StorageNodeMessage{
							Message: &pb.StorageNodeMessage_Succeeded{
								Succeeded: &pb.TransferSucceeded{
									OriginalPieceId:      m.TransferPiece.OriginalPieceId,
									OriginalPieceHash:    originalPieceHash,
									OriginalOrderLimit:   &orderLimit,
									ReplacementPieceHash: signedNewPieceHash,
								},
							},
						}
					} else if tt.badUplinkSig {
						pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
						require.NoError(t, err)

						header, err := pieceReader.GetPieceHeader()
						require.NoError(t, err)

						orderLimit := header.OrderLimit
						orderLimit.UplinkPublicKey = storj.PiecePublicKey{}

						originalPieceHash := &pb.PieceHash{
							PieceId:   orderLimit.PieceId,
							Hash:      header.GetHash(),
							PieceSize: pieceReader.Size(),
							Timestamp: header.GetCreationTime(),
							Signature: header.GetSignature(),
						}

						newPieceHash := &pb.PieceHash{
							PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
							Hash:      originalPieceHash.Hash,
							PieceSize: originalPieceHash.PieceSize,
							Timestamp: time.Now(),
						}

						receivingNode := storageNodes[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
						require.NotNil(t, receivingNode)
						signer := signing.SignerFromFullIdentity(receivingNode.Identity)

						signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
						require.NoError(t, err)

						tt.message = &pb.StorageNodeMessage{
							Message: &pb.StorageNodeMessage_Succeeded{
								Succeeded: &pb.TransferSucceeded{
									OriginalPieceId:      m.TransferPiece.OriginalPieceId,
									OriginalPieceHash:    originalPieceHash,
									OriginalOrderLimit:   &orderLimit,
									ReplacementPieceHash: signedNewPieceHash,
								},
							},
						}
					} else if tt.badNodeSig {
						pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
						require.NoError(t, err)

						header, err := pieceReader.GetPieceHeader()
						require.NoError(t, err)

						orderLimit := header.OrderLimit

						originalPieceHash := &pb.PieceHash{
							PieceId:   orderLimit.PieceId,
							Hash:      header.GetHash(),
							PieceSize: pieceReader.Size(),
							Timestamp: header.GetCreationTime(),
							Signature: header.GetSignature(),
						}

						newPieceHash := &pb.PieceHash{
							PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
							Hash:      originalPieceHash.Hash,
							PieceSize: originalPieceHash.PieceSize,
							Timestamp: time.Now(),
						}

						wrongSigner := signing.SignerFromFullIdentity(exitingNode.Identity)

						signedNewPieceHash, err := signing.SignPieceHash(ctx, wrongSigner, newPieceHash)
						require.NoError(t, err)

						tt.message = &pb.StorageNodeMessage{
							Message: &pb.StorageNodeMessage_Succeeded{
								Succeeded: &pb.TransferSucceeded{
									OriginalPieceId:      m.TransferPiece.OriginalPieceId,
									OriginalPieceHash:    originalPieceHash,
									OriginalOrderLimit:   &orderLimit,
									ReplacementPieceHash: signedNewPieceHash,
								},
							},
						}
					} else {
						pieceReader, err := exitingNode.Storage2.Store.Reader(ctx, satellite.ID(), m.TransferPiece.OriginalPieceId)
						require.NoError(t, err)

						header, err := pieceReader.GetPieceHeader()
						require.NoError(t, err)

						orderLimit := header.OrderLimit
						originalPieceHash := &pb.PieceHash{
							PieceId:   orderLimit.PieceId,
							Hash:      header.GetHash(),
							PieceSize: pieceReader.Size(),
							Timestamp: header.GetCreationTime(),
							Signature: header.GetSignature(),
						}

						newPieceHash := &pb.PieceHash{
							PieceId:   m.TransferPiece.AddressedOrderLimit.Limit.PieceId,
							Hash:      originalPieceHash.Hash,
							PieceSize: originalPieceHash.PieceSize,
							Timestamp: time.Now(),
						}

						receivingNode := storageNodes[m.TransferPiece.AddressedOrderLimit.Limit.StorageNodeId]
						require.NotNil(t, receivingNode)
						signer := signing.SignerFromFullIdentity(receivingNode.Identity)

						signedNewPieceHash, err := signing.SignPieceHash(ctx, signer, newPieceHash)
						require.NoError(t, err)

						tt.message = &pb.StorageNodeMessage{
							Message: &pb.StorageNodeMessage_Succeeded{
								Succeeded: &pb.TransferSucceeded{
									OriginalPieceId:      m.TransferPiece.OriginalPieceId,
									OriginalPieceHash:    originalPieceHash,
									OriginalOrderLimit:   &orderLimit,
									ReplacementPieceHash: signedNewPieceHash,
								},
							},
						}
					}
					err = processClient.Send(tt.message)
					require.NoError(t, err)
				case *pb.SatelliteMessage_ExitCompleted:
					// TODO test completed signature stuff
					break
				case *pb.SatelliteMessage_ExitFailed:
					status, err := satellite.DB.OverlayCache().GetExitStatus(ctx, exitingNode.ID())
					require.NoError(t, err)
					require.False(t, status.ExitSuccess)
					require.Equal(t, m.ExitFailed.Reason, pb.ExitFailed_OVERALL_FAILURE_PERCENTAGE_EXCEEDED)
					break
				default:
					//t.FailNow()
				}
			}
			// TODO uncomment once progress reflects updated success and fail counts
			// check that the exit has completed and we have the correct transferred/failed values
			//progress, err := satellite.DB.GracefulExit().GetProgress(ctx, exitingNode.ID())
			//require.NoError(t, err)
			//
			//if tt.success {
			//	require.Equal(t, int64(1), progress.PiecesTransferred, tt.name)
			//	require.Equal(t, int64(0), progress.PiecesFailed, tt.name)
			//} else {
			//	require.Equal(t, int64(0), progress.PiecesTransferred, tt.name)
			//	require.Equal(t, int64(1), progress.PiecesFailed, tt.name)
			//}
		})
	}
}

func testTransfers(t *testing.T, objects int, verifier func(ctx *testcontext.Context, storageNodes map[storj.NodeID]*storagenode.Peer, satellite *testplanet.SatelliteSystem, processClient exitProcessClient, exitingNode *storagenode.Peer, numPieces int)) {
	successThreshold := 8
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: successThreshold + 1,
		UplinkCount:      1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		uplinkPeer := planet.Uplinks[0]
		satellite := planet.Satellites[0]

		satellite.GracefulExit.Chore.Loop.Pause()

		storageNodes := make(map[storj.NodeID]*storagenode.Peer)
		for _, node := range planet.StorageNodes {
			storageNodes[node.ID()] = node
		}

		rs := &uplink.RSConfig{
			MinThreshold:     4,
			RepairThreshold:  6,
			SuccessThreshold: successThreshold,
			MaxThreshold:     successThreshold,
		}

		for i := 0; i < objects; i++ {
			err := uplinkPeer.UploadWithConfig(ctx, satellite, rs, "testbucket", "test/path"+strconv.Itoa(i), testrand.Bytes(5*memory.KiB))
			require.NoError(t, err)
		}
		// check that there are no exiting nodes.
		exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
		require.NoError(t, err)
		require.Len(t, exitingNodes, 0)

		exitingNode, err := findNodeToExit(ctx, planet, objects)
		require.NoError(t, err)

		// connect to satellite so we initiate the exit.
		conn, err := exitingNode.Dialer.DialAddressID(ctx, satellite.Addr(), satellite.Identity.ID)
		require.NoError(t, err)
		defer func() {
			err = errs.Combine(err, conn.Close())
		}()

		client := conn.SatelliteGracefulExitClient()

		c, err := client.Process(ctx)
		require.NoError(t, err)

		response, err := c.Recv()
		require.NoError(t, err)

		// should get a NotReady since the metainfo loop would not be finished at this point.
		switch response.GetMessage().(type) {
		case *pb.SatelliteMessage_NotReady:
			// now check that the exiting node is initiated.
			exitingNodes, err := satellite.DB.OverlayCache().GetExitingNodes(ctx)
			require.NoError(t, err)
			require.Len(t, exitingNodes, 1)

			require.Equal(t, exitingNode.ID(), exitingNodes[0].NodeID)
		default:
			//t.FailNow()
		}
		// close the old client
		require.NoError(t, c.CloseSend())

		// trigger the metainfo loop chore so we can get some pieces to transfer
		satellite.GracefulExit.Chore.Loop.TriggerWait()

		// make sure all the pieces are in the transfer queue
		incompleteTransfers, err := satellite.DB.GracefulExit().GetIncomplete(ctx, exitingNode.ID(), objects, 0)
		require.NoError(t, err)

		// connect to satellite again to start receiving transfers
		c, err = client.Process(ctx)
		require.NoError(t, err)
		defer func() {
			err = errs.Combine(err, c.CloseSend())
		}()

		verifier(ctx, storageNodes, satellite, c, exitingNode, len(incompleteTransfers))
	})
}

func findNodeToExit(ctx context.Context, planet *testplanet.Planet, objects int) (*storagenode.Peer, error) {
	satellite := planet.Satellites[0]
	keys, err := satellite.Metainfo.Database.List(ctx, nil, objects)
	if err != nil {
		return nil, err
	}

	pieceCountMap := make(map[storj.NodeID]int, len(planet.StorageNodes))
	for _, sn := range planet.StorageNodes {
		pieceCountMap[sn.ID()] = 0
	}

	for _, key := range keys {
		pointer, err := satellite.Metainfo.Service.Get(ctx, string(key))
		if err != nil {
			return nil, err
		}
		pieces := pointer.GetRemote().GetRemotePieces()
		for _, piece := range pieces {
			pieceCountMap[piece.NodeId]++
		}
	}

	var exitingNodeID storj.NodeID
	maxCount := 0
	for k, v := range pieceCountMap {
		if exitingNodeID.IsZero() {
			exitingNodeID = k
			maxCount = v
			continue
		}
		if v > maxCount {
			exitingNodeID = k
			maxCount = v
		}
	}

	for _, sn := range planet.StorageNodes {
		if sn.ID() == exitingNodeID {
			return sn, nil
		}
	}

	return nil, nil
}

// getNodeCounts piece counts for all nodes.
func getAllNodePieceCounts(ctx context.Context, satellite *testplanet.SatelliteSystem, objects int) (_ map[storj.NodeID]int, err error) {
	keys, err := satellite.Metainfo.Database.List(ctx, nil, objects)
	if err != nil {
		return nil, err
	}
	nodePieceCounts := make(map[storj.NodeID]int)
	for _, key := range keys {
		pointer, err := satellite.Metainfo.Service.Get(ctx, string(key))
		if err != nil {
			return nil, err
		}
		pieces := pointer.GetRemote().GetRemotePieces()
		for _, piece := range pieces {
			value, _ := nodePieceCounts[piece.NodeId]
			nodePieceCounts[piece.NodeId] = value + 1
		}
	}

	return nodePieceCounts, nil
}

// getNodePieceCounts tallies all the pieces per node.
func getNodePieceCounts(ctx context.Context, node *storagenode.Peer) (_ int, err error) {
	nodePieceCounts := 0
	namespaces, err := node.DB.Pieces().ListNamespaces(ctx)
	if err != nil {
		return nodePieceCounts, err
	}
	for _, ns := range namespaces {
		err = node.DB.Pieces().WalkNamespace(ctx, ns, func(blobInfo storage.BlobInfo) error {
			nodePieceCounts++
			return nil
		})
		if err != nil {
			return nodePieceCounts, err
		}
	}

	return nodePieceCounts, err
}
