package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator interface {
	Generate() bool
	QuestionString() string
	Answer() string
	FalseAnswer() string
}

type AdditionGenerator struct {
	Max        int32
	Min        int32
	BlankIndex *int32
	Left       int32
	Right      int32
	Result     int32
}

func NewAdditionGenerator(min, max int32, blankIndex *int32) *AdditionGenerator {
	g := &AdditionGenerator{
		max,
		min,
		blankIndex,
		-1,
		-1,
		-1,
	}

	g.Generate()

	return g
}

func (a *AdditionGenerator) Generate() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	a.Left = rand.Int31n(a.Max-a.Min+1) + a.Min
	a.Right = rand.Int31n(a.Max-a.Min+1) + a.Min
	a.Result = a.Left + a.Right

	blankIndex := rand.Int31n(3)

	if a.BlankIndex != nil {
		blankIndex = *a.BlankIndex
	}

	switch blankIndex {
	case 0:
		a.Left = -1
	case 1:
		a.Right = -1
	case 2:
		a.Result = -1
	}

	return true
}

func (a *AdditionGenerator) QuestionString() string {
	q := "____ "

	if a.Left != -1 {
		q = fmt.Sprintf("%d", a.Left)
	}

	q += " + "

	if a.Right != -1 {
		q += fmt.Sprintf("%d", a.Right)
	} else {
		q += " ____ "
	}

	q += " = "

	if a.Result != -1 {
		q += fmt.Sprintf("%d", a.Result)
	} else {
		q += " ____ "
	}

	return q
}

func (a *AdditionGenerator) Answer() string {
	var answer int32

	if a.Left == -1 {
		answer = a.Result - a.Right
	} else if a.Right == -1 {
		answer = a.Result - a.Left
	} else {
		answer = a.Left + a.Right
	}

	return fmt.Sprintf("%d", answer)
}

func (a *AdditionGenerator) FalseAnswer() string {
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		ans := fmt.Sprintf("%d", rand.Int31n(a.Max-a.Min+1)+a.Min)

		if ans != a.Answer() {
			return ans
		}
	}
}
