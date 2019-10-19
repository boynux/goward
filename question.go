package main

import (
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	True  = true
	False = false
)

type Question interface {
	IsAnswerCorrect() *bool
	Draw(x, y, w, h float32)
	Reset()
	Tries() int32
}

type ChoiceQuestion struct {
	gen     Generator
	choice  *string
	answers []string
	tries   int32
}

func NewQuestion(g Generator /*question string, answers []string, correct int32*/) Question {
	return &ChoiceQuestion{
		g,
		nil,
		[]string{},
		0,
	}
}

func (q *ChoiceQuestion) Reset() {
	q.choice = nil
	q.answers = []string{}
	q.tries = 0
}

func (q *ChoiceQuestion) IsAnswerCorrect() *bool {
	if q.choice == nil {
		return nil
	} else if *q.choice == q.gen.Answer() {
		return &True
	} else {
		return &False
	}
}

func (q *ChoiceQuestion) Draw(x, y, w, h float32) {
	if len(q.answers) == 0 {
		q.gen.Generate()
		q.answers = make([]string, q.gen.MaxChoices())
		q.answers[0] = q.gen.Answer()

		for i := 1; i <  q.gen.MaxChoices(); i++ {
			q.answers[i] = q.gen.FalseAnswer()
		}

		// shuffle answers :)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(q.answers), func(i, j int) { q.answers[i], q.answers[j] = q.answers[j], q.answers[i] })
	}

	raygui.Label(rl.NewRectangle(x+5, y+5, 20, 20), q.gen.QuestionString())
	for i, _ := range q.answers {
		chosen := raygui.Button(rl.NewRectangle((x+5+float32(45*i)), y+20+5+5, 40, 20), q.answers[i])

		if chosen {
			if q.choice != &q.answers[i] {
				q.tries = q.tries + 1
			}

			q.choice = &q.answers[i]
		}
	}
}

func (q *ChoiceQuestion) Tries() int32 {
	return q.tries
}
