package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func GenerateRandomId() (string, error) {
	length := 10
	keys := "1234567890abcdefghijklmnopqrstuvwxyz"
	keyLen := byte(len(keys))

	var sb strings.Builder
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to generate random bytes: %v", err)
	}

	for _, b := range randomBytes {
		sb.WriteByte(keys[b%keyLen])
	}

	return sb.String(), nil
}
