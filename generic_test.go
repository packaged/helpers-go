package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	tests := []struct {
		cond   bool
		input  [2]any
		output any
	}{
		{true, [2]any{"hello", "world"}, "hello"},
		{false, [2]any{"hello", "world"}, "world"},
		{true, [2]any{1, 2}, 1},
		{false, [2]any{1, 2}, 2},
	}

	for _, test := range tests {
		result := If(test.cond, test.input[0], test.input[1])
		assert.Equal(t, test.output, result)
	}
}

func TestCoalesce(t *testing.T) {
	tests := []struct {
		input  []any
		output any
	}{
		{[]any{"", "", "hello"}, "hello"},
		{[]any{"", "world", ""}, "world"},
		{[]any{"", "", ""}, ""},
		{[]any{0, 0, 1}, 1},
		{[]any{0, 2, 1}, 2},
		{[]any{0, 0, 0}, 0},
		{[]any{nil, nil, "hello"}, "hello"},
		{[]any{nil, "world", nil}, "world"},
		{[]any{nil, "", nil}, nil},
		{[]any{nil, nil, nil}, nil},
		{[]any{nil, 2, nil}, 2},
		{[]any{nil, 0, nil}, nil},
	}

	for _, test := range tests {
		result := Coalesce(test.input...)
		assert.Equal(t, test.output, result)
	}
}
