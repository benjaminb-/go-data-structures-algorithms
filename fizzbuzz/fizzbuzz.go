package fizzbuzz

import (
	"strconv"
)

// Return a slice of numbers as string from 1 to n
// if the number is multiple of 3 replace the number by Fizz
// if the number is multiple of 5 replace the number by Buzz
// if the number is multiple of 3 and 5 replace the number by FizzBuzz
func FizzBuzz(max int) []string {

	results := []string{}

	for i := 1; i <= max; i++ {
		// we check first multiple of 3 and 5
		if isMultipleOf(i, 15) {
			results = append(results, "FizzBuzz")
		} else if i%3 == 0 {
			results = append(results, "Fizz")
		} else if i%5 == 0 {
			results = append(results, "Buzz")
		} else {
			results = append(results, strconv.Itoa(i))
		}
	}

	return results
}

// Just a simple module implementation
func isMultipleOf(i int, j int) bool {
	div := i / j
	return (i - (j * div)) == 0

}
