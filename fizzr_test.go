package fizzr

import (
	"testing"
)

func TestEvenOdd(t *testing.T) {
	expected := "Odd\nEven\nOdd\nEven\nOdd\n"
	mappings := []F{&Even{}, &Odd{}}
	out := Fizzr(5, mappings)
	if out != expected {
		t.Errorf("Expected %s got %s", expected, out)
	}
}

func TestFizzBuzz(t *testing.T) {
	expected := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n" +
		"14\nFizzBuzz\n16\n"
	mappings := []F{&Fizz{}, &Buzz{}}
	out := Fizzr(16, mappings)
	if out != expected {
		t.Errorf("Expected %s got %s", expected, out)
	}
}

func TestFizzBuzzBazz(t *testing.T) {
	expected := "1\n2\nFizz\n4\nBuzz\nFizz\nBazz\n8\nFizz\nBuzz\n11\nFizz\n13\n" +
		"Bazz\nFizzBuzz\n16\n17\nFizz\n19\nBuzz\nFizzBazz\n22\n23\nFizz\nBuzz\n" +
		"26\nFizz\nBazz\n29\nFizzBuzz\n31\n32\nFizz\n34\nBuzzBazz\nFizz\n37\n38\n" +
		"Fizz\nBuzz\n41\nFizzBazz\n43\n"
	mappings := []F{&Fizz{}, &Buzz{}, &Bazz{}}
	out := Fizzr(43, mappings)
	if out != expected {
		t.Errorf("Expected %s got %s", expected, out)
	}
}
