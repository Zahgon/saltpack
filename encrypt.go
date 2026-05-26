// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"bytes"
	"crypto/sha512"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

type encryptStream struct {
	version    Version
	output     io.Writer
	encoder    encoder
	payloadKey SymmetricKey
	buffer     bytes.Buffer
	headerHash headerHash
	macKeys    []macKey

	numBlocks encryptionBlockNumber // the lower 64 bits of the nonce

	err error
}

func (es *encryptStream) Write(plaintext []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If es.buffer.Len() == encryptionBlockSize, we don't want to
// write it out just yet, since for V2 we need to be sure this
// isn't the last block.

func makeEncryptionBlock(version Version, ciphertext []byte, authenticators []payloadAuthenticator, isFinal bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func checkEncryptBlockRead(version Version, isFinal bool, blockSize, plaintextLen, bufLen int) {
	_ = "STUB: not implemented"
	return
}

// We shouldn't read more than a full block's worth.

// If we read less than a full block's worth, then we
// shouldn't have anything left in the buffer.

// isFinal must be equivalent to plaintextLen being 0
// (which, by the above, implies that bufLen == 0).

// If isFinal, then plaintextLen can be any number,
// buf bufLen must be 0.

func (es *encryptStream) encryptBlock(isFinal bool) error {
	_ = "STUB: not implemented"
	// NOTE: plaintext is a slice into es.buffer's buffer, so make
	// sure not to stash it anywhere.
	return nil
}

// Compute the digest to authenticate, and authenticate it for each
// recipient.

func checkKnownVersion(version Version) error { _ = "STUB: not implemented"; return nil }

// checkEncryptReceivers does some sanity checking on the
// receivers. Check that receivers aren't sent to twice; check that
// there's at least one receiver and not too many receivers.
func checkEncryptReceivers(receivers []BoxPublicKey) error { _ = "STUB: not implemented"; return nil }

// Make sure that each receiver only shows up in the set once.

// Make sure each key hasn't been used before.

func shuffleEncryptReceivers(receivers []BoxPublicKey) ([]BoxPublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encryptRNG is an interface encapsulating all the randomness (aside
// from ephemeral key generation) that happens during
// encryption. Tests can override it to make encryption deterministic.
type encryptRNG interface {
	createSymmetricKey() (*SymmetricKey, error)
	shuffleReceivers(receivers []BoxPublicKey) ([]BoxPublicKey, error)
}

func (es *encryptStream) init(
	version Version, sender BoxSecretKey, receivers []BoxPublicKey,
	ephemeralKeyCreator EphemeralKeyCreator, rng encryptRNG,
) error {
	if err := checkKnownVersion(version); err != nil {
		return err
	}

	if err := checkEncryptReceivers(receivers); err != nil {
		return err
	}

	receivers, err := rng.shuffleReceivers(receivers)
	if err != nil {
		return err
	}

	ephemeralKey, err := ephemeralKeyCreator.CreateEphemeralKey()
	if err != nil {
		return err
	}

	// If we have a nil Sender key, then we really want the ephemeral key
	// as the main encryption key.
	if sender == nil {
		sender = ephemeralKey
	}

	eh := EncryptionHeader{
		FormatName: FormatName,
		Version:    version,
		Type:       MessageTypeEncryption,
		Ephemeral:  ephemeralKey.GetPublicKey().ToKID(),
		Receivers:  make([]receiverKeys, 0, len(receivers)),
	}
	payloadKey, err := rng.createSymmetricKey()
	if err != nil {
		return err
	}
	es.payloadKey = *payloadKey

	nonce := nonceForSenderKeySecretBox()
	eh.SenderSecretbox = secretbox.Seal([]byte{}, sender.GetPublicKey().ToKID(), (*[24]byte)(&nonce), (*[32]byte)(&es.payloadKey))

	for i, receiver := range receivers {
		sharedKey := ephemeralKey.Precompute(receiver)
		//nolint:gosec // i is a valid slice index, conversion is safe
		nonce := nonceForPayloadKeyBox(version, uint64(i))
		payloadKeyBox := sharedKey.Box(nonce, es.payloadKey[:])

		keys := receiverKeys{PayloadKeyBox: payloadKeyBox}

		// Don't specify the receivers if this public key wants to hide
		if !receiver.HideIdentity() {
			keys.ReceiverKID = receiver.ToKID()
		}

		eh.Receivers = append(eh.Receivers, keys)
	}

	// Encode the header to bytes, hash it, then double encode it.
	headerBytes, err := encodeToBytes(eh)
	if err != nil {
		return err
	}
	es.headerHash = sha512.Sum512(headerBytes)
	err = es.encoder.Encode(headerBytes)
	if err != nil {
		return err
	}

	// Use the header hash to compute the MAC keys.
	//
	// TODO: Plumb the pre-computed shared keys above through to
	// computeMACKeysSender.
	es.macKeys = computeMACKeysSender(version, sender, ephemeralKey, receivers, es.headerHash)

	return nil
}

func computeMACKeySender(version Version, index uint64, secret, eSecret BoxSecretKey, public BoxPublicKey, headerHash headerHash) macKey {
	_ = "STUB: not implemented"
	// Switch on the whole version (i.e., not just the major
	// version) since we're writing.
	return *new(macKey)
}

func computeMACKeysSender(version Version, sender, ephemeralKey BoxSecretKey, receivers []BoxPublicKey, headerHash headerHash) []macKey {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // i is a valid slice index, conversion is safe

func (es *encryptStream) Close() error { _ = "STUB: not implemented"; return nil }

func newEncryptStream(version Version, ciphertext io.Writer, sender BoxSecretKey, receivers []BoxPublicKey, ephemeralKeyCreator EphemeralKeyCreator, rng encryptRNG) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

type defaultEncryptRNG struct{}

func (defaultEncryptRNG) createSymmetricKey() (*SymmetricKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (defaultEncryptRNG) shuffleReceivers(receivers []BoxPublicKey) ([]BoxPublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// receiversToEphemeralKeyCreator retrieves the EphemeralKeyCreator
// from the first receiver; this is to preserve API behavior.
func receiversToEphemeralKeyCreator(receivers []BoxPublicKey) (EphemeralKeyCreator, error) {
	_ = "STUB: not implemented"
	return *new(EphemeralKeyCreator), nil
}

// NewEncryptStream creates a stream that consumes plaintext data.
// It will write out encrypted data to the io.Writer passed in as ciphertext.
// The encryption is from the specified sender, and is encrypted for the
// given receivers.
//
// If initialization succeeds, returns an io.WriteCloser that accepts
// plaintext data to be encrypted and a nil error. Otherwise, returns
// nil and the initialization error.
func NewEncryptStream(version Version, ciphertext io.Writer, sender BoxSecretKey, receivers []BoxPublicKey) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func seal(version Version, plaintext []byte, sender BoxSecretKey, receivers []BoxPublicKey, ephemeralKeyCreator EphemeralKeyCreator, rng encryptRNG) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Seal a plaintext from the given sender, for the specified receiver groups.
// Returns a ciphertext, or an error if something bad happened.
func Seal(version Version, plaintext []byte, sender BoxSecretKey, receivers []BoxPublicKey) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
