package test

import (
	"testing"

	"github.com/garfix/grapple/src/embedding"
	"github.com/garfix/grapple/src/tokenizer"
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

	bpeTok := tokenizer.CreateBpeTokenizer()
	bpeTok.Train(corpus)

	tokens := bpeTok.Tokenize("This is not a token.")

	println()

	for _, token := range tokens {
		print(token + " ")
	}

	println()

	t.Error("err!")

}
