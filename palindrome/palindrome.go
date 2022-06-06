package palindrome

// Check if a string is a palindrome
func IsPalindrome(str string) bool {
	characters := []rune(str)
	var reverted string

	for i := 0; i < len(characters); i++ {
		reverted = string(characters[i]) + reverted
	}

	return reverted == str
}
