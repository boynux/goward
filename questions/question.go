package questions

import (
	"math/rand"
	"time"

	"github.com/boynux/goward/generators"
)

func generateAnswerCollection(g generators.Generator) []string {
	g.Generate()
	a := make([]string, g.MaxChoices())
	a[0] = g.Answer()

	for i := 1; i < g.MaxChoices(); i++ {
		a[i] = g.FalseAnswer()
	}

	// shuffle answers :)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}
