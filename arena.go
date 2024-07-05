package main

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const inColour = true

type Arena struct {
	mu sync.Mutex
	*tview.TextView

	grid  *Grid
	score int
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
		mu:       sync.Mutex{},
		TextView: textView,
		grid:     &Grid{},
	}
	a.Reset()

	return &a
}

// ResetGrid resets the arena.
func (a *Arena) Reset() {
	score = 0
	a.grid.ResetGrid()
	a.render()
}

func (a *Arena) render() {
	a.SetText(a.grid.Render(inColour))
}
