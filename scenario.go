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
		r := rl.NewRectangle(40, 90, 20, 20)

		if lastAnswerWasCorrect {
			raygui.LabelEx(r, fmt.Sprintf("%s!", correctAnswer), rl.Green, raygui.BackgroundColor(), raygui.BackgroundColor())
		} else {
			raygui.LabelEx(r, fmt.Sprintf("%s!", incorrectAnswer), rl.Red, raygui.BackgroundColor(), raygui.BackgroundColor())
		}
		raygui.Label(rl.NewRectangle(float32(40+rl.MeasureText(incorrectAnswer, 0xa)), 90, 20, 20), "Ready for next question ....")
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
