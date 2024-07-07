package widget

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rivo/tview"
)

const highScoreFile = ".highscore.bruh"

var (
	bestScore    int
	currentScore = 0
)

// CurrentScore returns the value of the current score.
func CurrentScore() int {
	return currentScore
}

// CurrentScore sets the current score.
func SetCurrentScore(s int) {
	currentScore = s
}

// AddToCurrentScore adds a number to the current score.
func AddToCurrentScore(s int) {
	currentScore += s
}

type Score struct{ *tview.TextView }

// NewScore returns the current score widget.
func NewScore() *Score {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	view.SetBorder(true).SetTitle(" Score ")

	view.SetText("\n 0")

	return &Score{view}
}

// Update updates the score widget to show the value of the score variable.
func (s *Score) Update() {
	s.SetText(fmt.Sprintf("\n %d", currentScore))
}

// Reset resets the current score.
func (s *Score) Reset() {
	currentScore = 0
	s.SetText(fmt.Sprintf("\n %d", currentScore))
}

type HighScore struct{ *tview.TextView }

// NewScore returns the current score widget.
func NewHighScore() *HighScore {
	titleView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	titleView.SetBorder(true).SetTitle(" Best ")

	// Load high score into memory if file exists
	file, err := os.Open(highScoreFile)
	if err != nil {
		bestScore = 0
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			bestScore, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
		}
	}

	titleView.SetText(fmt.Sprintf("\n %d", bestScore))

	return &HighScore{titleView}
}

// Update updates the high score widget.
func (s *HighScore) Update() {
	if currentScore > bestScore {
		bestScore = currentScore
		// Overwrite high score file with new score
		go func() {
			file, err := os.OpenFile(highScoreFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				panic(err)
			} else {
				defer file.Close()
			}
			_, err = file.WriteString(fmt.Sprint(currentScore))
			if err != nil {
				panic(err)
			}
		}()
	}
	s.SetText(fmt.Sprintf("\n %d", bestScore))
}
