package widget

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ResetButton struct{ *tview.Button }

// ResetButton returns the reset button widget. Provide a callback to be run
// when the button is pressed.
func NewResetButton(callback func()) *ResetButton {
	button := tview.NewButton("Reset").SetSelectedFunc(callback)
	button.SetBorder(true)
	return &ResetButton{button}
}

type ExitButton struct{ *tview.Button }

// NewExitButton returns the exit button widget.
func NewExitButton(callback func()) *ExitButton {
	button := tview.NewButton("Exit").
		SetBackgroundColorActivated(tcell.ColorDarkRed).
		SetSelectedFunc(callback)
	button.SetBorder(true)
	return &ExitButton{button}
}
