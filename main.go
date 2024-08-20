package main

import (
	"github.com/rivo/tview"
	"github.com/z-riley/go-2048/pkg/widget/arena"
)

func main() {
	app := tview.NewApplication()
	game := NewGame(app)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.timer, 0, 1, false).
			AddItem(game.currentScore, 0, 1, false).
			AddItem(game.highScore, 0, 1, false).
			AddItem(game.resetButton, 0, 1, false).
			AddItem(game.exitButton, 0, 1, false),
			16, 0, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(game.title, 0, 1, false).
			AddItem(game.arena, arena.GridHeight*arena.TileHeight+4, 0, false),
			arena.GridWidth*arena.TileWidth+4, 0, false).
		AddItem(game.guide, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
