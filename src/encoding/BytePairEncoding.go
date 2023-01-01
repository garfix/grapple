package encoding

import "fmt"

type BytePairEncoding struct {
}

func CreateBytePairEncoding() *BytePairEncoding {
	return &BytePairEncoding{}
}

func (e *BytePairEncoding) Encode(strings []string, numberOfMerges int) bool {
	vocabulary := e.AllUniqueCharacters(strings)
	fmt.Printf("%c", vocabulary)

	for i := 1; i <= numberOfMerges; i++ {

	}

	return false
}

// AllUniqueCharacters returns an array of all characters that appear in strings, order by decreasing order of occurrence
func (e *BytePairEncoding) AllUniqueCharacters(strings []string) []rune {
	var charsFound []rune
	var countsFound []int

	for _, s := range strings {
		for _, r := range s {
			found := false
			for i, c := range charsFound {
				if c == r {
					countsFound[i] = countsFound[i] + 1
					found = true
					break
				}
			}
			if !found {
				charsFound = append(charsFound, r)
				countsFound = append(countsFound, 1)
			}
		}
	}

	var orderedChars []rune
	var orderedCounts []int

	for j, foundChar := range charsFound {
		foundCount := countsFound[j]
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
