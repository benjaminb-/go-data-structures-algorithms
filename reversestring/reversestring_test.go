package reversestring

import "testing"

func TestReverseString(t *testing.T) {

	reverseStringTests := []struct {
		test string
		got  string
		want string
	}{
		{test: "0 should be 0", got: "0", want: "0"},
		{test: "ab should be ba", got: "ab", want: "ba"},
		{test: "test should be tset", got: "test", want: "tset"},
		{test: "@)lw should be wl)@", got: "@)lw", want: "wl)@"},
	}

	for _, tt := range reverseStringTests {
		t.Run(tt.test, func(t *testing.T) {
			if ReverseString(tt.got) != tt.want {
				t.Errorf("got %s want %s", ReverseString(tt.got), tt.want)
			}
		})
	}

}
