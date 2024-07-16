package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zac460/go-2048/pkg/widget"
	"github.com/zac460/go-2048/pkg/widget/arena"
)

type Game struct {
	timer        *widget.Timer
	currentScore *widget.Score
	highScore    *widget.HighScore
	resetButton  *widget.ResetButton
	exitButton   *widget.ExitButton
	title        *widget.Title
	arena        *arena.Arena
	guide        *widget.Guide
}

// NewGame returns the top-level struct for the game.
func NewGame(app *tview.Application) *Game {
	g := Game{
		timer:        widget.NewTimer(app),
		currentScore: widget.NewScore(),
		highScore:    widget.NewHighScore(),
		title:        widget.NewTitle(),
		arena:        arena.NewArena(app),
		guide:        widget.NewGuide(),
	}
	g.resetButton = widget.NewResetButton(g.Reset)
	g.exitButton = widget.NewExitButton(app.Stop)
	app.SetInputCapture(g.UserInput)

	return &g
}

// UserInput is the callback given to the tview app to handle keypresses.
func (g *Game) UserInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyUp:
		go g.ExecuteMove(arena.DirUp)
		g.timer.Begin()
	case tcell.KeyDown:
		go g.ExecuteMove(arena.DirDown)
		g.timer.Begin()
	case tcell.KeyLeft:
		go g.ExecuteMove(arena.DirLeft)
		g.timer.Begin()
	case tcell.KeyRight:
		go g.ExecuteMove(arena.DirRight)
		g.timer.Begin()
	case tcell.KeyCtrlR:
		g.Reset()
	}
	return event
}

// ExecuteMove carries out a move (up, down, left, right) in the given direction.
func (g *Game) ExecuteMove(dir arena.Direction) {
	g.arena.Update(dir)

	g.currentScore.Update()
	g.highScore.Update()

	if g.arena.IsLoss() {
		g.title.Lose()
		g.guide.Lose()
		g.timer.Pause()
	}
	if g.arena.IsWin() {
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
	g.timer.Reset()
}
