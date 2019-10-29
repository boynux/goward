package generators

import (
	"fmt"
	"math/rand"
	"time"
)

type ClockGenerator struct {
	min      int32
	hour     int32
	military bool
}

func NewClockGenerator(military bool) *ClockGenerator {
	return &ClockGenerator{
		military: military,
	}
}

func (c *ClockGenerator) Generate() bool {
	rand.Seed(time.Now().UTC().UnixNano())

	c.min = rand.Int31n(60) / 15 * 15

	if c.military {
		c.hour = rand.Int31n(24) + 1
	} else {
		c.hour = rand.Int31n(12) + 1
	}

	return true
}

func (c *ClockGenerator) QuestionString() string {
	return "" //fmt.Sprintf("What time is it?")
}

func (c *ClockGenerator) Answer() string {
	return fmt.Sprintf("%d:%d", c.hour, c.min)
}

func (c *ClockGenerator) FalseAnswer() string {
	maxHour := int32(12)
	if c.military {
		maxHour = 24
	}

	for {
		s := fmt.Sprintf("%d:%d", rand.Int31n(maxHour)+1, rand.Int31n(60)/15*15)
		if s != c.Answer() {
			return s
		}
	}
}

func (c *ClockGenerator) MaxChoices() int {
	return 4
}

func (c *ClockGenerator) GetMinute() int32 {
	return c.min
}

func (c *ClockGenerator) GetHour() int32 {
	return c.hour
}
