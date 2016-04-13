package fizzr

// ModulusMapping is a mapper that checks if a number is evenly divisible.
type ModulusMapping struct {
	output string
	number int
}

// NewModulusMapping returns a ModulusMapping with the selected settings.
func NewModulusMapping(number int, output string) *ModulusMapping {
	return &ModulusMapping{output: output, number: number}
}

// Out returns the output if n is evenly divisible with number.
func (m *ModulusMapping) Out(n int) string {
	if n%m.number == 0 {
		return m.output
	}
	return ""
}
