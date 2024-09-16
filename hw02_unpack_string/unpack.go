package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return s, nil
	}
	var builder strings.Builder
	var prev rune

	for _, r := range s {
		if unicode.IsDigit(r) {
			if unicode.IsDigit(prev) || prev == 0 {
				return "", ErrInvalidString
			}
			if r != '0' {
				builder.WriteString(
					strings.Repeat(string(prev), int(r-'0')))
			}
			prev = 0
			continue
		}

		if prev != 0 {
			builder.WriteRune(prev)
		}
		prev = r
	}

	if prev != 0 {
		builder.WriteRune(prev)
	}

	return builder.String(), nil
}
