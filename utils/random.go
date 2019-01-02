package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// Random : Random Utils (Hex)
type Random struct{}

// Hex : Random hexacode with length = numBytes * 2
func (Random) Hex(numBytes int) string {
	// new seed for random
	rand.Seed(time.Now().UnixNano())

	token := make([]byte, numBytes)
	rand.Read(token)

	// encode to hex with length = numBytes * 2
	return hex.EncodeToString(token)
}
