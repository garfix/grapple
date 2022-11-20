package test

import (
	"strings"
	"testing"

	"github.com/garfix/grapple/src/tokenizer"
)

func TestSimpleTokenizer(t *testing.T) {

	tok := tokenizer.CreateSimpleTokenizer()

	tests := []struct {
		input    string
		expected []string
	}{
		{"Pen, apple", []string{"Pen", ",", " ", "apple"}},
	}

	for _, test := range tests {
		result := tok.Tokenize(test.input)
		resultAsString := strings.Join(result, " ")
		expectedAsString := strings.Join(test.expected, " ")
		if resultAsString != expectedAsString {
			t.Error("Expected: " + expectedAsString + "; Got: " + resultAsString)
		}
	}
}

func TestAddBeginToken(t *testing.T) {

	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"Pen", ",", " ", "apple"}, []string{"Pen", ",", "Ġapple"}},
	}

	for _, test := range tests {
		result := tokenizer.AddBeginToken(test.input, "Ġ")
		resultAsString := strings.Join(result, " ")
		expectedAsString := strings.Join(test.expected, " ")
		if resultAsString != expectedAsString {
			t.Error("Expected: " + expectedAsString + "; Got: " + resultAsString)
		}
	}
}
