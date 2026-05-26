package saltpack

const nonceBytes = 24

// Nonce is a NaCl-style nonce, with 24 bytes of data, some of which can be
// counter values, and some of which can be random-ish values.
type Nonce [nonceBytes]byte

func nonceForSenderKeySecretBox() Nonce { _ = "STUB: not implemented"; return *new(Nonce) }

func nonceForPayloadKeyBoxV2(recip uint64) Nonce { _ = "STUB: not implemented"; return *new(Nonce) }

func nonceForPayloadKeyBox(version Version, recip uint64) Nonce {
	_ = "STUB: not implemented"
	// Switch on the major version since this is called during
	// both writing and reading, and in the latter we may
	// encounter headers written by unknown minor versions.
	return *new(Nonce)
}

// Let caller be responsible for filtering out unknown
// versions.

func nonceForDerivedSharedKey() Nonce { _ = "STUB: not implemented"; return *new(Nonce) }

func nonceForMACKeyBoxV1(headerHash headerHash) Nonce {
	_ = "STUB: not implemented"
	return *new(Nonce)
}

func nonceForMACKeyBoxV2(headerHash headerHash, ephemeral bool, recip uint64) Nonce {
	_ = "STUB: not implemented"
	return *new(Nonce)
}

// Set LSB of last byte based on ephemeral.

// Construct the nonce for the ith block of encryption payload.
func nonceForChunkSecretBox(i encryptionBlockNumber) Nonce {
	_ = "STUB: not implemented"
	return *new(Nonce)
}

// Construct the nonce for the ith block of signcryption
// payload.
func nonceForChunkSigncryption(headerHash headerHash, isFinal bool, i encryptionBlockNumber) Nonce {
	_ = "STUB: not implemented"
	return *new(Nonce)
}

// Set LSB of last byte based on isFinal.

// sigNonce is a nonce for signatures.
type sigNonce [16]byte

// newSigNonce creates a sigNonce with random bytes.
func newSigNonce() (sigNonce, error) { _ = "STUB: not implemented"; return *new(sigNonce), nil }
