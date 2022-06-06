package reverseint

import (
	"math"
	"strconv"
)

// Reverse an integer, preserve the sign
func ReverseInt(i int) int {
	// check if the input is negative
	isNegative := math.Signbit(float64(i))

	// set the input as positive
	if isNegative {
		i = -(i)
	}

	// convert integer to string
	integers := []rune(strconv.Itoa(i))
	var reverted string
	for i := 0; i < len(integers); i++ {
		// revert
		reverted = string(integers[i]) + reverted
	}

	// convert back to int
	result, _ := strconv.Atoi(reverted)

	if isNegative {
		// set negative
		return -(result)
	}

	return result
}
