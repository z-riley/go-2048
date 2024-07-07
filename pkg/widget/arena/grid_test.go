package arena

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	g := Grid{}
	exponent := 1.0
	for x := range g.tiles {
		for y := range g.tiles {
			g.tiles[x][y] = Tile{val: int(math.Pow(2, exponent))}
			exponent++
		}
	}
	fmt.Println(g.string(true))
}

func TestMove(t *testing.T) {
	type tc struct {
		input    Grid
		dir      Direction
		expected Grid
	}

	for n, tc := range []tc{
		{
			input: Grid{
				tiles: [4][4]Tile{
					{{val: 0}, {val: 2}, {val: 2}, {val: 2}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				},
			},
			dir: DirRight,
			expected: Grid{
				tiles: [4][4]Tile{
					{{val: 0}, {val: 0}, {val: 2}, {val: 4}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				},
			},
		},
	} {
		got := tc.input.clone()
		got.move(DirRight, func() {})
		if !reflect.DeepEqual(tc.expected.tiles, got.tiles) {
			t.Errorf("[%d] \nExpected:\n<%v>\nGot:\n<%v>", n, tc.expected.debug(), got.debug())
		}
	}
}

func TestMoveStep(t *testing.T) {
	type tc struct {
		input    [4]Tile
		dir      Direction
		expected [4]Tile
		moved    bool
	}

	for n, tc := range []tc{
		// 2 2 2 2 --[left]--> 4 4 0 0
		{
			input:    [4]Tile{{val: 2}, {val: 2}, {val: 2}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4, cmb: true}, {val: 0}, {val: 2}, {val: 2}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4, cmb: true}, {val: 0}, {val: 2}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4, cmb: true}, {val: 2}, {val: 0}, {val: 2}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4, cmb: true}, {val: 2}, {val: 0}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4, cmb: true}, {val: 2}, {val: 2}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4, cmb: true}, {val: 2}, {val: 2}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4, cmb: true}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
			moved:    true,
		},
		// 0 4 2 2 --[left]--> 4 4 0 0
		{
			input:    [4]Tile{{val: 0}, {val: 4}, {val: 2}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 0}, {val: 2}, {val: 2}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4}, {val: 0}, {val: 2}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 2}, {val: 0}, {val: 2}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4}, {val: 2}, {val: 0}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 2}, {val: 2}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4}, {val: 2}, {val: 2}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
			moved:    false,
		},
		// // 2 2 2 2 --[right]--> 4 4 0 0
		{
			input:    [4]Tile{{val: 2}, {val: 2}, {val: 2}, {val: 2}},
			dir:      DirRight,
			expected: [4]Tile{{val: 2}, {val: 2}, {val: 0}, {val: 4, cmb: true}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 2}, {val: 2}, {val: 0}, {val: 4, cmb: true}},
			dir:      DirRight,
			expected: [4]Tile{{val: 2}, {val: 0}, {val: 2}, {val: 4, cmb: true}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 2}, {val: 0}, {val: 2}, {val: 4, cmb: true}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 2}, {val: 2}, {val: 4, cmb: true}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 0}, {val: 2}, {val: 2}, {val: 4, cmb: true}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 0}, {val: 4, cmb: true}, {val: 4, cmb: true}},
			moved:    true,
		},
		// // 0 2 2 2 --[right]--> 0 0 2 4
		{
			input:    [4]Tile{{val: 0}, {val: 2}, {val: 2}, {val: 2}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 2}, {val: 0}, {val: 4, cmb: true}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 0}, {val: 2}, {val: 0}, {val: 4, cmb: true}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 0}, {val: 2}, {val: 4, cmb: true}},
			moved:    true,
		},
	} {
		got, moved := moveStep(tc.input, tc.dir)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("[%d] \nExpected:\n<%v>\nGot:\n<%v>", n, tc.expected, got)
		}
		if tc.moved != moved {
			t.Errorf("[%d] \nExpected:\n<%v>\nGot:\n<%v>", n, tc.moved, moved)
		}
	}
}

func TestTranspose(t *testing.T) {
	input := [4][4]Tile{
		{{val: 1}, {val: 2}, {val: 3}, {val: 4}},
		{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
		{{val: 6}, {val: 0}, {val: 0}, {val: 0}},
		{{val: 0}, {val: 0}, {val: 0}, {val: 5}},
	}
	expected := [4][4]Tile{
		{{val: 1}, {val: 0}, {val: 6}, {val: 0}},
		{{val: 2}, {val: 0}, {val: 0}, {val: 0}},
		{{val: 3}, {val: 0}, {val: 0}, {val: 0}},
		{{val: 4}, {val: 0}, {val: 0}, {val: 5}},
	}
	got := transpose(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\nExpected:\n<%v>\nGot:\n<%v>", expected, got)
	}
}

func TestIsLoss(t *testing.T) {
	type tc struct {
		input    Grid
		expected bool
	}

	for _, tc := range []tc{
		{
			input: Grid{
				tiles: [4][4]Tile{
					{{val: 2}, {val: 0}, {val: 8}, {val: 0}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
					{{val: 0}, {val: 4}, {val: 0}, {val: 0}},
					{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				},
			},
			expected: false,
		},
		{
			input: Grid{
				tiles: [4][4]Tile{
					{{val: 4}, {val: 4}, {val: 2}, {val: 4}},
					{{val: 4}, {val: 2}, {val: 4}, {val: 2}},
					{{val: 2}, {val: 4}, {val: 2}, {val: 4}},
					{{val: 4}, {val: 2}, {val: 4}, {val: 2}},
				},
			},
			expected: false,
		},
		{
			input: Grid{
				tiles: [4][4]Tile{
					{{val: 2}, {val: 4}, {val: 2}, {val: 4}},
					{{val: 4}, {val: 2}, {val: 4}, {val: 2}},
					{{val: 2}, {val: 4}, {val: 2}, {val: 4}},
					{{val: 4}, {val: 2}, {val: 4}, {val: 2}},
				},
			},
			expected: true,
		},
		{
			input: Grid{
				tiles: [4][4]Tile{
					{{val: 2}, {val: 4}, {val: 16}, {val: 2}},
					{{val: 8}, {val: 32}, {val: 64}, {val: 16}},
					{{val: 4}, {val: 16}, {val: 8}, {val: 4}},
					{{val: 2}, {val: 8}, {val: 4}, {val: 2}},
				},
			},
			expected: true,
		},
		{
			input: Grid{
				tiles: [4][4]Tile{
					{{val: 4}, {val: 16}, {val: 4}, {val: 2}},
					{{val: 2}, {val: 32}, {val: 4}, {val: 2}},
					{{val: 4}, {val: 8}, {val: 4}, {val: 2}},
					{{val: 2}, {val: 8}, {val: 8}, {val: 8}},
				},
			},
			expected: false,
		},
	} {
		actual := tc.input.isLoss()
		if tc.expected != actual {
			t.Errorf("Expected:\n<%v>\nGot:\n<%v>", tc.expected, actual)
		}
	}
}
