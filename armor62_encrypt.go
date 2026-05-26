// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"
)

type closeForwarder []io.WriteCloser

func (c closeForwarder) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c closeForwarder) Close() error { _ = "STUB: not implemented"; return nil }

func newEncryptArmor62Stream(version Version, ciphertext io.Writer, sender BoxSecretKey, receivers []BoxPublicKey, ephemeralKeyCreator EphemeralKeyCreator, rng encryptRNG, brand string) (plaintext io.WriteCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// NewEncryptArmor62Stream creates a stream that consumes plaintext data.
// It will write out encrypted data to the io.Writer passed in as ciphertext.
// The encryption is from the specified sender, and is encrypted for the
// given receivers.
//
// The "brand" is the optional "brand" string to put into the header
// and footer.
//
// The ciphertext is additionally armored with the recommended armor62-style format.
//
// If initialization succeeds, returns an io.WriteCloser that accepts
// plaintext data to be encrypted and a nil error. Otherwise, returns
// nil and the initialization error.
func NewEncryptArmor62Stream(version Version, ciphertext io.Writer, sender BoxSecretKey, receivers []BoxPublicKey, brand string) (plaintext io.WriteCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func encryptArmor62Seal(version Version, plaintext []byte, sender BoxSecretKey, receivers []BoxPublicKey, ephemeralKeyCreator EphemeralKeyCreator, rng encryptRNG, brand string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// EncryptArmor62Seal is the non-streaming version of NewEncryptArmor62Stream, which
// inputs a plaintext (in bytes) and output a ciphertext (as a string).
func EncryptArmor62Seal(version Version, plaintext []byte, sender BoxSecretKey, receivers []BoxPublicKey, brand string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
