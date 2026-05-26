// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package basic

import (
	"github.com/keybase/saltpack"
	"golang.org/x/crypto/ed25519"
)

// EphemeralKeyCreator creates random ephemeral keys.
type EphemeralKeyCreator struct{}

// CreateEphemeralKey creates a random ephemeral key.
func (c EphemeralKeyCreator) CreateEphemeralKey() (saltpack.BoxSecretKey, error) {
	_ = "STUB: not implemented"
	return *

	// PublicKey is a basic implementation of a saltpack public key
	new(saltpack.BoxSecretKey), nil
}

type PublicKey struct {
	EphemeralKeyCreator
	saltpack.RawBoxKey
}

// SecretKey is a basic implementation of a saltpack private key
type SecretKey struct {
	sec saltpack.RawBoxKey
	pub PublicKey
}

// PrecomputedSharedKey is a basic implementation of a saltpack
// precomputed shared key, computed from a BasicPublicKey and a BasicPrivateKey
type PrecomputedSharedKey saltpack.RawBoxKey

// ToKID takes a Publickey and returns a "key ID" or a KID, which is
// just the key itself in this implementation. It can be used to identify
// the key.
func (k PublicKey) ToKID() []byte { _ = "STUB: not implemented"; return nil }

// ToRawBoxKeyPointer returns a RawBoxKey from a given public key.
// A RawBoxKey is just a bunch of bytes that can be used in
// the lower-level Box libraries.
func (k PublicKey) ToRawBoxKeyPointer() *saltpack.RawBoxKey { _ = "STUB: not implemented"; return nil }

// HideIdentity says not to hide the identity of this key.
func (k PublicKey) HideIdentity() bool { _ = "STUB: not implemented"; return false }

func generateBoxKey() (*SecretKey, error) { _ = "STUB: not implemented"; return nil, nil }

var _ saltpack.BoxPublicKey = PublicKey{}

