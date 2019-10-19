package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var index int

type Scenario struct {
	question  Question
	repeat    int32
	index     int32
	correct   int32
	maxErrors int32
}

var (
	nextQuestionTimer    *time.Timer
	lastAnswerWasCorrect bool
)

const (
	correctAnswer   = "correct"
	incorrectAnswer = "Incorrect"
)

func NewScenario(q Question, repeat, maxErrors int32) *Scenario {
	return &Scenario{
		q,
		repeat,
		0,
		0,
		maxErrors,
	}
}

func (s *Scenario) Play() bool {
	if s.index >= s.repeat {
		return false
	}

	if nextQuestionTimer == nil {
		s.question.Draw(50, 70, 100, 100)

	} else {
		c := rl.Green
		t := correctAnswer

		if lastAnswerWasCorrect == false {
			c = rl.Red
			t = incorrectAnswer
		}

		r := rl.NewRectangle(40, 90, 20, 20)
		o := rl.MeasureText(t, rl.GetFontDefault().BaseSize)

		raygui.LabelEx(r, fmt.Sprintf("%s!", t), c, raygui.BackgroundColor(), raygui.BackgroundColor())

		r.X = r.X + 10 + float32(o)
		raygui.Label(r, "Ready for next question ....")
	}

	if a := s.question.IsAnswerCorrect(); a != nil {
		if *a == true || s.question.Tries() >= s.maxErrors {
			lastAnswerWasCorrect = false

			if *a == true {
				s.correct = s.correct + 1
				lastAnswerWasCorrect = true
			}

			s.index = s.index + 1
			s.question.Reset()

			nextQuestionTimer = time.AfterFunc(2*time.Second, func() {
				nextQuestionTimer = nil
			})
		}
	}

	return true
}

func (s *Scenario) Repeats() (total int32, correct int32) {
	total = s.index
	correct = s.correct

	return
}
