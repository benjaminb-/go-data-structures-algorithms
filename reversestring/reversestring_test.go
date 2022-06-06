package reversestring

import "testing"

func TestReverseString(t *testing.T) {

	revereStringTests := []struct {
		got  string
		want string
	}{
		{got: "0", want: "0"},
		{got: "ab", want: "ba"},
		{got: "test", want: "tset"},
		{got: "@)lw", want: "wl)@"},
	}

	for _, tt := range revereStringTests {
		t.Run(tt.got+" "+tt.want, func(t *testing.T) {
			if ReverseString(tt.got) != tt.want {
				t.Errorf("got %s want %s", tt.got, tt.want)
			}
		})
	}

}
