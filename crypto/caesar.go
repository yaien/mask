package crypto

import (
	"fmt"
	"strings"
)

var dictionary = " abcdefghijklmnopqrstuvxyz1234567890"

// Encrypt apply caesar encryption to a message
func Encrypt(message string, offset int) (string, error) {
	builder := &strings.Builder{}

	for _, char := range strings.ToLower(message) {
		index := strings.IndexRune(dictionary, char)

		if index == -1 {
			return "", fmt.Errorf("Invalid Char: '%c'", char)
		}

		index += offset

		if index >= len(dictionary) {
			index -= len(dictionary)
		}

		builder.WriteByte(dictionary[index])

	}

	return builder.String(), nil
}

func Decrypt(message string, offset int) (string, error) {
	builder := &strings.Builder{}

	for _, char := range strings.ToLower(message) {
		index := strings.IndexRune(dictionary, char)

		if index == -1 {
			return "", fmt.Errorf("Invalid Char: '%c'", char)
		}

		index -= offset

		if index < 0 {
			index += len(dictionary)
		}

		builder.WriteByte(dictionary[index])

	}

	return builder.String(), nil
}
