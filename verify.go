// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"
)

// NewVerifyStream creates a stream that consumes data from reader
// r.  It returns the signer's public key and a reader that only
// contains verified data.  If the signer's key is not in keyring,
// it will return an error.
func NewVerifyStream(versionValidator VersionValidator, r io.Reader, keyring SigKeyring) (skey SigningPublicKey, vs io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey), *new(io.Reader), nil
}

// Verify checks the signature in signedMsg. It returns the
// signer's public key and a verified message.
func Verify(versionValidator VersionValidator, signedMsg []byte, keyring SigKeyring) (skey SigningPublicKey, verifiedMsg []byte, err error) {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey), nil, nil
}

// VerifyDetachedReader verifies that signature is a valid signature for
// entire message read from message Reader, and that the public key for
// the signer is in keyring. It returns the signer's public key.
func VerifyDetachedReader(versionValidator VersionValidator, message io.Reader, signature []byte, keyring SigKeyring) (skey SigningPublicKey, err error) {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey), nil
}

// Use a verifyStream to parse the header.

// Reach inside the verifyStream to parse the signature bytes.

// Get the public key.

// Compute the signed text hash, without requiring us to copy the whole
// signed text into memory at once.

// VerifyDetached verifies that signature is a valid signature for
// message, and that the public key for the signer is in keyring.
// It returns the signer's public key.
func VerifyDetached(versionValidator VersionValidator, message, signature []byte, keyring SigKeyring) (skey SigningPublicKey, err error) {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey), nil
}
