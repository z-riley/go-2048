package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Score struct {
	currentScore int
}

// NewScore returns the current score widget.
func NewScore() *tview.TextView {
	titleView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() })
	titleView.SetBackgroundColor(tcell.ColorBlack).SetBorder(true).SetTitle(" Score ")

	titleView.SetText("\n\n999")

	return titleView
}

// UpdateScore updates the score widget to the given number.
func (s *Score) UpdateScore(score int) {
	// TODO:
}
