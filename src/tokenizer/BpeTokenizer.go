package tokenizer

import (
	"sort"

	"github.com/garfix/grapple/src/utils"
)

// code ported from https://huggingface.co/course/chapter6/5

type BpeTokenizer struct {
	mergeOrder []StringPair
	merges     map[StringPair]string
}

func CreateBpeTokenizer() *BpeTokenizer {
	return &BpeTokenizer{
		mergeOrder: []StringPair{},
		merges:     map[StringPair]string{},
	}
}

func (t *BpeTokenizer) Train(corpus []string) {

	tok := CreateSimpleTokenizer()

	words := []string{}
	wordFreqs := map[string]int{}

	for _, text := range corpus {
		tokens := tok.Tokenize(text)
		for _, word := range AddBeginToken(tokens, "Ġ") {
			freq, found := wordFreqs[word]
			if found {
				freq++
			} else {
				freq = 1
				words = append(words, word)
			}
			wordFreqs[word] = freq
		}
	}

	alphabet := []string{}

	for _, word := range words {
		for _, letter := range word {
			if !utils.StringArrayContains(alphabet, string(letter)) {
				alphabet = append(alphabet, string(letter))
			}
		}

	}

	sort.Strings(alphabet)

	vocabulary := []string{}
	vocabulary = append(vocabulary, "<|endoftext|>")
	vocabulary = append(vocabulary, alphabet...)

	splits := map[string][]string{}
	for _, word := range words {
		splits[word] = []string{}
		for _, letter := range word {
			splits[word] = append(splits[word], string(letter))
		}
	}

	merges, mergeOrder, _ := t.mergeN(vocabulary, splits, words, wordFreqs, 50)
	t.merges = merges
	t.mergeOrder = mergeOrder
}

func (t *BpeTokenizer) computePairFreqs(words []string, wordFreqs map[string]int, splits map[string][]string) ([]StringPair, map[StringPair]int) {

	pairs := []StringPair{}
	pairFreqs := map[StringPair]int{}

	for _, word := range words {
		freq := wordFreqs[word]
		split := splits[word]
		if len(split) == 1 {
			continue
		}
		for i := 0; i < len(split)-1; i++ {
			pair := StringPair{split[i], split[i+1]}
			_, found := pairFreqs[pair]
			if found {
				pairFreqs[pair] += freq
			} else {
				pairs = append(pairs, pair)
				pairFreqs[pair] = freq
			}
		}
	}
	return pairs, pairFreqs
}

func (t *BpeTokenizer) mergePair(a string, b string, words []string, wordFreqs map[string]int, splits map[string][]string) map[string][]string {
	for _, word := range words {
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

func (t *BpeTokenizer) mergeN(vocabulary []string, splits map[string][]string, words []string, wordFreqs map[string]int, vocabularySize int) (map[StringPair]string, []StringPair, []string) {

	merges := map[StringPair]string{}
	mergeOrder := []StringPair{}

	for len(vocabulary) < vocabularySize {
		pairs, pairFreqs := t.computePairFreqs(words, wordFreqs, splits)

		bestPair := StringPair{}
		maxFreq := 0

		for _, pair := range pairs {
			freq := pairFreqs[pair]
			if maxFreq < freq {
				bestPair = pair
				maxFreq = freq
			}
		}

		a := bestPair.first
		b := bestPair.second

		splits = t.mergePair(a, b, words, wordFreqs, splits)
		merges[bestPair] = a + b
		mergeOrder = append(mergeOrder, bestPair)

		// fmt.Printf("%s: %s (%d)\n", bestPair, a+b, maxFreq)
		vocabulary = append(vocabulary, a+b)
	}

	return merges, mergeOrder, vocabulary
}

func (t *BpeTokenizer) Tokenize(text string) []string {

	tok := CreateSimpleTokenizer()
	tokens := tok.Tokenize(text)
	words := AddBeginToken(tokens, "Ġ")

	splits := map[string][]string{}
	for _, word := range words {
		splits[word] = []string{}
		for _, letter := range word {
			splits[word] = append(splits[word], string(letter))
		}
	}

	for _, pair := range t.mergeOrder {
		merge := t.merges[pair]
		for _, word := range words {
			split := splits[word]

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
			splits[word] = split
		}
	}

	result := []string{}
	for _, word := range words {
		split := splits[word]
		result = append(result, split...)
	}
	return result
}
