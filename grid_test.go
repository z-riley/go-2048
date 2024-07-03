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
			g.tiles[x][y] = Tile{val: int(math.Pow(2, exponent))}
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
				{{val: 2}, {val: 0}, {val: 8}, {val: 0}},
				{{val: 0}, {val: 4}, {val: 2}, {val: 0}},
				{{val: 0}, {val: 8}, {val: 4}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 16}},
			}},
			direction: dirUp,
			expected: &Grid{[4][4]Tile{
				{{val: 2}, {val: 4}, {val: 8}, {val: 16}},
				{{val: 0}, {val: 8}, {val: 2}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 4}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{val: 2}, {val: 0}, {val: 8}, {val: 0}},
				{{val: 0}, {val: 4}, {val: 2}, {val: 0}},
				{{val: 0}, {val: 8}, {val: 4}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 16}},
			}},
			direction: dirDown,
			expected: &Grid{[4][4]Tile{
				{{val: 0}, {val: 0}, {val: 8}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 0}},
				{{val: 0}, {val: 4}, {val: 4}, {val: 0}},
				{{val: 2}, {val: 8}, {val: 2}, {val: 16}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{val: 2}, {val: 0}, {val: 8}, {val: 0}},
				{{val: 0}, {val: 4}, {val: 2}, {val: 0}},
				{{val: 0}, {val: 8}, {val: 4}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 16}},
			}},
			direction: dirLeft,
			expected: &Grid{[4][4]Tile{
				{{val: 2}, {val: 8}, {val: 0}, {val: 0}},
				{{val: 4}, {val: 2}, {val: 0}, {val: 0}},
				{{val: 8}, {val: 4}, {val: 0}, {val: 0}},
				{{val: 2}, {val: 16}, {val: 0}, {val: 0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{val: 2}, {val: 0}, {val: 8}, {val: 0}},
				{{val: 0}, {val: 4}, {val: 2}, {val: 0}},
				{{val: 0}, {val: 8}, {val: 4}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 16}},
			}},
			direction: dirRight,
			expected: &Grid{[4][4]Tile{
				{{val: 0}, {val: 0}, {val: 2}, {val: 8}},
				{{val: 0}, {val: 0}, {val: 4}, {val: 2}},
				{{val: 0}, {val: 0}, {val: 8}, {val: 4}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 16}},
			}},
		},
	} {
		got := tc.input.clone()
		got.Move(tc.direction, nil)
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
				{{val: 2}, {val: 4}, {val: 8}, {val: 4}},
				{{val: 2}, {val: 0}, {val: 16}, {val: 4}},
				{{val: 2}, {val: 0}, {val: 2}, {val: 4}},
				{{val: 2}, {val: 0}, {val: 2}, {val: 0}},
			}},
			direction: dirUp,
			expected: &Grid{[4][4]Tile{
				{{val: 4}, {val: 4}, {val: 8}, {val: 8}},
				{{val: 4}, {val: 0}, {val: 16}, {val: 4}},
				{{val: 0}, {val: 0}, {val: 4}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{val: 0}, {val: 0}, {val: 2}, {val: 0}},
				{{val: 2}, {val: 0}, {val: 0}, {val: 2}},
				{{val: 4}, {val: 2}, {val: 0}, {val: 2}},
				{{val: 4}, {val: 2}, {val: 4}, {val: 2}},
			}},
			direction: dirDown,
			expected: &Grid{[4][4]Tile{
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 2}, {val: 0}, {val: 2}, {val: 2}},
				{{val: 8}, {val: 4}, {val: 4}, {val: 4}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{val: 16}, {val: 8}, {val: 4}, {val: 4}},
				{{val: 0}, {val: 0}, {val: 0}, {val: 2}},
				{{val: 0}, {val: 0}, {val: 4}, {val: 8}},
				{{val: 4}, {val: 2}, {val: 4}, {val: 4}},
			}},
			direction: dirLeft,
			expected: &Grid{[4][4]Tile{
				{{val: 16}, {val: 8}, {val: 8}, {val: 0}},
				{{val: 2}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 4}, {val: 8}, {val: 0}, {val: 0}},
				{{val: 4}, {val: 2}, {val: 8}, {val: 0}},
			}},
		},
		{
			input: &Grid{tiles: [4][4]Tile{
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 2}, {val: 0}},
				{{val: 2}, {val: 2}, {val: 2}, {val: 2}},
			}},
			direction: dirRight,
			expected: &Grid{[4][4]Tile{
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 0}, {val: 0}},
				{{val: 0}, {val: 0}, {val: 0}, {val: 2}},
				{{val: 0}, {val: 0}, {val: 4}, {val: 4}},
			}},
		},
	} {
		got := tc.input.clone()
		got.Move(tc.direction, nil)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("Expected:\n<%s>\nGot:\n<%s>", tc.expected.debug(), got.debug())
		}
	}
}

func TestMoveVector(t *testing.T) {
	type tc struct {
		input    []Tile
		expected []Tile
	}

	for n, tc := range []tc{
		{
			input:    []Tile{{val: 4}, {val: 0}, {val: 2}, {val: 0}},
			expected: []Tile{{val: 4}, {val: 2}, {val: 0}, {val: 0}},
		},
		{
			input:    []Tile{{val: 8}, {val: 4}, {val: 2}, {val: 2}},
			expected: []Tile{{val: 8}, {val: 4}, {val: 4}, {val: 0}},
		},
		{
			input:    []Tile{{val: 2}, {val: 2}, {val: 2}, {val: 2}},
			expected: []Tile{{val: 4}, {val: 4}, {val: 0}, {val: 0}},
		},
		{
			input:    []Tile{{val: 0}, {val: 4}, {val: 2}, {val: 2}},
			expected: []Tile{{val: 4}, {val: 4}, {val: 0}, {val: 0}},
		},
	} {
		got := MoveVector(tc.input)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("[%d] \nExpected:\n<%v>\nGot:\n<%v>", n, tc.expected, got)
		}
	}
}

func TestMoveStep(t *testing.T) {
	type tc struct {
		input    []Tile
		expected []Tile
	}

	for n, tc := range []tc{
		// 2 2 2 2 --[left]--> 4 4 0 0
		{
			input:    []Tile{{val: 2}, {val: 2}, {val: 2}, {val: 2}},
			expected: []Tile{{val: 2}, {val: 2}, {val: 4, cmb: true}, {val: 0}},
		},
		{
			input:    []Tile{{val: 2}, {val: 2}, {val: 4, cmb: true}, {val: 0}},
			expected: []Tile{{val: 4, cmb: true}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
		},
		{
			input:    []Tile{{val: 4, cmb: true}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
			expected: []Tile{{val: 4, cmb: true}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
		},
		// 0 4 2 2 --[left]--> 4 4 0 0
		{
			input:    []Tile{{val: 0}, {val: 4}, {val: 2}, {val: 2}},
			expected: []Tile{{val: 0}, {val: 4}, {val: 4, cmb: true}, {val: 0}},
		},
		{
			input:    []Tile{{val: 0}, {val: 4}, {val: 4, cmb: true}, {val: 0}},
			expected: []Tile{{val: 4}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
		},
		{
			input:    []Tile{{val: 4}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
			expected: []Tile{{val: 4}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
		},
	} {
		got := MoveStep(tc.input)
		if !reflect.DeepEqual(tc.expected, got) {
			t.Errorf("[%d] \nExpected:\n<%v>\nGot:\n<%v>", n, tc.expected, got)
		}
	}
}
