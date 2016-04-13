package fizzr_test

import (
	"fmt"

	"github.com/walle/fizzr"
)

func ExampleFizzr() {
	mappings := []fizzr.Mapping{fizzr.Fizz, fizzr.Buzz}
	ret := fizzr.Fizzr(16, mappings)
	fmt.Println(ret)
	// Output: 1
	//2
	//Fizz
	//4
	//Buzz
	//Fizz
	//7
	//8
	//Fizz
	//Buzz
	//11
	//Fizz
	//13
	//14
	//FizzBuzz
	//16
	//
}
