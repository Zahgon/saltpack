// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"
)

type signcryptOpenStream struct {
	mps              *msgpackStream
	payloadKey       *SymmetricKey
	signingPublicKey SigningPublicKey
	senderAnonymous  bool
	headerHash       headerHash
	keyring          SigncryptKeyring
	resolver         SymmetricKeyResolver
}

func (sos *signcryptOpenStream) getNextChunk() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sos *signcryptOpenStream) readHeader() error {
	_ = "STUB: not implemented"
	// Read the header bytes.
	return nil
}

// Compute the header hash.

// Parse the header bytes.

func (sos *signcryptOpenStream) tryBoxSecretKeys(hdr *SigncryptionHeader, ephemeralPub BoxPublicKey) (*SymmetricKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try each of the box secret keys against each of the receiver pairs in
// the message header. The actual expected number of box secret keys is
// one, so this shouldn't be as quadratic as it looks.

//nolint:gosec // receiverIndex is a valid slice index, conversion is safe

// This is the right key! Open the sender secretbox and return the sender key.
//nolint:gosec // receiverIndex is a valid slice index, conversion is safe

// None of the box keys worked. We'll fall back to the secretbox keys.

func (sos *signcryptOpenStream) trySharedSymmetricKeys(hdr *SigncryptionHeader, ephemeralPub BoxPublicKey) (*SymmetricKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This key didn't resolve.

// We got a key. It should decrypt the corresponding receiver secretbox.

// should be statically impossible, if the slice above is the right length

//nolint:gosec // index is a valid slice index, conversion is safe

// If we get out of the loop, all the resolved keys were nil (failed to resolve).

func (sos *signcryptOpenStream) processHeader(hdr *SigncryptionHeader) error {
	_ = "STUB: not implemented"
	return nil
}

// Decrypt the sender's public key, and check for anonymous mode.

// anonymous mode, an all zero sender signing public key

// regular mode, with a real signing public key

func (sos *signcryptOpenStream) processBlock(payloadCiphertext []byte, isFinal bool, seqno packetSeqno) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle anonymous sender mode by skipping signature verification. By
// convention the signature bytes are all zeroes, but here we ignore them.

// NewSigncryptOpenStream starts a streaming verification and decryption. It
// synchronously ingests and parses the given Reader's encryption header. It
// consults the passed keyring for the decryption keys needed to decrypt the
// message. On failure, it returns a null Reader and an error message. On
// success, it returns a Reader with the plaintext stream, and a nil error. In
// either case, it will return a `MessageKeyInfo` which tells about who the
// sender was, and which of the Receiver's keys was used to decrypt the
// message.
//
// Note that the caller has an opportunity not to ingest the plaintext if he
// doesn't trust the sender revealed in the MessageKeyInfo.
func NewSigncryptOpenStream(r io.Reader, keyring SigncryptKeyring, resolver SymmetricKeyResolver) (senderPub SigningPublicKey, plaintext io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey), *new(io.Reader), nil
}

// SymmetricKeyResolver is an interface for resolving identifiers to keys.
type SymmetricKeyResolver interface {
	ResolveKeys(identifiers [][]byte) ([]*SymmetricKey, error)
}

// SigncryptOpen simply opens a ciphertext given the set of keys in the specified keyring.
// It returns a plaintext on success, and an error on failure. It returns the header's
// MessageKeyInfo in either case.
func SigncryptOpen(ciphertext []byte, keyring SigncryptKeyring, resolver SymmetricKeyResolver) (senderPub SigningPublicKey, plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey), nil, nil
}

// SigncryptKeyring is a combination of the Keyring and SigKeyring
// interfaces.
type SigncryptKeyring interface {
	Keyring
	SigKeyring
}
