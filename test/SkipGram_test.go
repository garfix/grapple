package test

import (
	"fmt"
	"testing"

	"github.com/garfix/grapple/src/embedding"
)

func TestSkipGram(t *testing.T) {

	skipGram := embedding.CreateSkipGram(10000, 300, 2, 0.5)

	// the black cat sat on the couch and the brown dog slept on the sat
	input := []int{1, 2, 3, 4, 5, 1, 6, 7, 1, 8, 9, 10, 5, 1, 11}

	skipGram.Train(input)

	// expected := []float64{2.0, 3.1, 5.4}
	blackFeatures := skipGram.GetWordFeatures(2)
	brownFeatures := skipGram.GetWordFeatures(8)
	satFeatures := skipGram.GetWordFeatures(4)

	blackBrownSimilarity := embedding.CalculateCosineSimilarity(blackFeatures, brownFeatures)
	blackSatSimilarity := embedding.CalculateCosineSimilarity(blackFeatures, satFeatures)
	brownSatSimilarity := embedding.CalculateCosineSimilarity(brownFeatures, satFeatures)

	fmt.Printf("black-brown: %v\n", blackBrownSimilarity)
	fmt.Printf("black-sat: %v\n", blackSatSimilarity)
	fmt.Printf("brown-sat: %v\n", brownSatSimilarity)

	t.Errorf("fail")

	// for i := range expected {
	// 	if output[i] != expected[i] {
	// 		t.Errorf("expected: %v\ngot   : %v", expected, output)
	// 		break
	// 	}
	// }

}
