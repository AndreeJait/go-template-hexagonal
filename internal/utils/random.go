package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateRandomToken returns a secure random base64 string of n bytes.
func GenerateRandomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	// Base64 encode (URL safe version to avoid + and /)
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b), nil
}
