package database

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

// Hash will make a brand new hashed sum of bytes which represent the orginal bytes
func Hash(b, s []byte) []byte {
	return argon2.IDKey(b, s, 3, 64 * 1024, 2, 64)
}

// Salt will return the sum of random bytes
func Salt(b int) []byte {
	alloc := make([]byte, b)
	if _, err := rand.Read(alloc); err != nil {
		return Salt(b)
	}

	return alloc
}

// MatchHash will attempt to match the hashed source against the unhashed source
func MatchHash(hash, unhash, salt []byte) bool {
	return bytes.Equal(hash, Hash(unhash, salt))
}