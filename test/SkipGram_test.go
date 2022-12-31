package test

import (
	"testing"

	"github.com/garfix/grapple/src/embedding"
)

func TestSkipGram(t *testing.T) {

	skipGram := embedding.CreateSkipGram(10000, 300, 2)

	// the black cat sat on the couch and the brown dog slept on the rug
	input := []int{1, 2, 3, 4, 5, 1, 6, 7, 1, 8, 9, 10, 5, 1, 11}

	expected := []float32{2.0, 3.1, 5.4}

	skipGram.Train(input)

	for i := range input {
		if output[i] != expected[i] {
			t.Errorf("Expected at %d: %f, got: %f", i, expected[i], output[i])
			break
		}
	}

}