// Box runs the NaCl box for the given sender and receiver key.
func (k SecretKey) Box(receiver saltpack.BoxPublicKey, nonce saltpack.Nonce, msg []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Unbox runs the NaCl unbox operation on the given ciphertext and nonce,
// using the receiver as the secret key.
func (k SecretKey) Unbox(sender saltpack.BoxPublicKey, nonce saltpack.Nonce, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKey returns the public key that corresponds to this secret key.
func (k SecretKey) GetPublicKey() saltpack.BoxPublicKey {
	_ = "STUB: not implemented"

	// GetRawPublicKey returns the raw public key that corresponds to this secret key.
	return *new(saltpack.BoxPublicKey)
}

func (k SecretKey) GetRawPublicKey() *[32]byte { _ = "STUB: not implemented"; return nil }

// GetRawSecretKey returns the raw secret key.
func (k SecretKey) GetRawSecretKey() *[32]byte { _ = "STUB: not implemented"; return nil }

// Precompute computes a shared key with the passed public key.
func (k SecretKey) Precompute(peer saltpack.BoxPublicKey) saltpack.BoxPrecomputedSharedKey {
	_ = "STUB: not implemented"
	return *new(saltpack.BoxPrecomputedSharedKey)
}

// NewSecretKey makes a new SecretKey from the raw 32-byte arrays
// the represent Box public and secret keys.
func NewSecretKey(pub, sec *[32]byte) SecretKey { _ = "STUB: not implemented"; return *new(SecretKey) }

var _ saltpack.BoxSecretKey = SecretKey{}

// Box runs the box computation given a precomputed key.
func (k PrecomputedSharedKey) Box(nonce saltpack.Nonce, msg []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Unbox runs the unbox computation given a precomputed key.
func (k PrecomputedSharedKey) Unbox(nonce saltpack.Nonce, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ saltpack.BoxPrecomputedSharedKey = PrecomputedSharedKey{}

// Keyring holds signing and box secret/public keypairs.
type Keyring struct {
	EphemeralKeyCreator
	encKeys map[PublicKey]SecretKey
	sigKeys map[SigningPublicKey]SigningSecretKey
}

// NewKeyring makes an empty new basic keyring.
func NewKeyring() *Keyring { _ = "STUB: not implemented"; return nil }

// ImportBoxKey imports an existing Box key into this keyring, from a raw byte arrays,
// first the public, and then the secret key halves.
func (k *Keyring) ImportBoxKey(pub, sec *[32]byte) { _ = "STUB: not implemented"; return }

// GenerateBoxKey generates a new Box secret key and imports it into the keyring.
func (k *Keyring) GenerateBoxKey() (*SecretKey, error) { _ = "STUB: not implemented"; return nil, nil }

// GenerateSigningKey generates a signing key and import it into the keyring.
func (k *Keyring) GenerateSigningKey() (*SigningSecretKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ImportSigningKey imports the raw signing key into the keyring.
func (k *Keyring) ImportSigningKey(pub *[ed25519.PublicKeySize]byte, sec *[ed25519.PrivateKeySize]byte) {
	_ = "STUB: not implemented"
	return
}

func kidToPublicKey(kid []byte) PublicKey { _ = "STUB: not implemented"; return *new(PublicKey) }

// LookupBoxSecretKey tries to find one of the secret keys in its keyring
// given the possible key IDs. It returns the index and the key, if found, and -1
// and nil otherwise.
func (k *Keyring) LookupBoxSecretKey(kids [][]byte) (int, saltpack.BoxSecretKey) {
	_ = "STUB: not implemented"
	return 0, *new(saltpack.BoxSecretKey)
}

// LookupBoxPublicKey returns the public key that corresponds to the
// given key ID (or "kid")
func (k *Keyring) LookupBoxPublicKey(kid []byte) saltpack.BoxPublicKey {
	_ = "STUB: not implemented"
	return *new(saltpack.BoxPublicKey)
}

// GetAllBoxSecretKeys returns all secret Box keys in the keyring.
func (k *Keyring) GetAllBoxSecretKeys() []saltpack.BoxSecretKey {
	_ = "STUB: not implemented"
	return nil
}

// ImportBoxEphemeralKey takes a key ID and returns a public key
// useful for encryption/decryption.
func (k *Keyring) ImportBoxEphemeralKey(kid []byte) saltpack.BoxPublicKey {
	_ = "STUB: not implemented"
	return *new(saltpack.BoxPublicKey)
}

var _ saltpack.Keyring = (*Keyring)(nil)

// SigningPublicKey is a basic public key used for verifying signatures.
// It's just a wrapper around an array of bytes.
type SigningPublicKey saltpack.RawBoxKey

type rawSigningSecretKey [ed25519.PrivateKeySize]byte

// SigningSecretKey is a basic secret key used for creating signatures
// and also for verifying signatures. It's a wrapper around an array of bytes
// and also the corresponding public key.
type SigningSecretKey struct {
	pub SigningPublicKey
	sec rawSigningSecretKey
}

// Sign runs the NaCl signature scheme on the input message, returning
// a signature.
func (k SigningSecretKey) Sign(msg []byte) (ret []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ saltpack.SigningSecretKey = SigningSecretKey{}

// ToKID returns the key id for this signing key. It just returns
// the key itself.
func (k SigningPublicKey) ToKID() []byte {
	_ = "STUB: not implemented"

	// GetPublicKey gets the public key that corresponds to this
	// secret signing key
	return nil
}

func (k SigningSecretKey) GetPublicKey() saltpack.SigningPublicKey {
	_ = "STUB: not implemented"

	// GetRawPublicKey returns the raw public key that corresponds to this secret key.
	return *new(saltpack.SigningPublicKey)
}

func (k SigningSecretKey) GetRawPublicKey() *[ed25519.PublicKeySize]byte {
	_ = "STUB: not implemented"
	return nil
}

// GetRawSecretKey returns the raw secret key.
func (k SigningSecretKey) GetRawSecretKey() *[ed25519.PrivateKeySize]byte {
	_ = "STUB: not implemented"
	return nil
}

// Verify runs the NaCl verification routine on the given msg / sig
// input.
func (k SigningPublicKey) Verify(msg []byte, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

var _ saltpack.SigningPublicKey = SigningPublicKey{}

// NewSigningSecretKey creates a new basic signing key from byte arrays.
func NewSigningSecretKey(pub *[ed25519.PublicKeySize]byte, sec *[ed25519.PrivateKeySize]byte) SigningSecretKey {
	_ = "STUB: not implemented"
	return *new(SigningSecretKey)
}

// NewSigningPublicKey creates a new public signing key from a byte array.
func NewSigningPublicKey(pub *[ed25519.PublicKeySize]byte) SigningPublicKey {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey)
}

func kidToSigningPublicKey(kid []byte) SigningPublicKey {
	_ = "STUB: not implemented"
	return *new(SigningPublicKey)
}

// LookupSigningPublicKey turns the given key ID ("kid") into a corresponding
// signing public key.
func (k *Keyring) LookupSigningPublicKey(kid []byte) saltpack.SigningPublicKey {
	_ = "STUB: not implemented"
	return *new(saltpack.SigningPublicKey)
}

var _ saltpack.SigKeyring = (*Keyring)(nil)
