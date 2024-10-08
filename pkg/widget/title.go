package widget

import (
	"github.com/rivo/tview"
)

const title2048 = `
____    ___   _  _     ___  
|___ \  / _ \ | || |   ( _ ) 
  __) || | | || || |_  / _ \ 
 / __/ | |_| ||__   _|| (_) |
|_____| \___/    |_|   \___/ 
`

const (
	colourNormal = "[#bbada0]"
	colourWin    = "[#40ff40]"
	colourLose   = "[#ff2021]"
)

type Title struct{ *tview.TextView }

// Title returns the game title widget.
func NewTitle() *Title {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	view.SetText(colourNormal + title2048)
	return &Title{view}
}

// Win changes the appearance of the title when 2048 is reached.
func (t *Title) Win() {
	t.SetText(colourWin + title2048)
}

// Lose changes the appearance of the title when the game is lost.
func (t *Title) Lose() {
	t.SetText(colourLose + title2048)
}

// Reset sets the apperance of the title to its original state.
func (t *Title) Reset() {
	t.SetText(colourNormal + title2048)
}
