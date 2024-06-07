package util

import (
	"strings"
	"unicode"
)

func IsChinese(text string) bool {
	return strings.ContainsFunc(text, func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	})
}
