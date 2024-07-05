package main

import (
	"github.com/rivo/tview"
)

// ResetButton returns the reset button widget.
func ResetButton() *tview.Button {
	// Reset button
	button := tview.NewButton("Reset").SetSelectedFunc(func() {
		app.Stop()
	})
	button.SetBorder(true)
	return button
}

