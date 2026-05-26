// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import "io"

func applyBrand(f func(Version, io.Writer, SigningSecretKey, string) (io.WriteCloser, error), brand string) func(Version, io.Writer, SigningSecretKey) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return nil
}

// NewSignArmor62Stream creates a stream that consumes plaintext data.
// It will write out signed data to the io.Writer passed in as
// signedtext. The `brand` is optional, and allows you to specify
// your "white label" brand to this signature. NewSignArmor62Stream only
// generates attached signatures.
//
// The signed data is armored with the recommended armor62-style
// format.
func NewSignArmor62Stream(version Version, signedtext io.Writer, signer SigningSecretKey, brand string) (stream io.WriteCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// SignArmor62 creates an attached armored signature message of plaintext from signer.
func SignArmor62(version Version, plaintext []byte, signer SigningSecretKey, brand string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewSignDetachedArmor62Stream creates a stream that consumes plaintext data.
// It will write out the detached signature to the io.Writer passed in as
// detachedsig. The `brand` is optional, and allows you to specify
// your "white label" brand to this signature.
//
// The signed data is armored with the recommended armor62-style
// NewSignDetachedArmor62Stream only generates detached signatures.
//
// The signature is armored with the recommended armor62-style
// format.
func NewSignDetachedArmor62Stream(version Version, detachedsig io.Writer, signer SigningSecretKey, brand string) (stream io.WriteCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// SignDetachedArmor62 returns a detached armored signature of plaintext from signer.
func SignDetachedArmor62(version Version, plaintext []byte, signer SigningSecretKey, brand string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
