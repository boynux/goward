package generators

import (
	"fmt"
	"math/rand"
	"time"
)

type ClockGenerator struct {
	min  int32
	hour int32
}

func NewClockGenerator() *ClockGenerator {
	return &ClockGenerator{}
}

func (c *ClockGenerator) Generate() bool {
	rand.Seed(time.Now().UTC().UnixNano())

	c.min = rand.Int31n(60) / 15 * 15
	c.hour = rand.Int31n(12) + 1

	return true
}

func (c *ClockGenerator) QuestionString() string {
	return "" //fmt.Sprintf("What time is it?")
}

func (c *ClockGenerator) Answer() string {
	return fmt.Sprintf("%d:%d", c.hour, c.min)
}

func (c *ClockGenerator) FalseAnswer() string {
	return fmt.Sprintf("%d:%d", rand.Int31n(12)+1, rand.Int31n(60)/15*15)
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
