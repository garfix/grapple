package utils

import (
	"fmt"

	"github.com/garfix/grapple/src/tokenizer"
)

func StringArrayContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

type stringPair struct {
	first  string
	second string
}

type stringPairFreq struct {
	pair stringPair
	freq int
}

type stringPairFreqs struct {
	freqs map[stringPair]int
	order []stringPair
}

func CreateStringPairFreqs() stringPairFreqs {
	return stringPairFreqs{
		freqs: map[stringPair]int{},
		order: []stringPair{},
	}
}

func (f stringPairFreqs) add(pair stringPair, freq int) stringPairFreqs {
	_, found := f.freqs[pair]
	if found {
		f.freqs[pair] += freq
	} else {
		f.order = append(f.order, pair)
		f.freqs[pair] = freq
	}
	return f
}

func (f stringPairFreqs) getAll() []stringPairFreq {
	items := []stringPairFreq{}
	for _, item := range f.order {
		items = append(items, stringPairFreq{item, f.freqs[item]})
	}
	return items
}

func ComputePairFreqs(wordFreqs map[string]int, splits map[string][]string) stringPairFreqs {
	pairFreqs := CreateStringPairFreqs()
	for word, freq := range wordFreqs {
		split := splits[word]
		if len(split) == 1 {
			continue
		}
		for i := 0; i < len(split)-1; i++ {
			pair := stringPair{split[i], split[i+1]}
			pairFreqs = pairFreqs.add(pair, freq)
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

func MergeN(vocabulary []string, splits map[string][]string, wordFreqs map[string]int, vocabularySize int) (map[stringPair]string, []string) {

	merges := map[stringPair]string{}

	for len(vocabulary) < vocabularySize {
		pairFreqs := ComputePairFreqs(wordFreqs, splits)

		bestPair := stringPair{}
		maxFreq := 0

		for _, pairFreq := range pairFreqs.getAll() {
			if maxFreq < pairFreq.freq {
				bestPair = pairFreq.pair
				maxFreq = pairFreq.freq
			}
		}

		a := bestPair.first
		b := bestPair.second

		splits = MergePair(a, b, wordFreqs, splits)
		merges[bestPair] = a + b

		fmt.Printf("%s: %s (%d)\n", bestPair, a+b, maxFreq)
		vocabulary = append(vocabulary, a+b)
	}

	return merges, vocabulary
}

func Tokenize(text string, merges map[stringPair]string) []string {
	// preTokenizeResult = tokenizer._tokenizer.pre_tokenizer.pre_tokenize_str(text)
	// pre_tokenized_text = [word for word, offset in pre_tokenize_result]

	tok := tokenizer.CreateSimpleTokenizer()
	tokens := tok.Tokenize(text)
	words := tokenizer.AddBeginToken(tokens, "Ġ")

	splits := map[string][]string{}
	for _, word := range words {
		splits[word] = []string{}
		for _, letter := range word {
			splits[word] = append(splits[word], string(letter))
		}
	}

	// splits = [[l for l in word] for word in pre_tokenized_text]
	for pair, merge := range merges {
		for idx, split := range splits {
			i := 0
			for i < len(split)-1 {
				if split[i] == pair.first && split[i+1] == pair.second {
					newSplit := []string{}
					newSplit = append(newSplit, split[:i]...)
					newSplit = append(newSplit, merge)
					newSplit = append(newSplit, split[i+2:]...)
					split = newSplit
				} else {
					i += 1
				}
			}
			splits[idx] = split
		}
	}

	result := []string{}
	for _, split := range splits {
		result = append(result, split...)
	}
	return result
}
