// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"bytes"
	"hash"
	"io"
)

type signAttachedStream struct {
	version    Version
	headerHash headerHash
	encoder    encoder
	buffer     bytes.Buffer
	seqno      packetSeqno
	secretKey  SigningSecretKey
}

func newSignAttachedStream(version Version, w io.Writer, signer SigningSecretKey) (*signAttachedStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encode the header bytes.

// Compute the header hash.

// Create the attached stream object.

// Double encode the header bytes onto the wire.

func (s *signAttachedStream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// If s.buffer.Len() == signatureBlockSize, we don't want to
// write it out just yet, since for V2 we need to be sure this
// isn't the last block.

func (s *signAttachedStream) Close() error { _ = "STUB: not implemented"; return nil }

func makeSignatureBlock(version Version, sig, chunk []byte, isFinal bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func checkSignBlockRead(version Version, isFinal bool, blockSize, chunkLen, bufLen int) {
	_ = "STUB: not implemented"
	return
}

// We shouldn't read more than a full block's worth.

// If we read less than a full block's worth, then we
// shouldn't have anything left in the buffer.

// isFinal must be equivalent to chunkLen being 0
// (which, by the above, implies that bufLen == 0).

// If isFinal, then chunkLen can be any number,
// but bufLen must be 0.

func (s *signAttachedStream) signBlock(isFinal bool) error {
	_ = "STUB: not implemented"
	// NOTE: chunk is a slice into s.buffer's buffer, so make sure
	// not to stash it anywhere.
	return nil
}

func (s *signAttachedStream) computeSig(payloadChunk []byte, seqno packetSeqno, isFinal bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type signDetachedStream struct {
	encoder   encoder
	secretKey SigningSecretKey
	hasher    hash.Hash
}

func newSignDetachedStream(version Version, w io.Writer, signer SigningSecretKey) (*signDetachedStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encode the header bytes.

// Compute the header hash.

// Create the detached stream object.

// Double encode the header bytes onto the wire.

// Start off the message digest with the header hash. Subsequent calls to
// Write() will push message bytes into this digest.

func (s *signDetachedStream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *signDetachedStream) Close() error { _ = "STUB: not implemented"; return nil }
