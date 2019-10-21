package generators

type Generator interface {
	Generate() bool
	QuestionString() string
	Answer() string
	FalseAnswer() string
	MaxChoices() int
}
