package main

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const inColour = true

type Arena struct {
	*tview.TextView

	mu   sync.Mutex
	grid *Grid
}

// NewArena returns the game arena arena widget.
func NewArena() *Arena {
	textView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() { app.Draw() })
	textView.SetBackgroundColor(tcell.NewHexColor(gridColour)).
		SetBorder(true).SetBackgroundColor(tcell.ColorBlack)

	a := Arena{
		TextView: textView,
		mu:       sync.Mutex{},
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

func (a *Arena) render() {
	a.SetText(a.grid.Render(inColour))
}
