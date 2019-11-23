package questions

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxTimeout = 30 * time.Second
)

type TimerDecorator struct {
	component  Question
	timer      *time.Ticker
	elapsed    time.Duration
	maxTimeout time.Duration
}

func NewTimerDecorator(component Question) Question {
	t := &TimerDecorator{
		component,
		time.NewTicker(1 * time.Second),
		0,
		MaxTimeout,
	}

	go func() {
		for _ = range t.timer.C {
			t.elapsed = t.elapsed + 1*time.Second
		}
	}()

	return t
}

func (t *TimerDecorator) IsAnswerCorrect() *bool {
	if t.elapsed >= t.maxTimeout {
		return &False
	} else {
		return t.component.IsAnswerCorrect()
	}
}

func (t *TimerDecorator) Draw(x, y, w, h float32) {
	t.component.Draw(x, y, w, h)

	ratio := float64(t.maxTimeout-t.elapsed) / float64(t.maxTimeout)
	height := int32(40 * ratio)
	top := 40 - height

	rl.DrawRectangle(int32(x-20), int32(y+20)+top, 5, height, rl.Red)
}

func (t *TimerDecorator) Reset() {
	t.elapsed = 0
	t.component.Reset()
}

func (t *TimerDecorator) Tries() int32 {
	if t.elapsed >= t.maxTimeout {
		return -1
	} else {
		return t.component.Tries()
	}
}
