// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"

	"github.com/keybase/saltpack/encoding/basex"
)

// Armor62Params are the armoring parameters we recommend for use with
// a generic armorer.  It specifies the spaces between words, the spacing
// between lines, some simple punctuation, and an encoding alphabet.
var Armor62Params = armorParams{
	BytesPerWord: 15,
	WordsPerLine: 200,
	Punctuation:  byte('.'),
	Encoding:     basex.Base62StdEncoding,
}

// NewArmor62EncoderStream makes a new Armor 62 encoding stream, using the base62-alphabet
// and a 32/43 encoding rate strategy. Pass it an `encoded` stream writer to write the
// encoded stream to.  Also pass an optional "brand" .  It will
// return an io.WriteCloser on success, that you can write raw (unencoded) data to.
// An error will be returned if there is trouble writing the header to encoded.
//
// To make the output look pretty, a space is inserted every 15 characters of output,
// and a newline is inserted every 200 words.
func NewArmor62EncoderStream(encoded io.Writer, typ MessageType, brand string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// Armor62Seal takes an input plaintext and returns and output armor encoding
// as a string, or an error if a problem was encountered. Also provide a header
// and a footer to frame the message. Uses Base62 encoding scheme
func Armor62Seal(plaintext []byte, typ MessageType, brand string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewArmor62DecoderStream is used to decode input base62-armoring format. It returns
// a stream you can read from, and also a Frame you can query to see what the open/close
// frame markers were. hc and fc are optional and can be nil.
func NewArmor62DecoderStream(r io.Reader, hc HeaderChecker, fc FrameChecker) (io.Reader, Frame, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), *new(Frame), nil
}

// Armor62Open runs armor stream decoding, but on a string, and it outputs
// a string. It does not do any validation on the header and footer.
//
// Deprecated: use Armor62OpenWithValidation instead.
func Armor62Open(msg string) (body []byte, header string, footer string, err error) {
	_ = "STUB: not implemented"
	return nil, "", "", nil
}

// Armor62OpenWithValidation runs armor stream decoding, but on a string, and it outputs
// a string. It validates header and footer with the provided checkers (which are optional and can be nil).
func Armor62OpenWithValidation(msg string, hc HeaderChecker, fc FrameChecker) (body []byte, brand string, header string, footer string, err error) {
	_ = "STUB: not implemented"
	return nil, "", "", "", nil
}

// CheckArmor62Frame checks that the frame matches our standard
// begin/end frame
func CheckArmor62Frame(frame Frame, typ MessageType) (brand string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CheckArmor62 checks that the frame matches our standard
// begin/end frame
func CheckArmor62(hdr string, ftr string, typ MessageType) (brand string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
