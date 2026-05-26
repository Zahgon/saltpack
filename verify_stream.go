// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"
)

type verifyStream struct {
	mps        *msgpackStream
	header     *SignatureHeader
	headerHash headerHash
	publicKey  SigningPublicKey
}

func newVerifyStream(versionValidator VersionValidator, r io.Reader, msgType MessageType) (*verifyStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *verifyStream) getNextChunk() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *verifyStream) readHeader(versionValidator VersionValidator, msgType MessageType) error {
	_ = "STUB: not implemented"
	return nil
}

func readSignatureBlock(version Version, mps *msgpackStream) (signature, payloadChunk []byte, isFinal bool, seqno packetSeqno, err error) {
	_ = "STUB: not implemented"
	return nil, nil, false, *new(packetSeqno), nil
}

func (v *verifyStream) processBlock(signature, payloadChunk []byte, isFinal bool, seqno packetSeqno) error {
	_ = "STUB: not implemented"
	return nil
}
