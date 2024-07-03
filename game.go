package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Game struct {
	*tview.TextView
	grid *Grid
}

// Game returns the game arena widget.
func NewGame() *Game {
	textView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() { app.Draw() })
	textView.SetBackgroundColor(tcell.NewHexColor(gridColour)).
		SetBorder(true).SetBackgroundColor(tcell.ColorBlack)

	g := Game{textView, NewGrid(func(s string) { textView.SetText(s) })}
	return &g
}

func (g *Game) ExecuteMove(dir direction) {
	// 1. Grid moves and re-renders itself
	didMove := g.grid.Move(dir)
	g.grid.renderFunc(g.grid.Render(true))

	// 2. Grid spawns tile and re-renders itself
	if didMove {
		g.grid.SpawnTile()
	}

	// 3. Check win/lose (todo)

	// 4. Update score (todo)

	// 5. Re-render widgets
	g.grid.renderFunc(g.grid.Render(true))
}

// ResetGrid resets the game.
func (g *Game) Reset() {
	g.grid.ResetGrid()
}
