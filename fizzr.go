// Package fizzr contains logic for building extensible "FizzBuzz" output.
package fizzr

import (
	"fmt"
)

// Mapping is a mapping interface. Each mapping implements the Out function.
type Mapping interface {
	Out(n int) string
}

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

// Define the common modulus mappings
var (
	Fizz Mapping = NewModulusMapping(3, "Fizz")
	Buzz Mapping = NewModulusMapping(5, "Buzz")
	Bazz Mapping = NewModulusMapping(7, "Bazz")
)

// Even is the even number mapping.
type Even bool

// Out returns Even if the number n is even.
func (e *Even) Out(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return ""
}

// Odd is the odd number mapping.
type Odd bool

// Out returns Odd if the number n is odd.
func (o *Odd) Out(n int) string {
	if n%2 == 1 {
		return "Odd"
	}
	return ""
}

// Special is the special number mapping.
type Special struct {
	number int
}

// Out returns The Answer if the number n is the answer.
func (s *Special) Out(n int) string {
	if n == s.number {
		return "The Answer"
	}
	return ""
}

// Fizzr runs the mappings against numbers from 1 to n and returns the generated
// string.
func Fizzr(n int, mappings []Mapping) string {
	ret := ""

	for i := 1; i <= n; i++ {
		res := ""
		for _, m := range mappings {
			res += m.Out(i)
		}
		if res == "" {
			ret += fmt.Sprintf("%d", i)
		} else {
			ret += res
		}
		ret += "\n"
	}

	return ret
}
