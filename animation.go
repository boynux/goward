package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	shapes []Shape
}

func NewAnimation() *Animation {
	a := &Animation{}

	s := NewShape()
	a.shapes = make([]Shape, 10)

	for i := 0; i < len(a.shapes); i++ {
		a.shapes[i] = s.Clone()
		s.Update()
	}

	return a
}

func (a *Animation) Draw() {
	for i := 0; i < len(a.shapes); i++ {
		a.shapes[i].Draw()
	}
}

func (a *Animation) Update() {
	for i := 0; i < len(a.shapes); i++ {
		a.shapes[i].Color = rl.Fade(a.shapes[i].Color, float32(i)/10.0)
		a.shapes[i].Update()
	}
}

type Shape struct {
	Position []rl.Vector2
	velocity []rl.Vector2
	Color    rl.Color
}

func NewShape() Shape {
	p := Shape{
		Color: rl.Green,
	}

	v := rl.Vector2{
		float32(rand.Int31n(7) - 3),
		float32(rand.Int31n(11) - 5),
	}

	if v.X < 1 && v.X > -1 {
		v.X = 1
	}

	if v.Y < 1 && v.Y > -1 {
		v.Y = 1
	}

	p.Position = make([]rl.Vector2, 5)
	p.velocity = make([]rl.Vector2, 5)

	for i := 0; i < 5; i++ {
		p.Position[i] = rl.NewVector2(rand.Float32()*float32(rl.GetScreenWidth()),
			rand.Float32()*float32(rl.GetScreenHeight()))
		p.velocity[i] = v
	}

	return p
}

func (s Shape) Clone() Shape {
	c := Shape{
		Color: s.Color,
	}

	c.Position = make([]rl.Vector2, len(s.Position))
	c.velocity = make([]rl.Vector2, len(s.velocity))

	copy(c.Position, s.Position)
	copy(c.velocity, s.velocity)

	return c
}

func (s Shape) Update() {
	for i := 0; i < len(s.Position); i++ {
		s.Position[i].X = s.Position[i].X + s.velocity[i].X
		s.Position[i].Y = s.Position[i].Y + s.velocity[i].Y

		if s.Position[i].X < 0 || s.Position[i].X > float32(rl.GetScreenWidth()) {
			s.velocity[i].X = s.velocity[i].X * -1
		}

		if s.Position[i].Y < 0 || s.Position[i].Y > float32(rl.GetScreenHeight()) {
			s.velocity[i].Y = s.velocity[i].Y * -1
		}
	}
}

func (s Shape) Draw() {
	rl.DrawPolyExLines(s.Position, int32(len(s.Position)), s.Color)
}
