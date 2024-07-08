package arena

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/zac460/go-2048/pkg/util"
	"github.com/zac460/go-2048/pkg/widget"
)

const (
	saveFile      = ".grid.bruh"
	tileDelimiter = ";"
)

// Grid is the grid arena for the game. Position {0,0} is the top left square.
type Grid struct {
	tiles [GridWidth][GridHeight]Tile

	savemu *sync.Mutex // for saving grid to disk
}

// newGrid initialises a new grid.
func newGrid() *Grid {
	g := Grid{
		tiles:  [4][4]Tile{},
		savemu: &sync.Mutex{},
	}

	if err := g.load(); err != nil {
		g.resetGrid()
	} else {
		// Reset loaded grid if it's a previously lost game
		if g.isLoss() {
			g.resetGrid()
		}
	}

	return &g
}

// resetGrid resets the grid to a start-of-game state, spawning two '2' tiles in random locations.
func (g *Grid) resetGrid() {
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

	g.save()
}

// string constructs the grid arena into a single string. Set inColour to include tview colour tags.
func (g *Grid) string(inColour bool) string {
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
					render += g.tiles[row][col].renderTilePadding(inColour)
				}
			case midSection:
				// Construct string of coloured numbers with padding
				for col := range GridWidth {
					render += g.tiles[row][col].renderTileNumber(inColour)
				}
			}
			render += "\n"
		}
	}

	return render
}

// spawnTile adds a new tile in a random location on the grid.
// The value of the tile is 2 (90% chance) or 4 (10% chance).
func (g *Grid) spawnTile() {
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
		time.Sleep(20 * time.Millisecond)
	}

	// Clear all of the "combined this turn" flags
	for i := 0; i < GridWidth; i++ {
		for j := 0; j < GridHeight; j++ {
			g.tiles[i][j].cmb = false
		}
	}

	return moved
}

// moveStep executes one part of the a move. Call multiple times until false
// is returned to complete a full move. Optional: variable to place the number of points gained by the step.
func moveStep(g [GridWidth]Tile, dir Direction) ([GridWidth]Tile, bool) {

	// Iterate in the same direction as the move
	reverse := false
	if dir == DirRight || dir == DirDown {
		reverse = true
	}
	iter := util.NewIter(len(g), reverse)

	for iter.HasNext() {
		// Calculate the hypothetical next position for the tile
		i := iter.Next()
		var newPos int
		if !reverse {
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

// isLoss returns true if the grid is in a losing state (gridlocked).
func (g *Grid) isLoss() bool {
	// False if any empty spaces exist
	for i := 0; i < GridHeight; i++ {
		for j := 0; j < GridWidth; j++ {
			if g.tiles[i][j].val == emptyTile {
				return false
			}
		}
	}

	// False if any similar tiles exist next to each other
	for i := 0; i < GridHeight; i++ {
		for j := 0; j < GridWidth-1; j++ {
			if g.tiles[i][j].val == g.tiles[i][j+1].val {
				return false
			}
		}
	}
	t := transpose(g.tiles)
	for i := 0; i < GridHeight; i++ {
		for j := 0; j < GridWidth-1; j++ {
			if t[i][j].val == t[i][j+1].val {
				return false
			}
		}
	}

	return true
}

// highestTile returns the value of the highest tile on the grid.
func (g *Grid) highestTile() int {
	highest := 0
	for a := range GridHeight {
		for b := range GridWidth {
			if g.tiles[a][b].val > highest {
				highest = g.tiles[a][b].val
			}
		}
	}
	return highest
}

// save saves the grid state to the disk.
func (g *Grid) save() {
	g.savemu.Lock()
	defer g.savemu.Unlock()

	f, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Construct string to save
	var s string
	for i := range GridHeight {
		for j := range GridWidth {
			s += fmt.Sprint(g.tiles[i][j].val) + tileDelimiter
		}
	}

	f.WriteString(s)
}

// load loads the grid from the disk. Returns error if save file doesn't
// exist or is corrupt.
func (g *Grid) load() error {

	// Load high score into memory if file exists
	file, err := os.Open(saveFile)
	if err != nil {
		return err
	} else {
		defer file.Close()
	}

	b, err := os.ReadFile(saveFile)
	if err != nil {
		return err
	}

	tileStrings := strings.Split(string(b), tileDelimiter)
	tileIdx := 0
	for i := range GridHeight {
		for j := range GridWidth {
			val, err := strconv.Atoi(tileStrings[tileIdx])
			if err != nil {
				return err
			}
			g.tiles[i][j].val = val
			tileIdx++
		}
	}
	return nil
}

// wipeSave deletes the save file.
func (g *Grid) wipeSave() {
	if err := os.Remove(saveFile); err != nil {
		panic(err)
	}
}

// debug arranges the grid into a human readable debug for debugging purposes.
func (g *Grid) debug() string {
	var out string
	for row := 0; row < GridHeight; row++ {
		for col := range GridWidth {
			out += g.tiles[row][col].renderTileNumber(false) + "|"
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
