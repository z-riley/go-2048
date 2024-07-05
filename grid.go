package main

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/zac460/go-2048/pkg/iter"
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
	return &Grid{}
}

// ResetGrid resets the grid to a start-of-game state, spawning two '2' tiles in random locations.
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
					render += g.tiles[row][col].RenderTilePadding(inColour)
				}
			case midSection:
				// Construct string of coloured numbers with padding
				for col := range gridWidth {
					render += g.tiles[row][col].RenderTileNumber(inColour)
				}
			}
			render += "\n"
		}
	}

	render = render[:len(render)-2] // remove final coloured newline
	return render
}

// Move attempts to move all tiles in the specified direction, combining them if appropriate.
// Returns true if any tiles were moved from the attempt.
func (g *Grid) Move(dir direction, renderFunc func()) bool {
	moved := false

	// Execute steps, re-rendering each time
	for {
		movedThisTurn := false
		for row := 0; row < gridHeight; row++ {
			var rowMoved bool

			// The MoveStep function only operates on a row, so to move vertically
			// we must transpose the grid before and after the move operation.
			if dir == dirUp || dir == dirDown {
				g.tiles = transpose(g.tiles)
			}

			g.tiles[row], rowMoved = moveStep(g.tiles[row], dir)

			if dir == dirUp || dir == dirDown {
				g.tiles = transpose(g.tiles)
			}

			if rowMoved {
				movedThisTurn = true
				moved = true
			}
		}
		renderFunc()
		if !movedThisTurn {
			break
		}
		time.Sleep(25 * time.Millisecond)
	}

	// Clear all of the "combined this turn" flags
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			g.tiles[i][j].cmb = false
		}
	}

	return moved
}

// SpawnTile adds a new tile in a random location on the grid.
// The value of the tile is 2 (90% chance) or 4 (10% chance).
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

// string arranges the grid into a human readable string for debugging purposes.
func (g *Grid) string() string {
	var out string
	for row := 0; row < gridHeight; row++ {
		for col := range gridWidth {
			out += g.tiles[row][col].RenderTileNumber(false) + "|"
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

// moveStep executes one part of the a move. Call multiple times until false
// is returned to complete a full move. Optional: variable to place the number of points gained by the step.
func moveStep(g [gridWidth]Tile, dir direction) ([gridWidth]Tile, bool) {

	// Iterate in the opposite direction to the move
	reverse := false
	if dir == dirLeft || dir == dirUp {
		reverse = true
	}
	iter := iter.NewIter(len(g), reverse)

	for iter.HasNext() {
		// Calculate the hypothetical next position for the tile
		i := iter.Next()
		var newPos int
		if reverse {
			newPos = i - 1
		} else {
			newPos = i + 1
		}

		// Skip if new position is not valid (on the grid)
		if newPos < 0 || newPos >= len(g) {
			continue
		}

		// Skip if source tile is empty
		if g[i].val == emptyTile {
			continue
		}

		// Combine if similar tile exists at destination and end turn
		alreadyCombined := g[i].cmb || g[newPos].cmb
		if g[newPos].val == g[i].val && !alreadyCombined {
			g[newPos].val += g[i].val // update the new location
			g[newPos].cmb = true
			score += g[newPos].val
			g[i].val = emptyTile // clear the old location
			return g, true

		} else if g[newPos].val != emptyTile {
			// Move blocked by another tile
			continue
		}

		// Destination empty; move tile and end turn
		if g[newPos].val == emptyTile {
			g[newPos] = g[i]
			g[i] = Tile{}
			return g, true
		}
	}
	return g, false
}

// transpose returns a tranposed version of the grid.
func transpose(matrix [gridWidth][gridHeight]Tile) [gridHeight][gridWidth]Tile {
	var transposed [gridHeight][gridWidth]Tile
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}
