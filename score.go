package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const highScoreFile = "highscore.bruh"

var score = 0
var highScore int

type Score struct{ *tview.TextView }

// NewScore returns the current score widget.
func NewScore() *Score {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() })
	view.SetBackgroundColor(tcell.ColorBlack).SetBorder(true).SetTitle(" Score ")

	view.SetText("\n\n 0")

	return &Score{view}
}

// Update updates the score widget to show the value of the score variable.
func (s *Score) Update() {
	s.SetText(fmt.Sprintf("\n\n %d", score))
}

// Reset resets the current score.
func (s *Score) Reset() {
	score = 0
	s.SetText(fmt.Sprintf("\n\n %d", score))
}

type HighScore struct{ *tview.TextView }

// NewScore returns the current score widget.
func NewHighScore() *HighScore {
	titleView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() })
	titleView.SetBackgroundColor(tcell.ColorBlack).SetBorder(true).SetTitle(" Best ")

	// Load high score into memory if file exists
	file, err := os.Open(highScoreFile)
	if err != nil {
		highScore = 0
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			highScore, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
		}
	}

	titleView.SetText(fmt.Sprintf("\n\n %d", highScore))

	return &HighScore{titleView}
}

// Update updates the high score widget.
func (s *HighScore) Update() {
	if score > highScore {
		highScore = score
		// Overwrite high score file
		go func() {
			file, err := os.Create(highScoreFile)
			if err != nil {
				panic(err)
			} else {
				defer file.Close()
			}
			_, err = file.WriteString(fmt.Sprint(score))
			if err != nil {
				panic(err)
			}
		}()
	}
	s.SetText(fmt.Sprintf("\n\n %d", highScore))
}
