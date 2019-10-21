package questions

import (
	"strconv"

	"github.com/boynux/goward/generators"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleChoiceQuestion struct {
	ChoiceQuestion
}

func NewRectangleChoiceQuestion(g *generators.BasicAdditionGenerator) Question {
	return &RectangleChoiceQuestion{
		ChoiceQuestion{
			g,
			nil,
			[]string{},
			0,
		},
	}
}

func (q *RectangleChoiceQuestion) Draw(x, y, w, h float32) {
	q.ChoiceQuestion.Draw(x, y, w, h)

	a, _ := strconv.ParseInt(q.gen.Answer(), 10, 32)
	c := rl.Blue

	for i := 0; i < int(a/10+1); i++ {
		for j := 0; j < 10; j++ {

			if i*10+j+1 > int(q.gen.(*generators.BasicAdditionGenerator).GetLeft()) {
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
