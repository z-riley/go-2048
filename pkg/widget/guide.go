package widget

import (
	"github.com/rivo/tview"
)

type Guide struct{ *tview.TextView }

// NewGuide returns a new guide widget.
func NewGuide() *Guide {
	view := tview.NewTextView().SetDynamicColors(true)
	view.SetBorder(true)
	g := Guide{view}
	g.Reset()
	return &g
}

// Win changes the guide when the player wins the game.
func (g *Guide) Win() {
	g.SetTitle(" You Win! ")
	g.SetText("\n" + "Keep playing or restart the game.")
}

// Lose changes the guide when the player loses the game.
func (g *Guide) Lose() {
	g.SetTitle(" You Lose! ")
	g.SetText("\n" + "Press reset to try again.")
}

// Reset sets the guide to its starting state.
func (g *Guide) Reset() {
	g.SetTitle(" How to Play ")
	g.SetText("\n" + "Use the arrow keys to move the tiles." +
		"\n\n" + "Reach 2048 before the grid fills up!")
}
