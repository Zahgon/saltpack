// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package saltpack

import (
	"crypto/sha512"

	"github.com/keybase/go-codec/codec"
)

// maxReceiverCount is the maximum number of receivers allowed
// for a single encrypted saltpack message, which is the maximum length
// of a msgpack array.
const maxReceiverCount = (1 << 32) - 1

// encryptionBlockNumber describes which block number we're at in the sequence
// of encrypted blocks. Each encrypted block of course fits into a packet.
type encryptionBlockNumber uint64

func codecHandle() *codec.MsgpackHandle { _ = "STUB: not implemented"; return nil }

func (e encryptionBlockNumber) check() error { _ = "STUB: not implemented"; return nil }

// assertEndOfStream reads from stream, and converts a nil error into
// ErrTrailingGarbage. Thus, it always returns a non-nil error. This
// should be used in a context where io.EOF is expected, and anything
// else is an error.
func assertEndOfStream(stream *msgpackStream) error { _ = "STUB: not implemented"; return nil }

type headerHash [sha512.Size]byte

func attachedSignatureInput(version Version, headerHash headerHash, payloadChunk []byte, seqno packetSeqno, isFinal bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to do.

func detachedSignatureInput(headerHash headerHash, plaintext []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func detachedSignatureInputFromHash(plaintextAndHeaderHash []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func copyEqualSize(out, in []byte) { _ = "STUB: not implemented"; return }

func copyEqualSizeStr(out []byte, in string) { _ = "STUB: not implemented"; return }

func sliceToByte24(in []byte) [24]byte { _ = "STUB: not implemented"; return nil }

func stringToByte24(in string) [24]byte { _ = "STUB: not implemented"; return nil }

func sliceToByte32(in []byte) [32]byte { _ = "STUB: not implemented"; return nil }

func sliceToByte64(in []byte) [64]byte { _ = "STUB: not implemented"; return nil }

type macKey [cryptoAuthKeyBytes]byte

type payloadHash [sha512.Size]byte

type payloadAuthenticator [cryptoAuthBytes]byte

func (pa payloadAuthenticator) Equal(other payloadAuthenticator) bool {
	_ = "STUB: not implemented"
	return false
}

func computePayloadAuthenticator(macKey macKey, payloadHash payloadHash) payloadAuthenticator {
	_ = "STUB: not implemented"
	// Equivalent to crypto_auth, but using Go's builtin HMAC. Truncates
	// SHA512, instead of calling SHA512/256, which has different IVs.
	return *new(payloadAuthenticator)
}

func computeMACKeySingle(secret BoxSecretKey, public BoxPublicKey, nonce Nonce) macKey {
	_ = "STUB: not implemented"
	return *new(macKey)
}

func sum512Truncate256(in []byte) [32]byte {
	_ = "STUB: not implemented"
	// Consistent with computePayloadAuthenticator in that it
	// truncates SHA512 instead of calling SHA512/256, which has
	// different IVs.
	return nil
}

func computePayloadHash(version Version, headerHash headerHash, nonce Nonce, ciphertext []byte, isFinal bool) payloadHash {
	_ = "STUB: not implemented"
	return *new(payloadHash)
}

// Nothing to do.

func computeSigncryptionSignatureInput(headerHash headerHash, nonce Nonce, isFinal bool, chunkPlaintext []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// This is a bit redundant, as the nonce already contains part
// of the header hash and the isFinal flag. However, we
// truncate the header hash pretty severely for the nonce, so
// it seems a bit safer to be redundant.

func hashHeader(headerBytes []byte) headerHash { _ = "STUB: not implemented"; return *new(headerHash) }

// VersionValidator is a function that takes a version and returns nil
// if it's a valid version, and an error otherwise.
type VersionValidator func(version Version) error

// CheckKnownMajorVersion returns nil if the given version has a known
// major version. You probably want to use this with NewDecryptStream,
// unless you want to restrict to specific versions only.
func CheckKnownMajorVersion(version Version) error { _ = "STUB: not implemented"; return nil }

// SingleVersionValidator returns a VersionValidator that returns nil
// if its given version is equal to desiredVersion.
func SingleVersionValidator(desiredVersion Version) VersionValidator {
	_ = "STUB: not implemented"
	return *new(VersionValidator)
}

func checkChunkState(version Version, chunkLen int, blockIndex uint64, isFinal bool) error {
	_ = "STUB: not implemented"
	return nil
}

// For V1, we derive isFinal from the chunk length, so
// if there's a mismatch, that's a bug and not a
// stream error.

// TODO: Ideally, we'd have tests exercising this case.

// assertEncodedChunkState sanity-checks some encoded chunk parameters.
func assertEncodedChunkState(version Version, encodedChunk []byte, encodingOverhead int, blockIndex uint64, isFinal bool) {
	_ = "STUB: not implemented"
	return
}

// checkDecodedChunkState sanity-checks some decoded chunk
// parameters. A returned error means there's something wrong with the
// decoded stream.
func checkDecodedChunkState(version Version, chunk []byte, seqno packetSeqno, isFinal bool) error {
	_ = "STUB: not implemented"
	// The first decoded block has seqno 1, since the header bytes
	// are decoded first.
	return nil
}
