package fizzr

import (
	"testing"
)

func TestEvenOdd(t *testing.T) {
	expected := "Odd\nEven\nOdd\nEven\nOdd\n"
	mappings := []F{new(Even), new(Odd)}
	out := Fizzr(5, mappings)
	if out != expected {
		t.Errorf("Expected %s got %s", expected, out)
	}
}

func TestFizzBuzz(t *testing.T) {
	expected := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n" +
		"14\nFizzBuzz\n16\n"
	mappings := []F{new(Fizz), new(Buzz)}
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
	mappings := []F{new(Fizz), new(Buzz), new(Bazz)}
	out := Fizzr(43, mappings)
	if out != expected {
		t.Errorf("Expected %s got %s", expected, out)
	}
}

func TestFizzBuzzBazzSpecial(t *testing.T) {
	expected := "1\n2\nFizz\n4\nBuzz\nFizz\nBazz\n8\nFizz\nBuzz\n11\nFizz\n13\n" +
		"Bazz\nFizzBuzz\n16\n17\nFizz\n19\nBuzz\nFizzBazz\n22\n23\nFizz\nBuzz\n" +
		"26\nFizz\nBazz\n29\nFizzBuzz\n31\n32\nFizz\n34\nBuzzBazz\nFizz\n37\n38\n" +
		"Fizz\nBuzz\n41\nFizzBazzThe Answer\n43\n"
	mappings := []F{new(Fizz), new(Buzz), new(Bazz), &Special{42}}
	out := Fizzr(43, mappings)
	if out != expected {
		t.Errorf("Expected %s got %s", expected, out)
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	expected := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n" +
		"14\nFizzBuzz\n16\n"
	mappings := []F{new(Fizz), new(Buzz)}
	for i := 0; i < b.N; i++ {
		out := Fizzr(16, mappings)
		if out != expected {
			b.Errorf("Expected %s got %s", expected, out)
		}
	}
}

func BenchmarkEvenOdd(b *testing.B) {
	expected := "Odd\nEven\nOdd\nEven\nOdd\n"
	mappings := []F{new(Even), new(Odd)}
	for i := 0; i < b.N; i++ {
		out := Fizzr(5, mappings)
		if out != expected {
			b.Errorf("Expected %s got %s", expected, out)
		}
	}
}
