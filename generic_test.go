package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
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
		result := Ternary(test.cond, test.input[0], test.input[1])
		assert.Equal(t, test.output, result)
	}
}
