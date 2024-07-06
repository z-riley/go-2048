package main

import (
	"github.com/rivo/tview"
	"github.com/zac460/go-2048/pkg/widget"
	"github.com/zac460/go-2048/pkg/widget/arena"
)

func main() {
	app := tview.NewApplication()

	game := Game{
		currentScore: widget.NewScore(),
		highScore:    widget.NewHighScore(),
		title:        widget.NewTitle(),
		arena:        arena.NewArena(app),
		guide:        widget.NewGuide(),
	}
	game.resetButton = widget.NewResetButton(game.Reset)
	game.exitButton = widget.NewExitButton(app.Stop)

	app.SetInputCapture(game.UserInput)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.currentScore, 7, 0, false).
			AddItem(game.highScore, 7, 0, false).
			AddItem(game.resetButton, 7, 0, false).
			AddItem(game.exitButton, 7, 0, false),
			16, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.title, 7, 1, false).
			AddItem(game.arena, arena.GridHeight*arena.TileHeight+4, 0, false).
			AddItem(tview.NewBox(), 0, 1, false),
			arena.GridWidth*arena.TileWidth+4, 0, false).
		AddItem(game.guide, 20, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
