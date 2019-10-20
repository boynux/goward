package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var index int

type order int

const (
	Random order = iota
	Ordered
)

const (
	NextQuestionDelay = 1 * time.Second
)

type Scenario struct {
	question       []Question
	repeat         int32
	index          int32
	correct        int32
	maxErrors      int32
	activeQuestion Question
	order          order
}

var (
	nextQuestionTimer    *time.Timer
	lastAnswerWasCorrect bool
)

const (
	CorrectAnswer   = "correct"
	IncorrectAnswer = "Incorrect"
)

func NewScenario(q []Question, repeat, maxErrors int32) *Scenario {
	return &Scenario{
		q,
		repeat,
		0,
		0,
		maxErrors,
		q[0],
		Ordered,
	}
}

func (s *Scenario) Restart() {
	s.index = 0
	s.correct = 0

	if nextQuestionTimer != nil {
		nextQuestionTimer.Stop()
		nextQuestionTimer = nil
	}
}

func (s *Scenario) Play() bool {
	if s.index >= s.repeat {
		return false
	}

	if nextQuestionTimer == nil {
		s.activeQuestion.Draw(50, 70, 100, 100)
	} else {
		showResult(lastAnswerWasCorrect)
	}

	if a := s.activeQuestion.IsAnswerCorrect(); a != nil {
		if *a == true || s.activeQuestion.Tries() >= s.maxErrors {
			lastAnswerWasCorrect = false

			if *a == true {
				s.correct = s.correct + 1
				lastAnswerWasCorrect = true
			}

			s.index = s.index + 1
			s.activeQuestion.Reset()

			nextQuestionTimer = time.AfterFunc(NextQuestionDelay, func() {
				nextQuestionTimer = nil
				s.activeQuestion = s.RotateQuestions()
			})
		}
	}

	return true
}

func (s *Scenario) Order(o order) {
	s.order = o
}

func (s *Scenario) RotateQuestions() Question {
	rand.Seed(time.Now().UTC().UnixNano())
	set := rand.Intn(len(s.question))

	if s.order == Ordered {
		set = int(s.index * int32(len(s.question)) / s.repeat)
	}

	return s.question[set]
}

func (s *Scenario) Repeats() (total int32, correct int32) {
	total = s.index
	correct = s.correct

	return
}

func showResult(isCorrect bool) {
	c := rl.Green
	t := CorrectAnswer

	if isCorrect == false {
		c = rl.Red
		t = IncorrectAnswer
	}

	r := rl.NewRectangle(40, 90, 20, 20)
	o := rl.MeasureText(t, rl.GetFontDefault().BaseSize)

	raygui.LabelEx(r, fmt.Sprintf("%s!", t), c, raygui.BackgroundColor(), raygui.BackgroundColor())

	r.X = r.X + 10 + float32(o)
	raygui.Label(r, "Ready for next question ....")

}
