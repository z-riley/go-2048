package arena

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/zac460/go-2048/pkg/util"
	"github.com/zac460/go-2048/pkg/widget"
)

// Grid is the grid arena for the game. Position {0,0} is the top left square.
type Grid struct {
	tiles [GridWidth][GridHeight]Tile
}

// NewGrid initialises a new grid with a random starting arrangement.
func NewGrid() *Grid {
	return &Grid{}
}

// ResetGrid resets the grid to a start-of-game state, spawning two '2' tiles in random locations.
func (g *Grid) ResetGrid() {
	g.tiles = [GridWidth][GridHeight]Tile{}
	// Place two '2' tiles in two random positions
	type pos struct{ x, y int }
	tile1 := pos{rand.Intn(GridWidth), rand.Intn(GridHeight)}
	tile2 := pos{rand.Intn(GridWidth), rand.Intn(GridHeight)}
	for reflect.DeepEqual(tile1, tile2) {
		// Try again until they're unique
		tile2 = pos{rand.Intn(GridWidth), rand.Intn(GridHeight)}
	}
	g.tiles[tile1.x][tile1.y].val = 2
	g.tiles[tile2.x][tile2.y].val = 2
}

// String constructs the grid arena into a single string. Set inColour to include tview colour tags.
func (g *Grid) String(inColour bool) string {
	type tileSection int
	const (
		topSection tileSection = 0
		midSection tileSection = 1
		botSection tileSection = 2
	)

	render := "\n"

	// For every row of tiles...
	for row := 0; row < GridHeight; row++ {
		// For row of text that forms a tile...
		for x := 0; x < TileHeight; x++ {
			// Alternate between topSection, midSection, botSection
			switch tileSection(x % TileHeight) {
			case topSection, botSection:
				// Construct string of coloured space
				for col := range GridWidth {
					render += g.tiles[row][col].RenderTilePadding(inColour)
				}
			case midSection:
				// Construct string of coloured numbers with padding
				for col := range GridWidth {
					render += g.tiles[row][col].RenderTileNumber(inColour)
				}
			}
			render += "\n"
		}
	}

	render = render[:len(render)-2] // remove final coloured newline
	return render
}

// move attempts to move all tiles in the specified direction, combining them if appropriate.
// Returns true if any tiles were moved from the attempt.
func (g *Grid) move(dir Direction, renderFunc func()) bool {
	moved := false

	// Execute steps, re-rendering each time
	for {
		movedThisTurn := false
		for row := 0; row < GridHeight; row++ {
			var rowMoved bool

			// The MoveStep function only operates on a row, so to move vertically
			// we must transpose the grid before and after the move operation.
			if dir == DirUp || dir == DirDown {
				g.tiles = transpose(g.tiles)
			}

			g.tiles[row], rowMoved = moveStep(g.tiles[row], dir)

			if dir == DirUp || dir == DirDown {
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
	for i := 0; i < GridWidth; i++ {
		for j := 0; j < GridHeight; j++ {
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

	x, y := rand.Intn(GridWidth), rand.Intn(GridHeight)
	for g.tiles[x][y].val != emptyTile {
		// Try again until they're unique
		x, y = rand.Intn(GridWidth), rand.Intn(GridHeight)
	}

	g.tiles[x][y].val = val
}

// string arranges the grid into a human readable string for debugging purposes.
func (g *Grid) string() string {
	var out string
	for row := 0; row < GridHeight; row++ {
		for col := range GridWidth {
			out += g.tiles[row][col].RenderTileNumber(false) + "|"
		}
		out += "\n"
	}
	return out
}

// clone returns a deep copy for debugging purposes.
func (g *Grid) clone() *Grid {
	newGrid := &Grid{}
	for a := range GridHeight {
		for b := range GridWidth {
			newGrid.tiles[a][b] = g.tiles[a][b]
		}
	}
	return newGrid
}

// moveStep executes one part of the a move. Call multiple times until false
// is returned to complete a full move. Optional: variable to place the number of points gained by the step.
func moveStep(g [GridWidth]Tile, dir Direction) ([GridWidth]Tile, bool) {

	// Iterate in the opposite direction to the move
	reverse := false
	if dir == DirLeft || dir == DirUp {
		reverse = true
	}
	iter := util.NewIter(len(g), reverse)

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
			widget.AddToCurrentScore(g[newPos].val)
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
func transpose(matrix [GridWidth][GridHeight]Tile) [GridHeight][GridWidth]Tile {
	var transposed [GridHeight][GridWidth]Tile
	for i := 0; i < GridWidth; i++ {
		for j := 0; j < GridHeight; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}
