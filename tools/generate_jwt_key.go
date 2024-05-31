package tools

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateRandomKey generates a secure random key of the specified length.
func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

// Standalone function to generate and print a key, for testing purposes
func PrintGeneratedKey() {
	key, err := GenerateRandomKey(32) // 32 bytes key for JWT
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	fmt.Println("Generated JWT Key:", key)
}
