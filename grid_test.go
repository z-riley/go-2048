package main

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
			g.tiles[x][y] = Tile{int(math.Pow(2, exponent))}
			exponent++
		}
	}
	fmt.Println(g.Render(true))
}

func TestMove_noCombines(t *testing.T) {
	type tc struct {
		input     *Grid
		direction direction
		expected  *Grid
	}

	for _, tc := range []tc{
		{
			input: &Grid{tiles: [4][4]Tile{
				{{2}, {0}, {8}, {0}},
				{{0}, {4}, {2}, {0}},
				{{0}, {8}, {4}, {0}},
				{{0}, {0}, {2}, {16}},
			}},
			direction: dirUp,
			expected: &Grid{[4][4]Tile{
				{{2}, {4}, {8}, {16}},
				{{0}, {8}, {2}, {0}},
				{{0}, {0}, {4}, {0}},
				{{0}, {0}, {2}, {0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{2}, {0}, {8}, {0}},
				{{0}, {4}, {2}, {0}},
				{{0}, {8}, {4}, {0}},
				{{0}, {0}, {2}, {16}},
			}},
			direction: dirDown,
			expected: &Grid{[4][4]Tile{
				{{0}, {0}, {8}, {0}},
				{{0}, {0}, {2}, {0}},
				{{0}, {4}, {4}, {0}},
				{{2}, {8}, {2}, {16}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{2}, {0}, {8}, {0}},
				{{0}, {4}, {2}, {0}},
				{{0}, {8}, {4}, {0}},
				{{0}, {0}, {2}, {16}},
			}},
			direction: dirLeft,
			expected: &Grid{[4][4]Tile{
				{{2}, {8}, {0}, {0}},
				{{4}, {2}, {0}, {0}},
				{{8}, {4}, {0}, {0}},
				{{2}, {16}, {0}, {0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{2}, {0}, {8}, {0}},
				{{0}, {4}, {2}, {0}},
				{{0}, {8}, {4}, {0}},
				{{0}, {0}, {2}, {16}},
			}},
			direction: dirRight,
			expected: &Grid{[4][4]Tile{
				{{0}, {0}, {2}, {8}},
				{{0}, {0}, {4}, {2}},
				{{0}, {0}, {8}, {4}},
				{{0}, {0}, {2}, {16}},
			}},
		},
	} {
		got := tc.input.clone()
		got.Move(tc.direction)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("Expected:\n<%s>\nGot:\n<%s>", tc.expected.debug(), got.debug())
		}
	}
}

func TestMove_combineTiles(t *testing.T) {
	type tc struct {
		input     *Grid
		direction direction
		expected  *Grid
	}

	for _, tc := range []tc{
		{
			input: &Grid{tiles: [4][4]Tile{
				{{2}, {4}, {8}, {4}},
				{{2}, {0}, {16}, {4}},
				{{2}, {0}, {2}, {4}},
				{{2}, {0}, {2}, {0}},
			}},
			direction: dirUp,
			expected: &Grid{[4][4]Tile{
				{{4}, {4}, {8}, {8}},
				{{4}, {0}, {16}, {4}},
				{{0}, {0}, {4}, {0}},
				{{0}, {0}, {0}, {0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{0}, {0}, {2}, {0}},
				{{2}, {0}, {0}, {2}},
				{{4}, {2}, {0}, {2}},
				{{4}, {2}, {4}, {2}},
			}},
			direction: dirDown,
			expected: &Grid{[4][4]Tile{
				{{0}, {0}, {0}, {0}},
				{{0}, {0}, {0}, {0}},
				{{2}, {0}, {2}, {2}},
				{{8}, {4}, {4}, {4}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{16}, {8}, {4}, {4}},
				{{0}, {0}, {0}, {2}},
				{{0}, {0}, {4}, {8}},
				{{4}, {2}, {4}, {4}},
			}},
			direction: dirLeft,
			expected: &Grid{[4][4]Tile{
				{{16}, {8}, {8}, {0}},
				{{2}, {0}, {0}, {0}},
				{{4}, {8}, {0}, {0}},
				{{4}, {2}, {8}, {0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{0}, {0}, {0}, {0}},
				{{0}, {0}, {0}, {0}},
				{{0}, {0}, {2}, {0}},
				{{2}, {2}, {2}, {2}},
			}},
			direction: dirRight,
			expected: &Grid{[4][4]Tile{
				{{0}, {0}, {0}, {0}},
				{{0}, {0}, {0}, {0}},
				{{0}, {0}, {0}, {2}},
				{{0}, {0}, {4}, {4}},
			}},
		},
	} {
		got := tc.input.clone()
		got.Move(tc.direction)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("Expected:\n<%s>\nGot:\n<%s>", tc.expected.debug(), got.debug())
		}
	}
}
