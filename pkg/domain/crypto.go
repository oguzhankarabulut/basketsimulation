package domain

import (
	"crypto/rand"
	"fmt"
)

func generateId() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	return fmt.Sprintf("%x", b), err
}
