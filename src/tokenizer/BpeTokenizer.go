package tokenizer

import (
	"fmt"
	"sort"

	"github.com/garfix/grapple/src/utils"
)

type BpeTokenizer struct {
	merges map[StringPair]string
}

func CreateBpeTokenizer() *BpeTokenizer {
	return &BpeTokenizer{
		merges: map[StringPair]string{},
	}
}

func (t *BpeTokenizer) Train(corpus []string) {

	tok := CreateSimpleTokenizer()

	wordFreqs := map[string]int{}
	for _, text := range corpus {
		tokens := tok.Tokenize(text)
		words := AddBeginToken(tokens, "Ġ")
		for _, word := range words {
			freq, found := wordFreqs[word]
			if found {
				freq++
			} else {
				freq = 1
			}
			wordFreqs[word] = freq
		}
	}

	alphabet := []string{}

	for word := range wordFreqs {
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

	println()

	// for _, a := range vocabulary {
	// 	print(a)
	// }

	splits := map[string][]string{}
	for word := range wordFreqs {
		splits[word] = []string{}
		for _, letter := range word {
			splits[word] = append(splits[word], string(letter))
		}
	}

	// pairFreqs := utils.ComputePairFreqs(wordFreqs, splits)

	// for k, v := range pairFreqs {
	// 	println(k, v)
	// }

	// bestPair := ""
	// maxFreq := 0

	// for pair, freq := range pairFreqs {
	// 	if maxFreq < freq {
	// 		bestPair = pair
	// 		maxFreq = freq
	// 	}
	// }

	// println(bestPair, maxFreq)

	// splits = utils.MergePair("Ġ", "t", wordFreqs, splits)

	// fmt.Printf("%s\n", vocabulary)

	merges, vocabulary := t.mergeN(vocabulary, splits, wordFreqs, 50)

	t.merges = merges

	// fmt.Printf("%s\n", vocabulary)

	// tokens := utils.Tokenize("This is not a token.", merges)

	// for s, t := range merges {
	// 	fmt.Printf("%s: %s\n", s, t)
	// }

	// println()

	// for _, token := range tokens {
	// 	print(token + " ")
	// }

}

func (t *BpeTokenizer) computePairFreqs(wordFreqs map[string]int, splits map[string][]string) map[StringPair]int {
	pairFreqs := map[StringPair]int{}
	for word, freq := range wordFreqs {
		split := splits[word]
		if len(split) == 1 {
			continue
		}
		for i := 0; i < len(split)-1; i++ {
			pair := StringPair{split[i], split[i+1]}
			// pairFreqs = pairFreqs.add(pair, freq)
			_, found := pairFreqs[pair]
			if found {
				pairFreqs[pair] += freq
			} else {
				pairFreqs[pair] = freq
			}
		}
	}
	return pairFreqs
}

func (t *BpeTokenizer) mergePair(a string, b string, wordFreqs map[string]int, splits map[string][]string) map[string][]string {
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

func (t *BpeTokenizer) mergeN(vocabulary []string, splits map[string][]string, wordFreqs map[string]int, vocabularySize int) (map[StringPair]string, []string) {

	merges := map[StringPair]string{}

	for len(vocabulary) < vocabularySize {
		pairFreqs := t.computePairFreqs(wordFreqs, splits)

		bestPair := StringPair{}
		maxFreq := 0

		for pair, freq := range pairFreqs {
			if maxFreq < freq {
				bestPair = pair
				maxFreq = freq
			}
		}

		a := bestPair.first
		b := bestPair.second

		splits = t.mergePair(a, b, wordFreqs, splits)
		merges[bestPair] = a + b

		fmt.Printf("%s: %s (%d)\n", bestPair, a+b, maxFreq)
		vocabulary = append(vocabulary, a+b)
	}

	return merges, vocabulary
}

func (t *BpeTokenizer) Tokenize(text string) []string {
	// preTokenizeResult = tokenizer._tokenizer.pre_tokenizer.pre_tokenize_str(text)
	// pre_tokenized_text = [word for word, offset in pre_tokenize_result]

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

	// splits = [[l for l in word] for word in pre_tokenized_text]
	for pair, merge := range t.merges {
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
