// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"io"
)

type decryptStream struct {
	versionValidator VersionValidator
	version          Version
	ring             Keyring
	mps              *msgpackStream
	payloadKey       *SymmetricKey
	senderKey        *RawBoxKey
	headerHash       headerHash
	macKey           macKey
	position         int
	mki              MessageKeyInfo
}

// MessageKeyInfo conveys all of the data about the keys used in this encrypted message.
type MessageKeyInfo struct {
	// These fields are cryptographically verified
	SenderKey      BoxPublicKey
	SenderIsAnon   bool
	ReceiverKey    BoxSecretKey
	ReceiverIsAnon bool

	// These fields are not cryptographically verified, and are just repeated from what
	// we saw in the incoming message.
	NamedReceivers   [][]byte
	NumAnonReceivers int
}

func (ds *decryptStream) getNextChunk() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ds *decryptStream) readHeader(_ io.Reader) error {
	_ = "STUB: not implemented"
	// Read the header bytes.
	return nil
}

// Compute the header hash.

// Parse the header bytes.

func readEncryptionBlock(version Version, mps *msgpackStream) (ciphertext []byte, authenticators []payloadAuthenticator, isFinal bool, seqno packetSeqno, err error) {
	_ = "STUB: not implemented"
	return nil, nil, false, *new(packetSeqno), nil
}

func (ds *decryptStream) tryVisibleReceivers(hdr *EncryptionHeader, ephemeralKey BoxPublicKey) (BoxSecretKey, *SymmetricKey, int, error) {
	_ = "STUB: not implemented"
	return *new(BoxSecretKey), nil, 0, nil
}

// Keep track of where it was in the original list

//nolint:gosec // orig is a valid slice index, conversion is safe

func (ds *decryptStream) tryHiddenReceivers(hdr *EncryptionHeader, ephemeralKey BoxPublicKey) (BoxSecretKey, *SymmetricKey, int, error) {
	_ = "STUB: not implemented"
	return *new(BoxSecretKey), nil, 0, nil
}

//nolint:gosec // i is a valid slice index, conversion is safe

func (ds *decryptStream) processHeader(hdr *EncryptionHeader) error {
	_ = "STUB: not implemented"
	return nil
}

// Decrypt the sender's public key

// Lookup the sender's public key in our keyring, and import
// it for use. However, if the sender key is the same as the ephemeral
// key, then assume "anonymous mode", so use the already imported anonymous
// key.

// Compute the MAC key.

func computeMACKeyReceiver(version Version, index uint64, secret BoxSecretKey, public, ePublic BoxPublicKey, headerHash headerHash) macKey {
	_ = "STUB: not implemented"
	// Switch on the major version since we're reading, and so may
	// encounter headers written by unknown minor versions.
	return *new(macKey)
}

func (ds *decryptStream) processBlock(ciphertext []byte, authenticators []payloadAuthenticator, isFinal bool, seqno packetSeqno) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check the authenticator.

// The encoding of the empty buffer implies the EOF.  But otherwise, all mechanisms are the same.

// NewDecryptStream starts a streaming decryption. It synchronously ingests
// and parses the given Reader's encryption header. It consults the passed
// keyring for the decryption keys needed to decrypt the message. On failure,
// it returns a null Reader and an error message. On success, it returns a
// Reader with the plaintext stream, and a nil error. In either case, it will
// return a `MessageKeyInfo` which tells about who the sender was, and which of the
// Receiver's keys was used to decrypt the message.
//
// Note that the caller has an opportunity not to ingest the plaintext if he
// doesn't trust the sender revealed in the MessageKeyInfo.
func NewDecryptStream(versionValidator VersionValidator, r io.Reader, keyring Keyring) (mki *MessageKeyInfo, plaintext io.Reader, err error) {
	_ = "STUB: not implemented"
	return nil, *new(io.Reader), nil
}

// Open simply opens a ciphertext given the set of keys in the specified keyring.
// It returns a plaintext on success, and an error on failure. It returns the header's
// MessageKeyInfo in either case.
func Open(versionValidator VersionValidator, ciphertext []byte, keyring Keyring) (i *MessageKeyInfo, plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
