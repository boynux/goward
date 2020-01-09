package main

import (
	"time"
	"flag"
)

type config struct {
	TotalQuestionsPerScene int
	QuestionTimeout        time.Duration
	IFTTTKey               string
}

var Config *config

func init() {
	c := &config{}

	flag.IntVar(&c.TotalQuestionsPerScene, "total", 30, "Total number of questions to ask per try")
	flag.DurationVar(&c.QuestionTimeout, "timeout", 1 * time.Minute, "Allowed timeout for each qwustion")
	flag.StringVar(&c.IFTTTKey, "ifttt-key", "", "IFTTT Key")

	flag.Parse()

	Config = c
}
