package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// game is the top-level struct for the game.
type game struct {
	currentScore *Score
	highScore    *HighScore
	resetButton  *tview.Button
	title        *Title
	arena        *Arena
	guide        tview.Primitive
}

// UserInput is the callback given to the tview app to handle keypresses.
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
		g.Reset()
	}
	return event
}

// ExecuteMove carries out a move (up, down, left, right) in the given direction.
func (g *game) ExecuteMove(dir direction) {
	g.arena.mu.Lock()
	defer g.arena.mu.Unlock()

	// Attempt to move and spawn new tile
	didMove := g.arena.grid.Move(dir, g.arena.render)
	if didMove {
		g.arena.grid.SpawnTile()
	}

	g.updateScore()

	lose := false
	if lose {
		// TODO: check for lose
	}

	if score > 2048 {
		g.title.Win()
	}
}

// Reset resets the game.
func (g *game) Reset() {
	score = 0
	g.arena.Reset()
	g.currentScore.Reset()
}

// updateScore updates the displayed current score.
func (g *game) updateScore() {
	g.currentScore.Update()
	g.highScore.Update()
}
