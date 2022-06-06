package maxchar

import "testing"

func TestMaxChar(t *testing.T) {
	reverseIntTests := []struct {
		test string
		got  string
		want string
	}{
		{test: "max char of jhjhqwjqwdqwwwww should be w", got: "jhjhqwjqwdqwwwww", want: "w"},
		{test: "max char of aaabbaba should be a", got: "aaabbaba", want: "a"},
		{test: "max char of c should be c", got: "c", want: "c"},
		{test: "max char of zzzzzzzzzzzzzz should be z", got: "zzzzzzzzzzzzzz", want: "z"},
	}

	for _, tt := range reverseIntTests {
		t.Run(tt.test, func(t *testing.T) {
			if MaxChar(tt.got) != tt.want {
				t.Errorf("got %s want %s", MaxChar(tt.got), tt.want)
			}
		})
	}
}
