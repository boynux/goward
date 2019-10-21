package questions

import (
	"github.com/boynux/goward/generators"
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChoiceQuestion struct {
	gen     generators.Generator
	choice  *string
	answers []string
	tries   int32
}

func NewChoiceQuestion(g generators.Generator /*question string, answers []string, correct int32*/) Question {
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
