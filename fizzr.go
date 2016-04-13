package fizzr

import (
	"fmt"
)

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
