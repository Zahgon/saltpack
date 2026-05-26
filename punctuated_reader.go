// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"errors"
	"io"
)

// punctuatedReader is a stream reader that reads until it hits a usual
// error OR until it hits a punctuation character. In that latter case
// it returns an `ErrPunctuation` "error" so that way callers can tell
// the difference between a normal EOF and a "punctuated" EOF.
type punctuatedReader struct {
	r              io.Reader
	punctuation    [1]byte
	nextSegment    []byte
	thisSegment    []byte
	errThisSegment error
	buf            [4096]byte
}

// ErrPunctuated is produced when a punctuation character is found in the stream.
// It can be returned along with data, unlike usual errors.
var ErrPunctuated = errors.New("found punctuation in stream")

// ErrOverflow is returned if we were looking for punctuation but our quota was
// overflowed before we found the needed character.
var ErrOverflow = errors.New("buffer was overflowed before we found punctuation")

// Read from the punctuatedReader, potentially returning an `ErrPunctuation`
// if a punctuation character was found.
func (p *punctuatedReader) Read(out []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// First deal with the case that we had a "short copy" to our target buffer
	// in a previous call of the read function.
	return 0, nil
}

// In this case we have no previous short copy, so now we deal with
// two cases --- we had, from a previous iteration, some stuff after the
// punctuation. Or we need to read again.

// Now look for punctuation. If we find it, we need to keep the remaining
// data in the buffer (to the right of the punctuation mark) for the next
// time through the loop. Note that that new buffer can itself have subsequent
// punctuation, so we'll have to perform the check again here.

// If we used a buffer, copy into the output buffer, and potentially deal
// with a "short copy" situation in which we couldn't fit all of the data
// into the given buffer.

// If we found punctuation, we have to set an error accordingly. However,
// in the "short copy" situation just above, we can't return the error just
// yet, we need to do so when that buffer is drained.

// ReadUntilPunctuation reads from the stream until it find a desired
// punctuation byte. If it wasn't found before EOF, it will return io.ErrUnexpectedEOF.
// If it wasn't found before lim bytes are consumed, then it will return ErrOverflow.
func (p *punctuatedReader) ReadUntilPunctuation(lim int) (res []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newPunctuatedReader returns a new punctuatedReader given an underlying
// read stream and a punctuation byte.
func newPunctuatedReader(r io.Reader, p byte) *punctuatedReader {
	_ = "STUB: not implemented"
	return nil
}
