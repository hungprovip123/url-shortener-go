package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

const (
	CodeLength = 6
)

func GenerateShortCode() (string, error) {
	bytes := make([]byte, CodeLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	encoded := base64.URLEncoding.EncodeToString(bytes)
	code := strings.ReplaceAll(encoded, "=", "")

	if len(code) > CodeLength {
		code = code[:CodeLength]
	}

	return code, nil
}
