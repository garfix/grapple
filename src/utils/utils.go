package utils

func StringArrayContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func ComputePairFreqs(wordFreqs map[string]int, splits map[string][]string) map[string]int {
	pairFreqs := map[string]int{}
	for word, freq := range wordFreqs {
		split := splits[word]
		if len(split) == 1 {
			continue
		}
		for i := 0; i < len(split)-1; i++ {
			pair := split[i] + split[i+1]
			_, found := pairFreqs[pair]
			if !found {
				pairFreqs[pair] = 0
			}
			pairFreqs[pair] += freq
		}
	}
	return pairFreqs
}

func MergePair(a string, b string, wordFreqs map[string]int, splits map[string][]string) map[string][]string {
	for word := range wordFreqs {
		split := splits[word]
		if len(split) == 1 {
			continue
		}

		i := 0
		for i < len(split)-1 {
			if split[i] == a && split[i+1] == b {
				newSplit := []string{}
				newSplit = append(newSplit, split[:i]...)
				newSplit = append(newSplit, a+b)
				newSplit = append(newSplit, split[i+2:]...)
				split = newSplit
			} else {
				i += 1
			}
		}
		splits[word] = split
	}
	return splits
}

func MergeN(vocabulary []string, splits map[string][]string, wordFreqs map[string]int, vocabularySize int) (map[string]string, []string) {

	merges := map[string]string{}

	for len(vocabulary) < vocabularySize {
		pairFreqs := ComputePairFreqs(wordFreqs, splits)

		bestPair := ""
		maxFreq := 0

		for pair, freq := range pairFreqs {
			if maxFreq < freq {
				bestPair = pair
				maxFreq = freq
			}
		}

		splits = MergePair(string(bestPair[0]), string(bestPair[1]), wordFreqs, splits)
		merges[bestPair] = string(bestPair[0]) + string(bestPair[1])
		vocabulary = append(vocabulary, string(bestPair[0])+string(bestPair[1]))
	}

	return merges, vocabulary
}