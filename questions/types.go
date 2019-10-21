package questions

var (
	True  = true
	False = false
)

type Question interface {
	IsAnswerCorrect() *bool
	Draw(x, y, w, h float32)
	Reset()
	Tries() int32
}
