package test

import (
	"strings"
	"testing"

	"github.com/garfix/grapple/src/tokenizer"
)

func TestTokenizerUtilitiesEndToken(t *testing.T) {

	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"Pen", "apple", "."}, []string{"Pen_", "apple_", "._"}},
	}

	for _, test := range tests {
		result := tokenizer.AddEndToken(test.input, "_")
		resultAsString := strings.Join(result, " ")
		expectedAsString := strings.Join(test.expected, " ")
		if resultAsString != expectedAsString {
			t.Error("Expected: " + expectedAsString + "; Got: " + resultAsString)
		}
	}
}
