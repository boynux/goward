package generators

type BasicAdditionGenerator struct {
	*BasicGenerator
}

func NewBasicAdditionGenerator(min, max int32) *BasicAdditionGenerator {
	var index int32 = 2
	return &BasicAdditionGenerator{
		NewBasicGenerator(min, max, &index, []string{"+"}),
	}
}
