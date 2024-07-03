package main

import (
	"fmt"
	"strings"
)

const (
	emptyTile = 0
)

type Tile struct {
	val int  // the value of the number on the tile
	cmb bool // flag for whether tile was combined in the current turn
}

// NewTile spawns a new tile with a starting value of 2.
func NewTile() *Tile {
	return &Tile{
		val: 2,
	}
}

// RenderTileNumber returns the middle section of a tile, which displays its value.
func (t *Tile) RenderTileNumber(colour bool) string {
	if colour {
		return t.colour() + t.paddedString()
	} else {
		return t.paddedString()
	}
}

// RenderTilePadding returns the top or bottom padding of a tile (value not displayed).
func (t *Tile) RenderTilePadding(colour bool) string {
	padding := strings.Repeat(" ", tileWidth)
	if colour {
		return t.colour() + padding
	} else {
		return padding
	}
}

// colour gets the tview colour tag, depending on the tile's value.
func (t *Tile) colour() string {
	switch t.val {
	// Fomat: "text:background"
	case emptyTile:
		return "[#7b6b61:#cdc1b4]"
	case 2:
		// TODO: set correct font colours
		return "[#7b6b61:#ece5db]"
	case 4:
		return "[#7b6b61:#ebe0ca]"
	case 8:
		return "[#7b6b61:#e8b482]"
	case 16:
		return "[#7b6b61:#e89a6c]"
	case 32:
		return "[#7b6b61:#e68266]"
	case 64:
		return "[#7b6b61:#e46747]"
	case 128:
		return "[#7b6b61:#ead17f]"
	case 256:
		return "[#7b6b61:#e8ce71]"
	case 512:
		return "[#7b6b61:#edc651]"
	case 1024:
		return "[#7b6b61:#eec744]"
	case 2048:
		return "[#7b6b61:#eec130]"
	case 4096:
		return "[#7b6b61:#ff3b3b]"
	case 8192:
		return "[#7b6b61:#ff2021]"
	default:
		return "[#fcf8ed:#ff0000]"
	}
}

// paddedString generates a padded version of the tile's value so it's centred on the tile.
func (t *Tile) paddedString() string {
	if t.val == emptyTile {
		return strings.Repeat(" ", tileWidth)
	}

	s := fmt.Sprintf("%d", t.val)
	switch len(s) {
	case 1:
		return "   " + s + "   "
	case 2:
		return "   " + s + "  "
	case 3:
		return "  " + s + "  "
	case 4:
		return "  " + s + " "
	case 5:
		return " " + s + " "
	case 6:
		return " " + s
	default:
		return s
	}
}
