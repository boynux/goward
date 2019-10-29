package generators

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type multiAdditionGenerator struct {
	max        int32
	min        int32
	blankIndex *int32
	numbers    []int32
	result     int32
	holder     int32
}

func NewMultiAdditionGenerator(min, max, numnbers int32, blankIndex *int32) *multiAdditionGenerator {
	if blankIndex != nil && *blankIndex > numnbers {
		panic(fmt.Errorf("Blank index cannot be higher then total number of operands"))
	}

	return &multiAdditionGenerator{
		max,
		min,
		blankIndex,
		make([]int32, numnbers, numnbers),
		-1,
		-1,
	}
}

func (m *multiAdditionGenerator) Generate() bool {
	m.result = 0

	rand.Seed(time.Now().UTC().UnixNano())
	for i := range m.numbers {
		m.numbers[i] = rand.Int31n(m.max-m.min+1) + m.min
		m.result = m.result + m.numbers[i]
	}

	blankIndex := rand.Int31n(int32(len(m.numbers)) + 1)

	if m.blankIndex != nil {
		blankIndex = *m.blankIndex
	}

	if blankIndex < int32(len(m.numbers)) {
		m.holder = m.numbers[blankIndex]
		m.numbers[blankIndex] = -1
	} else {
		m.holder = m.result
		m.result = -1
	}

	return true
}

func (m *multiAdditionGenerator) QuestionString() string {
	blank := "____"
	q := ""

	left := make([]string, len(m.numbers))

	for i, n := range m.numbers {
		left[i] = strconv.Itoa(int(n))

		if n == -1 {
			left[i] = blank
		}
	}

	q = q + strings.Join(left, " + ") + " = "
	if m.result == -1 {
		q = q + blank
	} else {
		q = q + strconv.Itoa(int(m.result))
	}

	return q
}

func (m *multiAdditionGenerator) Answer() string {
	return fmt.Sprintf("%d", m.holder)
}

func (m *multiAdditionGenerator) FalseAnswer() string {
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		ans := fmt.Sprintf("%d", rand.Int31n(m.max-m.min+1)+m.min)

		if ans != m.Answer() {
			return ans
		}
	}
}

func (m *multiAdditionGenerator) MaxChoices() int {
	return 4
}
