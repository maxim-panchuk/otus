package hw02unpackstring

import (
	"errors"
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func TestUnpackPositive(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "b4c0", expected: "bbbb"},
		{input: "b4c0c", expected: "bbbbc"},
		{input: "b4c0c0", expected: "bbbb"},
		{input: "b2c2", expected: "bbcc"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			r, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, r)
		})
	}
}

func TestUnpackNegative(t *testing.T) {
	invalidStringList := []string{"34", "4c", "v34"}
	for _, tc := range invalidStringList {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
