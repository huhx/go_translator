package util

import (
	"math/rand"
	"strings"
	"unicode"
)

func IsChinese(text string) bool {
	return strings.ContainsFunc(text, func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	})
}

func RandomSalt(length int) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomBytes)
}
