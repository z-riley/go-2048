package widget

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/zac460/go-2048/pkg/store"
)

const (
	currentScoreLabel = "currentScore"
	highScoreLabel    = "highScore"
)

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

	ss, err := store.ReadSaveState()
	if err != nil {
		panic(err)
	}
	var ok bool
	currentScore, ok = ss[currentScoreLabel].(int)
	if !ok {
		currentScore = 0
	}

	view.SetText(fmt.Sprintf("\n %d", currentScore))

	return &Score{view}
}

// Update updates the score widget to show the value of the score variable.
func (s *Score) Update() {
	s.SetText(fmt.Sprintf("\n %d", currentScore))

	// Overwrite high score file with new score
	go func() {
		store.SaveKeyVal(currentScoreLabel, currentScore)
	}()
}

// Reset resets the current score.
func (s *Score) Reset() {
	currentScore = 0
	go func() {
		store.SaveKeyVal(currentScoreLabel, currentScore)
	}()
	s.SetText(fmt.Sprintf("\n %d", currentScore))
}

type HighScore struct{ *tview.TextView }

// NewScore returns the current score widget.
func NewHighScore() *HighScore {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	view.SetBorder(true).SetTitle(" Best ")

	ss, err := store.ReadSaveState()
	if err != nil {
		panic(err)
	}
	var ok bool
	bestScore, ok = ss[highScoreLabel].(int)
	if !ok {
		bestScore = 0
	}

	view.SetText(fmt.Sprintf("\n %d", bestScore))

	return &HighScore{view}
}

// Update updates the high score widget.
func (s *HighScore) Update() {
	if currentScore > bestScore {
		bestScore = currentScore
		// Overwrite high score file with new score
		go func() {
			store.SaveKeyVal(highScoreLabel, bestScore)
		}()
	}
	s.SetText(fmt.Sprintf("\n %d", bestScore))
}
