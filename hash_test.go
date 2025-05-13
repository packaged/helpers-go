package helpers

import (
	"testing"
)

func TestHashToLength(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{
			name:     "Valid input with length 10",
			input:    "test-input",
			length:   10,
			expected: "f17fbcdc7a",
		},
		{
			name:     "Length 0",
			input:    "test-input",
			length:   0,
			expected: "",
		},
		{
			name:     "Negative length",
			input:    "test-input",
			length:   -5,
			expected: "",
		},
		{
			name:     "Length greater than sha512.Size",
			input:    "test-input",
			length:   100,
			expected: "f17fbcdc7aee821520de21628749d456a3c0ca87ffbb7800a915e37e7076743914d6920459aeb89828aee6efc1488f092eae",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HashToLength(tt.input, tt.length)
			if result != tt.expected {
				t.Errorf("HashToLength(%q, %d) = %q; want %q", tt.input, tt.length, result, tt.expected)
			}
		})
	}
}
