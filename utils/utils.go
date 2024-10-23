package utils

import "math/rand"

func GenerateRandomId() string {
	length := 10
	keys := "1234567890abcdefghijklmnopqrstuvwxyz"

	s := ""

	for i := 0; i < length; i++ {
		idx := rand.Intn(len(keys))
		s += string(keys[idx])
	}

	return s
}
