package widget

import (
	"github.com/rivo/tview"
)

type Guide struct{ *tview.TextView }

// NewGuide returns a new guide widget.
func NewGuide() *Guide {
	view := tview.NewTextView().SetDynamicColors(true)
	view.SetBorder(true).SetTitle(" How to Play ")

	view.SetText("\n" + "Use the arrow keys to move the tiles." +
		"\n\n" + "Reach 2048 before the grid fills up!")

	return &Guide{view}
}
