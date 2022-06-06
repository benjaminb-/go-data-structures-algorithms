package reverseint

import "testing"

func TestReverseInt(t *testing.T) {

	reverseIntTests := []struct {
		test string
		got  int
		want int
	}{
		{test: "0 should be 0", got: 0, want: 0},
		{test: "23 should be 32", got: 23, want: 32},
		{test: "5000 should be 5", got: 5000, want: 5},
		{test: "-656 should be -656", got: -656, want: -656},
		{test: "-0005 should be -5", got: -0005, want: -5},
	}

	for _, tt := range reverseIntTests {
		t.Run(tt.test, func(t *testing.T) {
			if ReverseInt(tt.got) != tt.want {
				t.Errorf("got %d want %d", ReverseInt(tt.got), tt.want)
			}
		})
	}
}
