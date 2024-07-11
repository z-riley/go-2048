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
	a.Render()

	return &a
}

// Update updates the arena after the user makes a move in the game.
func (a *Arena) Update(dir Direction) {
	a.move(dir)
	go a.grid.save()
}

// ResetGrid resets the arena.
func (a *Arena) Reset() {
	a.grid.resetGrid()
	a.Render()
}

// move attempts executes a move in the specified direction, spawning a new tile if appropriate.
func (a *Arena) move(dir Direction) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	didMove := a.grid.move(dir, a.Render)
	if didMove {
		a.grid.spawnTile()
	}
	a.Render()
}

// Render displays the latest grid in the arena.
func (a *Arena) Render() {
	a.SetText(a.grid.string(inColour))
}

// IsLoss returns whether the game is in a losing state.
func (a *Arena) IsLoss() bool {
	return a.grid.isLoss()
}

// IsWin returns whether the game is in a winning state.
func (a *Arena) IsWin() bool {
	return a.grid.highestTile() >= 2048
}
