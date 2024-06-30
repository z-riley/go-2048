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
				{Tile{2}, Tile{0}, Tile{8}, Tile{0}},
				{Tile{0}, Tile{4}, Tile{2}, Tile{0}},
				{Tile{0}, Tile{8}, Tile{4}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{16}},
			}},
			direction: dirUp,
			expected: &Grid{[4][4]Tile{
				{Tile{2}, Tile{4}, Tile{8}, Tile{16}},
				{Tile{0}, Tile{8}, Tile{2}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{4}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{Tile{2}, Tile{0}, Tile{8}, Tile{0}},
				{Tile{0}, Tile{4}, Tile{2}, Tile{0}},
				{Tile{0}, Tile{8}, Tile{4}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{16}},
			}},
			direction: dirDown,
			expected: &Grid{[4][4]Tile{
				{Tile{0}, Tile{0}, Tile{8}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{0}},
				{Tile{0}, Tile{4}, Tile{4}, Tile{0}},
				{Tile{2}, Tile{8}, Tile{2}, Tile{16}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{Tile{2}, Tile{0}, Tile{8}, Tile{0}},
				{Tile{0}, Tile{4}, Tile{2}, Tile{0}},
				{Tile{0}, Tile{8}, Tile{4}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{16}},
			}},
			direction: dirLeft,
			expected: &Grid{[4][4]Tile{
				{Tile{2}, Tile{8}, Tile{0}, Tile{0}},
				{Tile{4}, Tile{2}, Tile{0}, Tile{0}},
				{Tile{8}, Tile{4}, Tile{0}, Tile{0}},
				{Tile{2}, Tile{16}, Tile{0}, Tile{0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{Tile{2}, Tile{0}, Tile{8}, Tile{0}},
				{Tile{0}, Tile{4}, Tile{2}, Tile{0}},
				{Tile{0}, Tile{8}, Tile{4}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{16}},
			}},
			direction: dirRight,
			expected: &Grid{[4][4]Tile{
				{Tile{0}, Tile{0}, Tile{2}, Tile{8}},
				{Tile{0}, Tile{0}, Tile{4}, Tile{2}},
				{Tile{0}, Tile{0}, Tile{8}, Tile{4}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{16}},
			}},
		},
	} {
		got := tc.input.Move(tc.direction)
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
				{Tile{0}, Tile{4}, Tile{8}, Tile{4}},
				{Tile{0}, Tile{0}, Tile{16}, Tile{4}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{4}},
				{Tile{0}, Tile{0}, Tile{2}, Tile{0}},
			}},
			direction: dirUp,
			expected: &Grid{[4][4]Tile{
				{Tile{0}, Tile{4}, Tile{8}, Tile{8}},
				{Tile{0}, Tile{0}, Tile{16}, Tile{4}},
				{Tile{0}, Tile{0}, Tile{4}, Tile{0}},
				{Tile{0}, Tile{0}, Tile{0}, Tile{0}},
			}},
		},
	} {
		got := tc.input.Move(tc.direction)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("Expected:\n<%s>\nGot:\n<%s>", tc.expected.debug(), got.debug())
		}
	}
}
