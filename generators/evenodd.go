package generators

import (
	"fmt"
	"math/rand"
	"time"
)

type EvenOddGenerator struct {
	Max    int32
	Min    int32
	Number int32
}

func NewEvenOddGenerator(min, max int32) *EvenOddGenerator {
	return &EvenOddGenerator{
		max,
		min,
		-1,
	}
}

func (a *EvenOddGenerator) Generate() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	a.Number = rand.Int31n(a.Max-a.Min+1) + a.Min

	return true
}

func (a *EvenOddGenerator) QuestionString() string {
	return fmt.Sprintf("Is %d even or odd?", a.Number)
}

func (a *EvenOddGenerator) Answer() string {
	if a.Number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func (a *EvenOddGenerator) FalseAnswer() string {
	if a.Number%2 == 1 {
		return "even"
	} else {
		return "odd"
	}
}

func (a *EvenOddGenerator) MaxChoices() int {
	return 2
}
