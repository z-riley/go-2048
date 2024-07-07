package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/zac460/go-2048/pkg/widget"
	"github.com/zac460/go-2048/pkg/widget/arena"
)

// Game is the top-level struct for the Game.
type Game struct {
	currentScore *widget.Score
	highScore    *widget.HighScore
	resetButton  *widget.ResetButton
	exitButton   *widget.ExitButton
	title        *widget.Title
	arena        *arena.Arena
	guide        *widget.Guide
}

// UserInput is the callback given to the tview app to handle keypresses.
func (g *Game) UserInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyUp:
		go g.ExecuteMove(arena.DirUp)
	case tcell.KeyDown:
		go g.ExecuteMove(arena.DirDown)
	case tcell.KeyLeft:
		go g.ExecuteMove(arena.DirLeft)
	case tcell.KeyRight:
		go g.ExecuteMove(arena.DirRight)
	case tcell.KeyCtrlR:
		g.Reset()
	}
	return event
}

// ExecuteMove carries out a move (up, down, left, right) in the given direction.
func (g *Game) ExecuteMove(dir arena.Direction) {
	// Mutate arena
	g.arena.Move(dir)
	g.arena.Save()

	// Update scores
	g.currentScore.Update()
	g.highScore.Update()

	// Check game state
	if g.arena.IsLoss() {
		g.title.Lose()
		g.guide.Lose()

	}
	if g.arena.HighestTile() >= 2048 {
		g.title.Win()
		g.guide.Win()
	}
}

// Reset resets the game.
func (g *Game) Reset() {
	widget.SetCurrentScore(0)
	g.arena.Reset()
	g.currentScore.Reset()
	g.title.Reset()
	g.guide.Reset()
}
