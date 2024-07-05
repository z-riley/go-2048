package widget

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Guide struct{ *tview.TextView }

// NewGuide returns a new guide widget.
func NewGuide() *Guide {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)
	view.SetBackgroundColor(tcell.ColorBlack).SetBorder(true).SetTitle(" How to Play ")

	view.SetText("\n\n" + "Use the arrow keys to move the tiles." +
		"\n\n" + "Try to reach 2048 before the grid fills up!")

	return &Guide{view}
}
