package arena

import (
	"testing"
)

func TestPaddedString(t *testing.T) {
	type tc struct {
		input    Tile
		expected string
	}

	for _, tc := range []tc{
		{
			input:    Tile{val: 2},
			expected: "   2   ",
		},
		{
			input:    Tile{val: 16},
			expected: "   16  ",
		},
		{
			input:    Tile{val: 128},
			expected: "  128  ",
		},
		{
			input:    Tile{val: 1024},
			expected: "  1024 ",
		},
		{
			input:    Tile{val: 16384},
			expected: " 16384 ",
		},
		{
			input:    Tile{val: 131072},
			expected: " 131072",
		},
	} {
		got := tc.input.paddedString()
		if tc.expected != got {
			t.Errorf("\nExpected:<%s>\nGot:<%s>", tc.expected, got)
		}
	}

}
