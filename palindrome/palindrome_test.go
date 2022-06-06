package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	palindromeTests := []struct {
		test string
		got  string
		want bool
	}{
		{test: "abba is a palindrome", got: "abba", want: true},
		{test: "ben is not a palindrome", got: "ben", want: false},
		{test: "hello is not a palindrome", got: "hello", want: false},
		{test: "radar is a palindrome", got: "radar", want: true},
	}

	for _, tt := range palindromeTests {
		t.Run(tt.test, func(t *testing.T) {
			if tt.want == true {
				if IsPalindrome(tt.got) != tt.want {
					t.Errorf("%s is not a palindrome", tt.got)
				}
			} else {
				if IsPalindrome(tt.got) != tt.want {
					t.Errorf("%s is a palindrome", tt.got)
				}
			}

		})
	}
}
