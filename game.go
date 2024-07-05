package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type game struct {
	currentScore *tview.TextView
	bestScore    tview.Primitive
	resetButton  *tview.Button
	title        tview.Primitive
	arena        *Arena
	guide        tview.Primitive
}

func (g *game) UserInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyUp:
		go g.ExecuteMove(dirUp)
	case tcell.KeyDown:
		go g.ExecuteMove(dirDown)
	case tcell.KeyLeft:
		go g.ExecuteMove(dirLeft)
	case tcell.KeyRight:
		go g.ExecuteMove(dirRight)
	case tcell.KeyCtrlR:
		g.arena.Reset()
	}
	return event
}

func (g *game) ExecuteMove(dir direction) {
	g.arena.mu.Lock()
	defer g.arena.mu.Unlock()

	// 1. Grid moves and re-renders itself as it goes
	didMove := g.arena.grid.Move(dir, g.arena.render)

	// 2. Grid spawns tile and re-renders itself
	if didMove {
		g.arena.grid.SpawnTile()
	}
	// 3. Check win/lose (todo)

	// 4. Update score
	g.currentScore.SetText(fmt.Sprintf("\n\n %d", score))

	// 5. Re-render if needed?
}
