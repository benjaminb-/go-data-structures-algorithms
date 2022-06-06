package anagrams

import "testing"

func TestAnagrams(t *testing.T) {
	anagramsTests := []struct {
		test string
		got  []string
		want bool
	}{
		{test: "Debit card and Bad credit are anagrams",
			got: []string{"Debit card", "Bad credit"}, want: true},
		{test: "Hello world! and Bye world are not anagrams",
			got: []string{"Hello world!", "Bye world"}, want: false},
		{test: "Hello and Hellos are not anagrams",
			got: []string{"Hello", "Hellos"}, want: false},
	}

	for _, tt := range anagramsTests {
		t.Run(tt.test, func(t *testing.T) {
			if Anagrams(tt.got[0], tt.got[1]) != tt.want {
				t.Error("failed test should be", tt.want)
			}
		})
	}
}
