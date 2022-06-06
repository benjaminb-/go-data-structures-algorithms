package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzzTests := []struct {
		test string
		got  int
		want []string
	}{
		{test: "until 5", got: 5, want: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{test: "until 5", got: 15, want: []string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz"}},
	}

	for _, tt := range fizzBuzzTests {
		t.Run(tt.test, func(t *testing.T) {
			if reflect.DeepEqual(FizzBuzz(tt.got), tt.want) == false {
				t.Error("failed FizzBuzz", FizzBuzz(tt.got), tt.want)
			}
		})
	}

}
