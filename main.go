package main

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {

	game := game{
		currentScore: NewScore(),
		bestScore:    tview.NewBox().SetBorder(true).SetTitle(" Best "),
		resetButton:  ResetButton(),
		title:        Title(),
		arena:        NewArena(),
		guide:        tview.NewBox().SetBorder(true).SetTitle(" How to Play "),
	}

	app.SetInputCapture(game.UserInput)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.currentScore, 7, 0, false).
			AddItem(game.bestScore, 7, 0, false).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(game.resetButton, 7, 0, false),
			16, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.title, 6, 1, false).
			AddItem(game.arena, gridHeight*tileHeight+4, 0, false).
			AddItem(tview.NewBox(), 0, 1, false),
			gridWidth*tileWidth+4, 0, false).
		AddItem(game.guide, 20, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
