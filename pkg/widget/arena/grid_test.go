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
			expected: [4]Tile{{val: 2}, {val: 2}, {val: 4, cmb: true}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 2}, {val: 2}, {val: 4, cmb: true}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4, cmb: true}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4, cmb: true}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4, cmb: true}, {val: 4, cmb: true}, {val: 0}, {val: 0}},
			moved:    true,
		},
		// 0 4 2 2 --[left]--> 4 4 0 0
		{
			input:    [4]Tile{{val: 0}, {val: 4}, {val: 2}, {val: 2}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 0}, {val: 4}, {val: 4, cmb: true}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 0}, {val: 4}, {val: 4, cmb: true}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 4}, {val: 0}, {val: 4, cmb: true}, {val: 0}},
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
		// 4 0 0 0 --[left]--> 4 0 0 0
		{
			input:    [4]Tile{{val: 4}, {val: 0}, {val: 0}, {val: 0}},
			dir:      DirLeft,
			expected: [4]Tile{{val: 4}, {val: 0}, {val: 0}, {val: 0}},
			moved:    false,
		},
		// 2 2 2 2 --[right]--> 4 4 0 0
		{
			input:    [4]Tile{{val: 2}, {val: 2}, {val: 2}, {val: 2}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 4, cmb: true}, {val: 2}, {val: 2}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 0}, {val: 4, cmb: true}, {val: 2}, {val: 2}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 4, cmb: true}, {val: 0}, {val: 4, cmb: true}},
			moved:    true,
		},
		{
			input:    [4]Tile{{val: 0}, {val: 4, cmb: true}, {val: 0}, {val: 4, cmb: true}},
			dir:      DirRight,
			expected: [4]Tile{{val: 0}, {val: 0}, {val: 4, cmb: true}, {val: 4, cmb: true}},
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

func TestX(t *testing.T) {

}
