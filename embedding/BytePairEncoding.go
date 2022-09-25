package embedding

import "fmt"

// BytePairEncoding https://en.wikipedia.org/wiki/Byte_pair_encoding
type BytePairEncoding struct {
}

func CreateBytePairEncoding() *BytePairEncoding {
	return &BytePairEncoding{}
}

func (e *BytePairEncoding) Encode(strings []string, numberOfMerges int) bool {
	vocabulary := e.allUniqueCharacters(strings)
	fmt.Printf("%c", vocabulary)

	for i := 1; i <= numberOfMerges; i++ {

	}

	return false
}

// returns an array of all characters that appear in strings, order by decreasing order of occurrence
func (e *BytePairEncoding) allUniqueCharacters(strings []string) []rune {
	charsFound := map[rune]int{}

	for _, s := range strings {
		for _, r := range s {
			count, found := charsFound[r]
			if found {
				count++
			} else {
				count = 1
			}
			charsFound[r] = count
		}
	}

	var orderedCounts []int
	var orderedChars []rune

	for foundChar, foundCount := range charsFound {
		found := false
		for i, orderedCount := range orderedCounts {
			if foundCount > orderedCount {
				orderedCounts = append(orderedCounts[:i+1], orderedCounts[i:]...)
				orderedCounts[i] = foundCount
				orderedChars = append(orderedChars[:i+1], orderedChars[i:]...)
				orderedChars[i] = foundChar
				found = true
				break
			}
		}
		if !found {
			orderedCounts = append(orderedCounts, foundCount)
			orderedChars = append(orderedChars, foundChar)
		}
	}

	return orderedChars
}
