package fizzr

// Mapping is a mapping interface. Each mapping implements the Out function.
type Mapping interface {
	Out(n int) string
}
