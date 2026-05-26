// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package basex

import "io"

// Much of this code is adopted from Go's encoding/base64

// EncodeToString returns the baseX encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string { _ = "STUB: not implemented"; return "" }

type encoder struct {
	err  error
	enc  *Encoding
	w    io.Writer
	buf  []byte // buffered data waiting to be encoded
	nbuf int    // number of bytes in buf
	out  []byte // output buffer
}

func (e *encoder) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Leading fringe.

// Large interior chunks.

// Trailing fringe.

// Close flushes any pending output from the encoder.
// It is an error to call Write after calling Close.
func (e *encoder) Close() error {
	_ = "STUB: not implemented"
	// If there's anything left in the buffer, flush it out
	return nil
}

// NewEncoder returns a new baseX stream encoder.  Data written to
// the returned writer will be encoded using enc and then written to w.
// Encodings operate in enc.baseXBlockLen-byte blocks; when finished
// writing, the caller must Close the returned encoder to flush any
// partially written blocks.
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// DecodeString returns the bytes represented by the baseX string s.
// It uses the liberal decoding strategy, ignoring any non-baseX-characters
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type decoder struct {
	err        error
	enc        *Encoding
	r          io.Reader
	out        []byte // leftover decoded output
	buf        []byte // leftover input
	nbuf       int    // the begin pointer of buf above
	scratchbuf []byte // a temporary scratch buf, for reuse
}

func (d *decoder) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Use leftover decoded output from last read.

// Try to read up to the next full block.

// The num bytes to decode should be along obl-aligned boundaries, unless
// we're at the end of file.

// If we have too many bytes for the given buffer, we can buffer
// the rest internally

// Shift the bytes in d.buf over from [numBytesToDecode:] to the start of the array

type filteringReader struct {
	wrapped io.Reader
	enc     *Encoding
	nRead   int
}

func (r *filteringReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO: Return n, i.e. partial results?

// We want this byte. We only need to rewrite it if
// offset is behind i (otherwise we'd just be writing the
// same byte again, over itself).

// Previous buffer entirely whitespace, read again

// NewDecoder constructs a new baseX stream decoder.
func NewDecoder(enc *Encoding, r io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func newDecoder(enc *Encoding, r io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}
