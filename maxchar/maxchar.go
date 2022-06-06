package maxchar

type CharactersMap struct {
	character string
	max       int
}

// Return the most present character in a string
func MaxChar(str string) string {
	charactersMap := []CharactersMap{}
	characters := []rune(str)
	for i := 0; i < len(characters); i++ {
		index, isPresent := inCharactersMap(charactersMap, string(characters[i]))
		if isPresent {
			charactersMap[index].max++
		} else {
			char := CharactersMap{character: string(characters[i]), max: 1}
			charactersMap = append(charactersMap, char)
		}
	}

	var maxChar string
	maxPresence := 0

	for _, v := range charactersMap {
		if v.max > maxPresence {
			maxPresence = v.max
			maxChar = v.character
		}
	}

	return maxChar
}

func inCharactersMap(s []CharactersMap, str string) (int, bool) {
	for i, v := range s {
		if v.character == str {
			return i, true
		}

	}
	return 0, false
}
