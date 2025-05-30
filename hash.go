package helpers

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashToLength(input string, length int) string {
	if length <= 0 {
		return ""
	}
	hash := sha512.Sum512([]byte(input))
	encoded := hex.EncodeToString(hash[:])
	return encoded[:length]
}
