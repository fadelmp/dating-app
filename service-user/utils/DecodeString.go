package utils

import "strconv"

// DecodeBase64 decodes a base64-encoded string to its original content
func DecodeString(encodedString string) (string, error) {

	decodedString, err := strconv.Unquote(`"` + encodedString + `"`)

	if err != nil {
		return "", err
	}

	return decodedString, nil
}
