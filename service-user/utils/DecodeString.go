package utils

import (
	"encoding/base64"
)

// DecodeBase64 decodes a base64-encoded string to its original content
func DecodeString(encodedPass string) (string, error) {

	// Decode the Base64 encoded string
	decoded, err := base64.StdEncoding.DecodeString(encodedPass)

	if err != nil {
		return "", err
	}

	// Convert the decoded bytes to a string
	return string(decoded), nil
}
