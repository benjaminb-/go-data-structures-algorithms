package anagrams

import (
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

// Check if the 2 provided strings are anagrams of each other
// A string is an anagram of another if it uses the same characters
// Ignore spaces, punctuation, capital letters are considered same as lowercase
func Anagrams(str1, str2 string) bool {
	// first keep only alphanumeric characters based on the specs
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// apply regex and lowercase
	cleanedStr1 := strings.ToLower(reg.ReplaceAllString(str1, ""))
	cleanedStr2 := strings.ToLower(reg.ReplaceAllString(str2, ""))

	// Method using character map
	// if len(cleanedStr1) != len(cleanedStr2) {
	// 	return false
	// }

	// // generate character maps
	// str1Map := createCharactersMap(cleanedStr1)
	// str2Map := createCharactersMap(cleanedStr2)

	// for _, char := range str1Map {
	// 	idx, exists := findIdxInCharactersMap(str2Map, char.char)
	// 	if exists == false {
	// 		return false
	// 	}

	// 	found := str2Map[idx]
	// 	if found.count != char.count {
	// 		return false
	// 	}
	// }

	// simpler version, solved after cleanup, sorting then compare
	cleanedStr1Sorted := strings.Split(cleanedStr1, "")
	sort.Strings(cleanedStr1Sorted)
	cleanedStr2Sorted := strings.Split(cleanedStr2, "")
	sort.Strings(cleanedStr2Sorted)

	return reflect.DeepEqual(cleanedStr1Sorted, cleanedStr2Sorted)
}

type charMap struct {
	char  string
	count int
}

func createCharactersMap(str string) []charMap {
	strings := []rune(str)
	cmap := []charMap{}

	for i := 0; i < len(strings); i++ {
		idx, exists := findIdxInCharactersMap(cmap, string(strings[i]))
		if exists {
			cmap[idx].count++
		} else {
			char := charMap{char: string(strings[i]), count: 1}
			cmap = append(cmap, char)
		}
	}

	return cmap
}

func findIdxInCharactersMap(cmap []charMap, char string) (int, bool) {
	for i, a := range cmap {
		if a.char == char {
			return i, true
		}
	}
	return -1, false
}
