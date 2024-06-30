package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	// Game widget
	textView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetBackgroundColor(tcell.NewHexColor(gridColour)).
		SetBorder(true).SetBackgroundColor(tcell.ColorBlack)

	grid := NewGrid()
	render := grid.Render(true)
	fmt.Fprintf(textView, "%s ", render)

	// Layout
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(" Score "), 7, 0, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(" Best "), 7, 0, false).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(ResetButton(), 7, 0, false),
			16, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(Title(), 6, 1, false).
			AddItem(textView, gridHeight*tileHeight+4, 0, false).
			AddItem(tview.NewBox(), 0, 1, false),
			gridWidth*tileWidth+4, 0, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" How to Play "), 20, 1, false)

	// Controls
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			return nil
		case tcell.KeyDown:
			return nil
		case tcell.KeyCtrlR:
			grid.ResetGrid()
			textView.SetText(grid.Render(true))
		}
		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
