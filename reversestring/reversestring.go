package reversestring

// Reverse a string
func ReverseString(str string) string {
	chars := []rune(str)
	var reversed string

	for i := 0; i < len(chars); i++ {
		reversed = string(chars[i]) + reversed
	}

	return reversed
}
