package util

import (
	"crypto/sha256"
	"fmt"
)

func EncryptText(text string) string {
	bytes := []byte(text)
	hash := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", hash)
}
