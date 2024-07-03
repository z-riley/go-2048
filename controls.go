package main

import (
	"github.com/gdamore/tcell/v2"
)

func initControls(game *Game) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			go game.ExecuteMove(dirUp)
		case tcell.KeyDown:
			go game.ExecuteMove(dirDown)
		case tcell.KeyLeft:
			go game.ExecuteMove(dirLeft)
		case tcell.KeyRight:
			go game.ExecuteMove(dirRight)
		case tcell.KeyCtrlR:
			game.Reset()
		}
		return event
	})
}
