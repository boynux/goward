package generators

import (
	"fmt"
	"math/rand"
	"time"
)

type BasicGenerator struct {
	Max        int32
	Min        int32
	BlankIndex *int32
	Left       int32
	Right      int32
	Operator   *string
	Result     int32
	Holder     int32
	Operators  []string
}

func NewBasicGenerator(min, max int32, blankIndex *int32, o []string) *BasicGenerator {
	if len(o) == 0 {
		o = []string{"+"}
	}

	g := &BasicGenerator{
		max,
		min,
		blankIndex,
		-1,
		-1,
		nil,
		-1,
		-1,
		o,
	}

	return g
}

func (a *BasicGenerator) Generate() bool {
	rand.Seed(time.Now().UTC().UnixNano())
	a.Left = rand.Int31n(a.Max-a.Min+1) + a.Min
	a.Right = rand.Int31n(a.Max-a.Min+1) + a.Min

	a.Operator = &a.Operators[rand.Intn(len(a.Operators))]

	switch *a.Operator {
	case "+":
		a.Result = a.Left + a.Right
	case "-":
		a.Result = a.Left - a.Right
		if a.Result < 0 {
			a.Result = a.Result * -1
			a.Left, a.Right = a.Right, a.Left
		}
	case "*":
		a.Result = a.Left * a.Right
	default:
		a.Result = a.Left + a.Right
	}

	blankIndex := rand.Int31n(3)

	if a.BlankIndex != nil {
		blankIndex = *a.BlankIndex
	}

	switch blankIndex {
	case 0:
		a.Holder = a.Left
		a.Left = -1
	case 1:
		a.Holder = a.Right
		a.Right = -1
	case 2:
		a.Holder = a.Result
		a.Result = -1
	}

	return true
}

func (a *BasicGenerator) QuestionString() string {
	q := "____ "

	if a.Left != -1 {
		q = fmt.Sprintf("%d", a.Left)
	}

	q += " " + *a.Operator + " "

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

func (a *BasicGenerator) Answer() string {
	return fmt.Sprintf("%d", a.Holder)
}

func (a *BasicGenerator) FalseAnswer() string {
	rand.Seed(time.Now().UTC().UnixNano())

	max := a.Max-a.Min+1

	if a.BlankIndex != nil && *a.BlankIndex == 2 {
		max = a.Max+a.Min
	}

	for {
		ans := fmt.Sprintf("%d", rand.Int31n(max)+a.Min)

		if ans != a.Answer() {
			return ans
		}
	}
}

func (a *BasicGenerator) MaxChoices() int {
	return 4
}

func (a *BasicGenerator) GetLeft() int32 {
	return a.Left
}

func (a *BasicGenerator) GetRight() int32 {
	return a.Right
}
