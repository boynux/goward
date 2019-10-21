package questions

import (
	"math"

	"github.com/boynux/goward/generators"
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ImageScale       = 0.5
	MinuteHandLength = 20
	HourHandLength   = 12
)

var clockTexture rl.Texture2D

type ClockChoiceQuestion struct {
	ChoiceQuestion
}

func NewClockChoiceQuestion(g *generators.ClockGenerator) Question {
	i := rl.LoadImage("images/clock.png")
	clockTexture = rl.LoadTextureFromImage(i)
	rl.UnloadImage(i)

	return &ClockChoiceQuestion{
		ChoiceQuestion{
			g,
			nil,
			[]string{},
			0,
		},
	}
}

func (q *ClockChoiceQuestion) Draw(x, y, w, h float32) {
	q.ChoiceQuestion.Draw(x, y, w, h)

	c := rl.Vector2{
		x + w/2 + float32(clockTexture.Width)*ImageScale*0.5,
		y - 50 + float32(clockTexture.Height)*ImageScale*0.5,
	}

	mx, my := math.Sincos(float64(q.gen.(*generators.ClockGenerator).GetMinute()) / 60 * math.Pi * 2)
	hx, hy := math.Sincos(float64(q.gen.(*generators.ClockGenerator).GetHour()) / 12 * math.Pi * 2)

	rl.DrawTextureEx(clockTexture, rl.Vector2{x + w/2, y - 50}, 0, ImageScale, raygui.BackgroundColor())

	rl.DrawLineEx(c, rl.Vector2{c.X + MinuteHandLength*float32(mx), c.Y - MinuteHandLength*float32(my)}, 1, rl.Black)
	rl.DrawLineEx(c, rl.Vector2{c.X + HourHandLength*float32(hx), c.Y - HourHandLength*float32(hy)}, 2, rl.Black)
}
