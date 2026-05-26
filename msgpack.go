// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"

	"github.com/keybase/go-codec/codec"
)

type encoder interface {
	Encode(v any) error
}

func newEncoder(w io.Writer) encoder { _ = "STUB: not implemented"; return *new(encoder) }

func encodeToBytes(i any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// If p has type {Encryption,Signcryption,Signature}Header, then
// decodeFromBytes would succeed even if the version numbers or mode
// aren't encoded as positive fixnums. Ideally, it would reject those
// as malformed, but there's no easy way to do that.
//
// Similarly, we'd ideally reject strings, byte arrays, or arrays that
// aren't minimally encoded, but there's no easy way to check that
// either.

func decodeFromBytes(p any, b []byte) error { _ = "STUB: not implemented"; return nil }

type msgpackStream struct {
	decoder *codec.Decoder
	seqno   packetSeqno
}

func newMsgpackStream(r io.Reader) *msgpackStream { _ = "STUB: not implemented"; return nil }

func (r *msgpackStream) Read(i any) (ret packetSeqno, err error) {
	_ = "STUB: not implemented"
	return *new(packetSeqno), nil
}
