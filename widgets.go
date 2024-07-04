package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Title returns the game title widget.
func Title() *tview.TextView {
	titleView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	titleView.SetBackgroundColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorBlack)
	titleView.SetText("[#bbada0]" +
		` ____    ___   _  _     ___  
|___ \  / _ \ | || |   ( _ ) 
  __) || | | || || |_  / _ \ 
 / __/ | |_| ||__   _|| (_) |
|_____| \___/    |_|   \___/ 
`)
	return titleView
}

// ResetButton returns the reset button widget.
func ResetButton() *tview.Button {
	// Reset button
	button := tview.NewButton("Reset").SetSelectedFunc(func() {
		app.Stop()
	})
	button.SetBorder(true)
	return button
}
