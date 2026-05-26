// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"
)

var (
	armor62EncryptionHeaderChecker HeaderChecker = func(header string) (string, error) {
		return parseFrame(header, MessageTypeEncryption, headerMarker)
	}
	armor62EncryptionFrameChecker FrameChecker = func(header, footer string) (string, error) {
		return CheckArmor62(header, footer, MessageTypeEncryption)
	}
)

// NewDearmor62DecryptStream makes a new stream that dearmors and decrypts the given
// Reader stream. Pass it a keyring so that it can lookup private and public keys
// as necessary. Returns the MessageKeyInfo recovered during header
// processing, an io.Reader stream from which you can read the plaintext, the armor branding, and
// maybe an error if there was a failure.
func NewDearmor62DecryptStream(versionValidator VersionValidator, ciphertext io.Reader, kr Keyring) (mki *MessageKeyInfo, ds io.Reader, brand string, err error) {
	_ = "STUB: not implemented"
	return nil, *new(io.Reader), "", nil
}

// Dearmor62DecryptOpen takes an armor62'ed, encrypted ciphertext and attempts to
// dearmor and decrypt it, using the provided keyring. Checks that the frames in the
// armor are as expected. Returns the MessageKeyInfo recovered during message
// processing, the plaintext (if decryption succeeded), the armor branding, and
// maybe an error if there was a failure.
func Dearmor62DecryptOpen(versionValidator VersionValidator, ciphertext string, kr Keyring) (*MessageKeyInfo, []byte, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}
