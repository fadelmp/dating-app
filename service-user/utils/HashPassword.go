package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// HashPassword hashes the given password using SHA-256 with a key from the environment variable
func HashPassword(password string) (string, error) {

	key := os.Getenv("HASH_KEY")

	if key == "" {
		return "", fmt.Errorf("environment variable HASH_KEY is not set")
	}

	hasher := sha256.New()
	hasher.Write([]byte(password))
	hasher.Write([]byte(key))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
