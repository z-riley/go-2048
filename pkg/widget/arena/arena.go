package arena

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	// Game setup
	GridWidth        = 4
	GridHeight       = 4
	gridColour int32 = 0xbbada0
	inColour         = true
)

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

type Arena struct {
	*tview.TextView

	Mu   sync.Mutex
	grid *Grid
}

// NewArena returns the game arena arena widget.
func NewArena(app *tview.Application) *Arena {
	textView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetRegions(true).SetChangedFunc(func() { app.Draw() })
	textView.SetBackgroundColor(tcell.NewHexColor(gridColour)).
		SetBorder(true).SetBackgroundColor(tcell.ColorBlack)

	a := Arena{
		TextView: textView,
		Mu:       sync.Mutex{},
		grid:     newGrid(),
	}
	a.render()

	return &a
}

// ResetGrid resets the arena.
func (a *Arena) Reset() {
	a.grid.resetGrid()
	a.render()
}

// Move attempts executes a move in the specified direction, spawning a new tile if appropriate.
func (a *Arena) Move(dir Direction) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	didMove := a.grid.move(dir, a.render)
	if didMove {
		a.grid.spawnTile()
	}
}

// IsLoss returns whether the game is lost.
func (a *Arena) IsLoss() bool {
	return a.grid.isLoss()
}

// HighestTile returns the highest tile in the arena.
func (a *Arena) HighestTile() int {
	return a.grid.highestTile()
}

// Save saves the current arena state to disk.
func (a *Arena) Save() {
	go a.grid.save()
}

// Load loads the arena state from disk.
func (a *Arena) Load() {
	a.grid.load()
}

// render generates the game areana in string format.
func (a *Arena) render() {
	a.SetText(a.grid.string(inColour))
}
