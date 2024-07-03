package main

import (
	"github.com/gdamore/tcell/v2"
)

func initControls(game *Game) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			game.ExecuteMove(dirUp)
		case tcell.KeyDown:
			game.ExecuteMove(dirDown)
		case tcell.KeyLeft:
			game.ExecuteMove(dirLeft)
		case tcell.KeyRight:
			game.ExecuteMove(dirRight)
		case tcell.KeyCtrlR:
			game.Reset()
		}
		return event
	})
}
