package backend

import (
	"encoding/hex"
	"crypto/rand"
)

// source: http://www.ashishbanerjee.com/home/go/go-generate-uuid
// API doc: https://golang.org/pkg/math/rand/#Read
func generateKey() string {
	uuid := make([]byte, 16)
	rand.Read(uuid)

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid)
}