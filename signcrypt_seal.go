// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"bytes"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/nacl/secretbox"
)

type signcryptSealStream struct {
	version       Version
	output        io.Writer
	encoder       encoder
	encryptionKey SymmetricKey
	signingKey    SigningSecretKey
	buffer        bytes.Buffer
	headerHash    headerHash

	numBlocks encryptionBlockNumber // the lower 64 bits of the nonce

	err error
}

func (sss *signcryptSealStream) Write(plaintext []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (sss *signcryptSealStream) signcryptBlock(isFinal bool) error {
	_ = "STUB: not implemented"
	// NOTE: plaintext is a slice into sss.buffer's buffer, so
	// make sure not to stash it anywhere.
	return nil
}

// Handle regular signing mode and anonymous mode (where we don't actually
// sign anything).

// Similar to the encryption format, we derive a symmetric key from our DH keys
// (one of which is ephemeral) by encrypting 32 bytes of zeros. We could have
// used crypto_box_beforenm directly instead, but that would be a slight abuse
// of that function, and also we don't expect all NaCl/libsodium wrapper libs
// to expose it. This key does *not* mix in the recipient index -- it will be
// the same for two different recipients if they claim the same public key.
func derivedEphemeralKeyFromBoxKeys(public BoxPublicKey, private BoxSecretKey) *SymmetricKey {
	_ = "STUB: not implemented"
	return nil
}

// should be statically impossible, if the slice above is the right length

// Compute the visible identifier that the recipient will use to find the right
// recipient entry. Include the entry index, so that this identifier is unique
// even if two recipients claim the same public key (though unfortunately that
// means that recipients will need to recompute the identifier for each entry
// in the recipients list). This identifier is somewhat redundant, because a
// recipient could instead just attempt to decrypt the payload key secretbox
// and see if it works, but including them adds a bit to anonymity by making
// box key recipients indistinguishable from symmetric key recipients.
func keyIdentifierFromDerivedKey(derivedKey *SymmetricKey, recipientIndex uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// A receiverKeysMaker is either a (wrapped) BoxPublicKey or a
// ReceiverSymmetricKey.
type receiverKeysMaker interface {
	makeReceiverKeys(ephemeralPriv BoxSecretKey, payloadKey SymmetricKey, index uint64) receiverKeys
}

type receiverBoxKey struct {
	pk BoxPublicKey
}

func (r receiverBoxKey) makeReceiverKeys(ephemeralPriv BoxSecretKey, payloadKey SymmetricKey, index uint64) receiverKeys {
	_ = "STUB: not implemented"
	return *new(receiverKeys)
}

// ReceiverSymmetricKey is a symmetric key paired with an identifier.
type ReceiverSymmetricKey struct {
	// In practice these identifiers will be KBFS TLF keys.
	Key SymmetricKey
	// In practice these identifiers will be KBFS TLF pseudonyms.
	Identifier []byte
}

func (r ReceiverSymmetricKey) makeReceiverKeys(ephemeralPriv BoxSecretKey, payloadKey SymmetricKey, index uint64) receiverKeys {
	_ = "STUB: not implemented"
	// Derive a message-specific shared secret by hashing the symmetric key and
	// the ephemeral public key together. This lets us use nonces that are
	// simple counters.
	return *new(receiverKeys)
}

// should be statically impossible, if the slice above is the right length

// Unlike the box key case, the identifier is supplied by the caller rather
// than computed. (These will be KBFS TLF pseudonyms.)

func checkSigncryptReceiverCount(receiverBoxKeyCount, receiverSymmetricKeyCount int) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle possible (but unlikely) overflow when adding
// together the two sizes.

// checkEncryptReceivers does some sanity checking on the
// receivers. Check that receivers aren't sent to twice; check that
// there's at least one receiver and not too many receivers.
func checkSigncryptReceivers(receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure that each receiver only shows up in the set once.

// Make sure each key hasn't been used before.

func shuffleSigncryptReceivers(receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey) ([]receiverKeysMaker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// signcryptRNG is an interface encapsulating all the randomness
// (aside from ephemeral key generation) that happens during
// signcryption. Tests can override it to make encryption
// deterministic.
type signcryptRNG interface {
	createSymmetricKey() (*SymmetricKey, error)
	shuffleReceivers(receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey) ([]receiverKeysMaker, error)
}

// This generates the payload key, and encrypts it for all the different
// recipients of the two different types. Symmetric key recipients and DH key
// recipients use different types of identifiers, but they are the same length,
// and should both be indistinguishable from random noise.
func (sss *signcryptSealStream) init(
	receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey,
	ephemeralKeyCreator EphemeralKeyCreator, rng signcryptRNG,
) error {
	if err := checkSigncryptReceivers(receiverBoxKeys, receiverSymmetricKeys); err != nil {
		return err
	}

	receivers, err := rng.shuffleReceivers(receiverBoxKeys, receiverSymmetricKeys)
	if err != nil {
		return err
	}

	ephemeralKey, err := ephemeralKeyCreator.CreateEphemeralKey()
	if err != nil {
		return err
	}

	eh := SigncryptionHeader{
		FormatName: FormatName,
		Version:    sss.version,
		Type:       MessageTypeSigncryption,
		Ephemeral:  ephemeralKey.GetPublicKey().ToKID(),
	}
	encryptionKey, err := rng.createSymmetricKey()
	if err != nil {
		return err
	}
	sss.encryptionKey = *encryptionKey

	// Prepare the secretbox that contains the sender's public key. If the
	// sender is anonymous, use an all-zeros key, so that the anonymity bit
	// doesn't leak out.
	nonce := nonceForSenderKeySecretBox()
	if sss.signingKey == nil {
		// anonymous sender mode, all zeros
		eh.SenderSecretbox = secretbox.Seal([]byte{}, make([]byte, ed25519.PublicKeySize), (*[24]byte)(&nonce), (*[32]byte)(&sss.encryptionKey))
	} else {
		// regular sender mode, an actual key
		signingPublicKeyBytes := sss.signingKey.GetPublicKey().ToKID()
		if len(signingPublicKeyBytes) != ed25519.PublicKeySize {
			panic("unexpected signing key length, anonymity bit will leak")
		}
		eh.SenderSecretbox = secretbox.Seal([]byte{}, sss.signingKey.GetPublicKey().ToKID(), (*[24]byte)(&nonce), (*[32]byte)(&sss.encryptionKey))
	}

	// Collect all the recipient identifiers, and encrypt the payload key for
	// all of them.
	for i, r := range receivers {
		//nolint:gosec // i is a valid slice index, conversion is safe
		eh.Receivers = append(eh.Receivers, r.makeReceiverKeys(ephemeralKey, sss.encryptionKey, uint64(i)))
	}

	// Encode the header to bytes, hash it, then double encode it.
	headerBytes, err := encodeToBytes(eh)
	if err != nil {
		return err
	}
	sss.headerHash = sha512.Sum512(headerBytes)
	err = sss.encoder.Encode(headerBytes)
	if err != nil {
		return err
	}

	return nil
}

func (sss *signcryptSealStream) Close() error { _ = "STUB: not implemented"; return nil }

func newSigncryptSealStream(ciphertext io.Writer, sender SigningSecretKey, receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey, ephemeralKeyCreator EphemeralKeyCreator, rng signcryptRNG) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

type defaultSigncryptRNG struct{}

func (defaultSigncryptRNG) createSymmetricKey() (*SymmetricKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (defaultSigncryptRNG) shuffleReceivers(receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey) ([]receiverKeysMaker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSigncryptSealStream creates a stream that consumes plaintext data. It
// will write out signed and encrypted data to the io.Writer passed in as
// ciphertext. The encryption is from the specified sender, and is encrypted
// for the given receivers.
//
// ephemeralKeyCreator should be the last argument; it's the 2nd one
// to preserve the public API.
//
// If initialization succeeds, returns an io.WriteCloser that accepts
// plaintext data to be encrypted and a nil error. Otherwise, returns
// nil and the initialization error.
func NewSigncryptSealStream(ciphertext io.Writer, ephemeralKeyCreator EphemeralKeyCreator, sender SigningSecretKey, receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func signcryptSeal(plaintext []byte, sender SigningSecretKey, receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey, ephemeralKeyCreator EphemeralKeyCreator, rng signcryptRNG) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SigncryptSeal a plaintext from the given sender, for the specified
// receiver groups.  Returns a ciphertext, or an error if something
// bad happened.
//
// ephemeralKeyCreator should be the last argument; it's the 2nd one
// to preserve the public API.
func SigncryptSeal(plaintext []byte, ephemeralKeyCreator EphemeralKeyCreator, sender SigningSecretKey, receiverBoxKeys []BoxPublicKey, receiverSymmetricKeys []ReceiverSymmetricKey) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
