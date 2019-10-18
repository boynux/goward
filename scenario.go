package main

import "fmt"

var index int

type Scenario struct {
	question  Question
	repeat    int32
	index     int32
	correct   int32
	maxErrors int32
}

func NewScenario(q Question, repeat, maxErrors int32) *Scenario {
	return &Scenario{
		q,
		repeat,
		0,
		0,
		maxErrors,
	}
}

func (s *Scenario) Play() bool {
	if s.index >= s.repeat {
		return false
	}

	s.question.Draw(50, 70, 100, 100)

	if a := s.question.IsAnswerCorrect(); a != nil {
		fmt.Printf("%v\n", s.question.Tries())
		if *a == true || s.question.Tries() >= s.maxErrors {
			if *a == true {
				s.correct = s.correct + 1
			}

			s.index = s.index + 1
			s.question.Reset()
		}
	}

	return true
}

func (s *Scenario) Repeats() (total int32, correct int32) {
	total = s.index
	correct = s.correct

	return
}
