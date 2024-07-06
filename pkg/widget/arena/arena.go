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
		grid:     NewGrid(),
	}
	a.Reset()

	return &a
}

// ResetGrid resets the arena.
func (a *Arena) Reset() {
	a.grid.ResetGrid()
	a.render()
}

// Move attempts executes a move in the specified direction, spawning a new tile if appropriate.
func (a *Arena) Move(dir Direction) {
	didMove := a.grid.move(dir, a.render)
	if didMove {
		a.grid.SpawnTile()
	}
}

func (a *Arena) render() {
	a.SetText(a.grid.String(inColour))
}
