package widget

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/z-riley/go-2048/pkg/store"
)

const currentTimeLabel = "currentTime"

var playTime time.Duration

type Timer struct {
	*tview.TextView

	ticker *time.Ticker
	done   chan bool
	paused bool
}

// NewTimer returns the timer widget.
func NewTimer(app *tview.Application) *Timer {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetChangedFunc(func() { app.Draw() })
	view.SetBorder(true).SetTitle(" Timer ")

	ss, err := store.ReadSaveState()
	if err != nil {
		panic(err)
	}
	intTime, ok := ss[currentTimeLabel].(int)
	if ok {
		playTime = time.Duration(intTime)
	}

	view.SetText(fmt.Sprintf("\n %s", format(playTime)))

	return &Timer{
		TextView: view,
		done:     make(chan bool),
		ticker:   time.NewTicker(1 * time.Second),
	}
}

// Begin starts the timer widget. If the timer has already started, nothing will happen.
func (t *Timer) Begin() {
	t.paused = false
	go func() {
		for {
			select {
			case <-t.done:
				return
			case <-t.ticker.C:
				if t.paused {
					return
				}
				// Display new time
				playTime += 1 * time.Second
				t.SetText("\n" + format(playTime))

				// Save time to disk
				go func() {
					err := store.SaveKeyVal(currentTimeLabel, int(playTime))
					if err != nil {
						panic(err)
					}
				}()
			}
		}
	}()
}

// Reset resets the timer widget, which remains paused until Begin() is called.
func (t *Timer) Reset() {
	playTime = 0
	t.paused = true
	t.SetText("\n" + format(playTime))
}

// Pause pauses the timer widget.
func (t *Timer) Pause() {
	t.paused = true
}

// format formats a duration into the format "HH:MM:SS".
func format(t time.Duration) string {
	hours := int(t.Hours())
	minutes := int(t.Minutes()) % 60
	seconds := int(t.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
