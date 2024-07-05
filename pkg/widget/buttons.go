package widget

import (
	"github.com/rivo/tview"
)

type ResetButton struct {
	*tview.Button
}

// ResetButton returns the reset button widget. Provide a callback to be run
// when the button is pressed.
func NewResetButton(callback func()) *ResetButton {
	button := ResetButton{tview.NewButton("Reset")}
	button.SetSelectedFunc(callback)
	button.SetBorder(true)
	return &button
}

type ExitButton struct{ *tview.Button }

// NewExitButton returns the exit button widget.
func NewExitButton(callback func()) *ExitButton {
	button := tview.NewButton("Exit").SetSelectedFunc(callback)
	button.SetBorder(true)
	return &ExitButton{button}
}
