package utils

import (
	"crypto/rand"
	"log"
	"math/big"
	"strings"
)

func GenerateRandomId() string {
	length := 10
	keys := "1234567890abcdefghijklmnopqrstuvwxyz"

	var sb strings.Builder

	for i := 0; i < length; i++ {
		// idx := rand.Intn(len(keys))
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
		if err != nil {
			log.Fatalf("Failed to generate random index: %v", err)
		}
		sb.WriteByte(keys[idx.Int64()])
	}

	return sb.String()
}
