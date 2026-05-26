// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"bytes"
	"io"

	"github.com/keybase/saltpack/encoding/basex"
)

// armorParams specify armor formatting, encoding and punctuation.
type armorParams struct {
	// BytesPerWord is the number of characters in each "word" of output.
	// We'll put spaces between words.
	BytesPerWord int
	// WordsPerLine is the number of words for each line of output. We'll
	// put newlines between two subsequent lines of output.
	WordsPerLine int
	// Punctuation is the byte inserted after the three "sentences" of
	// our encoding -- the header, the body and the footer.
	Punctuation byte
	// Encoding is the basex encoding to use, including strictness parameters
	Encoding *basex.Encoding
}

type armorEncoderStream struct {
	buf     *bytes.Buffer
	footer  string
	encoded io.Writer
	encoder io.WriteCloser
	nWords  int
	params  armorParams
}

func (s *armorEncoderStream) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *armorEncoderStream) spaceAndOutputBuffer() error { _ = "STUB: not implemented"; return nil }

func (s *armorEncoderStream) Close() (err error) { _ = "STUB: not implemented"; return nil }

// newArmorEncoderStream makes a new Armor encoding stream, using the given encoding
// Pass it an `encoded` stream writer to write the
// encoded stream to.  Also pass a header, and a footer string.  It will
// return an io.WriteCloser on success, that you can write raw (unencoded) data to.
// An error will be returned if there is trouble writing the header to encoded.
//
// To make the output look pretty, a space is inserted every 15 characters of output,
// and a newline is inserted every 200 words.
func newArmorEncoderStream(encoded io.Writer, header string, footer string, params armorParams) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// armorSeal takes an input plaintext and returns and output armor encoding
// as a string, or an error if a problem was encountered. Also provide a header
// and a footer to frame the message.
func armorSeal(plaintext []byte, header string, footer string, params armorParams) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Frame is a way to read the frame out of a Decoder stream.
type Frame interface {
	// GetHeader() returns the header of the frame associated with this stream, or an error
	GetHeader() (string, error)
	// GetFooter() returns the footer of the frame associated with this stream, or an error
	GetFooter() (string, error)
	// GetBrand() returns the brand contained in this frame's header, or an error
	GetBrand() (string, error)
}

type fdsState int

const (
	fdsHeader fdsState = iota
	fdsBody
	fdsFooter
	fdsEndOfStream
)

// HeaderChecker is a function intended to check that an header (the initial part of an armor
// before and not including the first punctuation) is valid. Optionally, it can return some
// information about such frame (tipycally the brand). If the frame is invalid, a non nil
// error should be returned.
type HeaderChecker func(header string) (string, error)

// FrameChecker is a function intended to check the frame of the armor is valid. Optionally,
// it can return some information about such frame (tipycally the brand). If the frame is invalid,
// a non nil error should be returned.
type FrameChecker func(header, footer string) (string, error)

type framedDecoderStream struct {
	header        []byte
	footer        []byte
	frameBrand    string
	state         fdsState
	params        armorParams
	r             *punctuatedReader
	headerChecker HeaderChecker
	frameChecker  FrameChecker
	frameLim      int // The largest frame we'll accept before we show an overflow.
}

func (s *framedDecoderStream) loadHeader() (err error) { _ = "STUB: not implemented"; return nil }

// Read from a framedDeecoderStream. The frame is the "BEGIN FOO." block
// at the beginning, and the "END FOO." block at the end.
func (s *framedDecoderStream) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// consume the stream until we hit an EOF. For all data we consume, make
// sure that it's a valid byte as far as our underlying decoder is concerned.
// We might considering clamping down here on the number of characters we're willing
// to accept after the message is over. But for now, we're quite liberal.
func (s *framedDecoderStream) consumeUntilEOF() error { _ = "STUB: not implemented"; return nil }

// isValidByteSequence checks if the byte sequence is valid as far as our
// underlying encoder is concerned.
func (s *framedDecoderStream) isValidByteSequence(p []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *framedDecoderStream) toASCII(buf []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *framedDecoderStream) GetFooter() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *framedDecoderStream) GetHeader() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *framedDecoderStream) GetBrand() (string, error) { _ = "STUB: not implemented"; return "", nil }

// newArmorDecoderStream is used to decode armored encoding. It returns a stream you
// can read from, and also a Frame you can query to see what the open/close
// frame markers were. Note that the footer of the Frame can be accessed only after the
// reader has been exhausted.
func newArmorDecoderStream(r io.Reader, params armorParams, headerChecker HeaderChecker, frameChecker FrameChecker) (io.Reader, Frame, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), *new(Frame), nil
}

// armorOpen runs armor stream decoding, but on a string, and it outputs a string.
func armorOpen(msg string, params armorParams, headerChecker HeaderChecker, frameChecker FrameChecker) (body []byte, brand string, header string, footer string, err error) {
	_ = "STUB: not implemented"
	return nil, "", "", "", nil
}
