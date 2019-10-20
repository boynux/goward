package main

import (
	"math/rand"
	"strconv"
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

type RectangleChoiceQuestion struct {
	ChoiceQuestion
}

func NewChoiceQuestion(g Generator /*question string, answers []string, correct int32*/) Question {
	return &ChoiceQuestion{
		g,
		nil,
		[]string{},
		0,
	}
}

func NewRectangleChoiceQuestion(g *BasicAdditionGenerator) Question {
	return &RectangleChoiceQuestion{
		ChoiceQuestion{
			g,
			nil,
			[]string{},
			0,
		},
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
		q.answers = generateAnswerCollection(q.gen)
	}

	raygui.Label(rl.NewRectangle(x, y+5, 20, 20), q.gen.QuestionString())
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

func generateAnswerCollection(g Generator) []string {
	g.Generate()
	a := make([]string, g.MaxChoices())
	a[0] = g.Answer()

	for i := 1; i < g.MaxChoices(); i++ {
		a[i] = g.FalseAnswer()
	}

	// shuffle answers :)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}

func (q *RectangleChoiceQuestion) Draw(x, y, w, h float32) {
	q.ChoiceQuestion.Draw(x, y, w, h)

	a, _ := strconv.ParseInt(q.gen.Answer(), 10, 32)
	c := rl.Blue

	for i := 0; i < int(a/10+1); i++ {
		for j := 0; j < 10; j++ {

			if i*10+j+1 > int(q.gen.(*BasicAdditionGenerator).GetLeft()) {
				c = rl.Green
			}

			top := int32(y-25) + int32(j/5+1)*9
			left := int32(x) + int32(j%5)*9 + 5 + 50*int32(i)

			if i*10+j+1 > int(a) {
				rl.DrawRectangleLines(left, top, 8, 8, rl.Blue)
			} else {
				rl.DrawRectangle(left, top, 8, 8, c)
			}
		}
	}
}
