package main

import (
	"fmt"
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
	tiles      [gridWidth][gridHeight]Tile
	renderFunc func(string)
}

// NewGrid initialises a new grid with a random starting arrangement.
func NewGrid(rf func(string)) *Grid {
	g := Grid{renderFunc: rf}
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

	g.renderFunc(g.Render(true))
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
// Returns whether any tiles moved from the move attempt.
func (g *Grid) Move(dir direction) bool {
	var (
		moved            = false
		prevState        = [gridWidth][gridHeight]Tile{}
		combinedThisTurn = [gridWidth][gridHeight]bool{}
	)

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
			if dir == dirDown || dir == dirRight {
				i = gridWidth - 1 - i
			}

			for j := range gridHeight {
				if dir == dirDown || dir == dirRight {
					j = gridHeight - 1 - j
				}

				tile := g.tiles[i][j]
				if tile.val == emptyTile {
					continue
				}

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
					// Similar tile exists at new location; combine tiles
					g.tiles[x][y].val += tile.val // update the new location
					g.tiles[i][j].val = emptyTile // clear the old location

					// Only allow one combination operation for each tile.
					// For example, so '2, 2, 2, 2' doesn't combine into '8', but '4, 4' like it should
					combinedThisTurn[x][y] = true

					moved = true
					continue

				} else if g.tiles[x][y].val != emptyTile {
					// Different tile exists in new location; don't move
					continue
				}

				// Destination empty; move tile
				g.tiles[x][y].val = tile.val  // populate the new location
				g.tiles[i][j].val = emptyTile // clear the old location
				moved = true

			}
		}
		// TODO: re-draw screen and delay here to show animation
	}
	return moved
}

// MoveNew moves all tiles in the specified direction, combining them if appropriate.
// Returns whether any tiles moved from the move attempt.
func (g *Grid) MoveNew(dir direction) bool {
	moved := false

	return moved
}

type vec []Tile

// JUST MOVES LEFT FOR NOW
func MoveVector(vec []Tile) []Tile {

	// For each tile:
	// Try to move 3 spaces
	// Try to move 2 spaces
	// Try to move 1 space

	// !!!!!!!!!!!!!!!!
	// IDEA: REMOVE THE ZEROS BEFORE DOING ANY OPERATIONS
	// THEN ADD THEM ON AT THE END
	// 2,2,2,2 -> 4,4. Then add the zeros on the right

	// FIXME: FIRST TRY LUT VERSION (IN TXT NOTE)

	for i := range vec {
		tile := vec[i]

		for _, n := range []int{1, 2, 3} {

			// Calculate the hypothetical next position for the tile
			newPos := i - n

			// Skip if new position is not valid (on the grid)
			if newPos < 0 || newPos >= len(vec) {
				break
			}

			// Skip if source tile is empty
			if tile.val == emptyTile {
				continue
			}

			// Combine if similar tile exists at destination
			if vec[newPos].val == tile.val {
				fmt.Println("!¬!!!!!!!!!!!!!!!!!!!!!!")
				vec[newPos].val += tile.val // update the new location
				vec[i].val = emptyTile      // clear the old location
				break
			} else if vec[newPos].val != emptyTile {
				// Break if move blocked by another tile

				break
			}

			// Destination empty; move tile
			if vec[newPos].val == emptyTile {
				vec[newPos].val = tile.val // populate the new location
				vec[i].val = emptyTile     // clear the old location
				continue
			}

			panic("THIS CODE SHOULDN'T RUN")

		}

	}

	// Clear all combinedThisTurn flags
	for i := range vec {
		vec[i].combinedThisTurn = false
	}

	return vec
}

// SpawnTile adds a new tile in a random location on the grid.
// The value of the tile is 2 (90% chance) or 4 (10% chance.)
func (g *Grid) SpawnTile() {
	val := 2
	if rand.Float64() >= 0.9 {
		val = 4
	}

	x, y := rand.Intn(gridWidth), rand.Intn(gridHeight)
	for g.tiles[x][y].val != emptyTile {
		// Try again until they're unique
		x, y = rand.Intn(gridWidth), rand.Intn(gridHeight)
	}

	g.tiles[x][y].val = val
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
