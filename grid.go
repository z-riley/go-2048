package main

import (
	"math/rand"
	"reflect"
)

type direction int

const (
	// Game setup
	gridWidth        = 4
	gridHeight       = 4
	tileWidth        = 7
	tileHeight       = 3
	gridColour int32 = 0xbbada0
)

const (
	// Direction enumerations
	dirUp direction = iota
	dirDown
	dirLeft
	dirRight
)

// Grid is the grid arena for the game. Position {0,0} is the top left square.
type Grid struct {
	tiles [gridWidth][gridHeight]Tile
}

// NewGrid initialises a new grid with a random starting arrangement.
func NewGrid() *Grid {
	g := Grid{}
	g.ResetGrid()
	return &g
}

// ResetGrid resets the grid to a start-of-game state.
func (g *Grid) ResetGrid() {
	g.tiles = [gridWidth][gridHeight]Tile{}
	// Place two '2' tiles in two random positions
	type pos struct{ x, y int }
	tile1 := pos{rand.Intn(gridWidth), rand.Intn(gridHeight)}
	tile2 := pos{rand.Intn(gridWidth), rand.Intn(gridHeight)}
	for reflect.DeepEqual(tile1, tile2) {
		// Try again until they're unique
		tile2 = pos{rand.Intn(gridWidth), rand.Intn(gridHeight)}
	}
	g.tiles[tile1.x][tile1.y].val = 2
	g.tiles[tile2.x][tile2.y].val = 2
}

// Render constructs the grid arena into a single string.
func (g *Grid) Render(inColour bool) string {
	type tileSection int
	const (
		topSection tileSection = 0
		midSection tileSection = 1
		botSection tileSection = 2
	)

	render := "\n"

	// For every row of tiles...
	for row := 0; row < gridHeight; row++ {
		// For row of text that forms a tile...
		for x := 0; x < tileHeight; x++ {
			// Alternate between topSection, midSection, botSection
			switch tileSection(x % tileHeight) {
			case topSection, botSection:
				// Construct string of coloured space
				for col := range gridWidth {
					render += g.tiles[row][col].TilePadding(inColour)
				}
			case midSection:
				// Construct string of coloured numbers with padding
				for col := range gridWidth {
					render += g.tiles[row][col].TileNumber(inColour)
				}
			}
			render += "\n"
		}
	}

	render = render[:len(render)-2] // remove final coloured newline
	return render
}

// Move moves all tiles in the specified direction, combining them if appropriate.
func (g *Grid) Move(dir direction) {
	var prevState = [gridWidth][gridHeight]Tile{}
	var combinedThisTurn = [gridWidth][gridHeight]bool{}

	// Repeat until no more moves occur
	for !reflect.DeepEqual(g.tiles, prevState) {
		// Save grid state for comparison
		for a := range gridHeight {
			for b := range gridWidth {
				prevState[a][b] = g.tiles[a][b]
			}
		}

		// For each position where there is a tile...
		for i := range gridWidth {
			for j := range gridHeight {
				tile := g.tiles[i][j]
				if tile.val != 0 {
					// Calculate the hypothetical next position for the tile
					x, y := 0, 0
					switch dir {
					case dirLeft:
						x, y = i, j-1
					case dirRight:
						x, y = i, j+1
					case dirUp:
						x, y = i-1, j
					case dirDown:
						x, y = i+1, j
					}

					// Check if new position is valid (on the grid)
					if x < 0 || y < 0 || y >= gridHeight || x >= gridWidth {
						continue
					}

					if g.tiles[x][y].val == tile.val && !combinedThisTurn[x][y] {
						// Similar tile exists at new location. Combine tiles

						g.tiles[x][y].val += tile.val // update the new location
						g.tiles[i][j].val = emptyTile // clear the old location

						// Only allow one combination operation for each tile.
						// For example, so '2, 2, 2, 2' doesn't combine into '8', but '4, 4' like it should
						combinedThisTurn[x][y] = true

						continue

					} else if g.tiles[x][y].val != emptyTile {
						// Different tile exists in new location. Don't move
						continue
					}

					g.tiles[x][y].val = tile.val  // populate the new location
					g.tiles[i][j].val = emptyTile // clear the old location
				}
			}
		}
		// TODO: re-draw screen and delay here to show animation
	}
}

// debug arranges the grid into a human readable string for debugging purposes.
func (g *Grid) debug() string {
	var out string
	for row := 0; row < gridHeight; row++ {
		for col := range gridWidth {
			out += g.tiles[row][col].TileNumber(false) + "|"
		}
		out += "\n"
	}
	return out
}

// clone returns a deep copy for debugging purposes.
func (g *Grid) clone() *Grid {
	newGrid := &Grid{}
	for a := range gridHeight {
		for b := range gridWidth {
			newGrid.tiles[a][b] = g.tiles[a][b]
		}
	}
	return newGrid
}
