// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gracefulexit

import (
	"bytes"
	"context"

	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/signing"
)

func (endpoint *Endpoint) validatePendingTransfer(ctx context.Context, transfer *pendingTransfer) error {
	if transfer.satelliteMessage == nil {
		return Error.New("Satellite message cannot be nil")
	}
	if transfer.satelliteMessage.GetTransferPiece() == nil {
		return Error.New("Satellite message transfer piece cannot be nil")
	}
	if transfer.satelliteMessage.GetTransferPiece().GetAddressedOrderLimit() == nil {
		return Error.New("Addressed order limit on transfer piece cannot be nil")
	}
	if transfer.satelliteMessage.GetTransferPiece().GetAddressedOrderLimit().GetLimit() == nil {
		return Error.New("Addressed order limit on transfer piece cannot be nil")
	}
	if transfer.path == nil {
		return Error.New("Transfer path cannot be nil")
	}
	if transfer.originalPointer == nil || transfer.originalPointer.GetRemote() == nil {
		return Error.New("could not get remote pointer from transfer item")
	}

	return nil
}

func (endpoint *Endpoint) verifyPieceTransferred(ctx context.Context, message *pb.StorageNodeMessage_Succeeded, transfer *pendingTransfer, receivingNodePeerID *identity.PeerIdentity) error {
	originalOrderLimit := message.Succeeded.GetOriginalOrderLimit()
	if originalOrderLimit == nil {
		return ErrInvalidArgument.New("Original order limit cannot be nil")
	}
	originalPieceHash := message.Succeeded.GetOriginalPieceHash()
	if originalPieceHash == nil {
		return ErrInvalidArgument.New("Original piece hash cannot be nil")
	}
	replacementPieceHash := message.Succeeded.GetReplacementPieceHash()
	if replacementPieceHash == nil {
		return ErrInvalidArgument.New("Replacement piece hash cannot be nil")
	}

	// verify that the original piece hash and replacement piece hash match
	if !bytes.Equal(originalPieceHash.Hash, replacementPieceHash.Hash) {
		return ErrInvalidArgument.New("Piece hashes for transferred piece don't match")
	}

	// verify that the satellite signed the original order limit
	err := signing.VerifyOrderLimitSignature(ctx, endpoint.signer, originalOrderLimit)
	if err != nil {
		return ErrInvalidArgument.Wrap(err)
	}

	// verify that the public key on the order limit signed the original piece hash
	err = signing.VerifyUplinkPieceHashSignature(ctx, originalOrderLimit.UplinkPublicKey, originalPieceHash)
	if err != nil {
		return ErrInvalidArgument.Wrap(err)
	}

	if originalOrderLimit.PieceId != message.Succeeded.OriginalPieceId {
		return ErrInvalidArgument.New("Invalid original piece ID")
	}

	receivingNodeID := transfer.satelliteMessage.GetTransferPiece().GetAddressedOrderLimit().GetLimit().StorageNodeId
	calculatedNewPieceID := transfer.originalPointer.GetRemote().RootPieceId.Derive(receivingNodeID, transfer.pieceNum)
	if calculatedNewPieceID != replacementPieceHash.PieceId {
		return ErrInvalidArgument.New("Invalid replacement piece ID")
	}

	signee := signing.SigneeFromPeerIdentity(receivingNodePeerID)

	// verify that the new node signed the replacement piece hash
	err = signing.VerifyPieceHashSignature(ctx, signee, replacementPieceHash)
	if err != nil {
		return ErrInvalidArgument.Wrap(err)
	}
	return nil
}
