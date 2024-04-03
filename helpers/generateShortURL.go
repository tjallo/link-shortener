package helpers

import (
	"math/rand"
	"strings"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var alphabetLen = len(alphabet)

func GenerateShortURL(urlLen int) string {
	sb := strings.Builder{}

	for range urlLen {
		sb.WriteByte(alphabet[rand.Intn(alphabetLen)])
	}

	return sb.String()
}
