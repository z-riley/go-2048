package main

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	game := NewGame()

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(" Score "), 7, 0, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(" Best "), 7, 0, false).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(ResetButton(), 7, 0, false),
			16, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(Title(), 6, 1, false).
			AddItem(game, gridHeight*tileHeight+4, 0, false).
			AddItem(tview.NewBox(), 0, 1, false),
			gridWidth*tileWidth+4, 0, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" How to Play "), 20, 1, false)

	initControls(game)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
