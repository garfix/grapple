package test

import (
	"github.com/garfix/grapple/src/embedding"
	"testing"
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
