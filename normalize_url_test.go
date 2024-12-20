package main

import (
	"testing"
)

func TestNormalizaUrl(t *testing.T) {
	cases := []struct {
		name     string
		inputUrl string
		expected string
	}{
		{
			name:     "remove scheme",
			inputUrl: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		}, {
			name:     "remove scheme",
			inputUrl: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizarUrl(tc.inputUrl)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v | Got: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
