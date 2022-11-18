package test

import (
	"sort"
	"testing"

	"github.com/garfix/grapple/src/embedding"
	"github.com/garfix/grapple/src/tokenizer"
	"github.com/garfix/grapple/src/utils"
)

func TestAllUniqueCharacters(t *testing.T) {

	encoding := embedding.CreateBytePairEncoding()

	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"Pen Penapple Apple Pen"}, "epPn laA"},
		{[]string{"pen_", "penapple_", "apple_", "pen_"}, "pe_nal"},
	}

	for _, test := range tests {
		vocabulary := encoding.AllUniqueCharacters(test.input)
		if string(vocabulary) != test.expected {
			t.Error("Expected: " + test.expected + "; Got: " + string(vocabulary))
		}
	}
}

func TestBPE(t *testing.T) {

	corpus := []string{
		"This is the Hugging Face Course.",
		"This chapter is about tokenization.",
		"This section shows several tokenizer algorithms.",
		"Hopefully, you will be able to understand how they are trained and generate tokens.",
	}

	tok := tokenizer.CreateSimpleTokenizer()

	wordFreqs := map[string]int{}
	for _, text := range corpus {
		tokens := tok.Tokenize(text)
		words := tokenizer.AddBeginToken(tokens, "Ġ")
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

	// fmt.Print(wordFreqs)

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

	pairFreqs := utils.ComputePairFreqs(wordFreqs, splits)

	// for k, v := range pairFreqs {
	// 	println(k, v)
	// }

	bestPair := ""
	maxFreq := 0

	for pair, freq := range pairFreqs {
		if maxFreq < freq {
			bestPair = pair
			maxFreq = freq
		}
	}

	println(bestPair, maxFreq)

	// splits = utils.MergePair("Ġ", "t", wordFreqs, splits)

	// merges, vocabulary := utils.MergeN(vocabulary, splits, wordFreqs, 50)

	// for s, t := range merges {
	// 	println(s, t)
	// }

	// fmt.Printf("%s", vocabulary)

	// bpe := embedding.CreateBytePairEncoding()
	// bpe.Encode()
	// t.Error("err!")
}
