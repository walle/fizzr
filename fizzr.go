// Package fizzr contains logic for building extensible "FizzBuzz" output.
package fizzr

import (
	"fmt"
)

// Mapping is a mapping interface. Each mapping implements the Out function.
type Mapping interface {
	Out(n int) string
}

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

// Fizz is the fizz number mapping.
type Fizz bool

// Out returns Fizz if the number n is evenly divisible by 3.
func (f *Fizz) Out(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	return ""
}

// Buzz is the buzz number mapping.
type Buzz bool

// Out returns Buzz if the number n is evenly divisible by 5.
func (b *Buzz) Out(n int) string {
	if n%5 == 0 {
		return "Buzz"
	}
	return ""
}

// Bazz is the bazz number mapping.
type Bazz bool

// Out returns Bazz if the number n is evenly divisible by 7.
func (b *Bazz) Out(n int) string {
	if n%7 == 0 {
		return "Bazz"
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
