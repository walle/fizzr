package fizzr

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
