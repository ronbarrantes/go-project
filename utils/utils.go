package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GenerateRandomId generates a random string of a fixed length.
// The generated string consists of digits and lowercase letters.
// It returns the generated string and an error if the random bytes
// could not be generated.
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

// MakeMapToArray converts a map with any key and value types to a slice of values.
// It returns the slice of values and an error if any occurs.
func MakeMapToArray[Key comparable, Value any](m map[Key]Value) ([]Value, error) {
	if m == nil {
		return []Value{}, nil
	}

	result := make([]Value, 0, len(m))
	for _, value := range m {
		result = append(result, value)
	}
	return result, nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
