package test

import (
	"fmt"
	"testing"

	"github.com/garfix/grapple/src/embedding"
)

func TestSkipGram(t *testing.T) {

	// rand.Seed(time.Now().UnixNano())

	skipGram := embedding.CreateSkipGram(10000, 300, 2, 1.0)

	// the black cat sat on the couch and the brown dog slept on the rug
	input := []int{1, 2, 3, 4, 5, 1, 6, 7, 1, 8, 9, 10, 5, 1, 11}

	// for i := 0; i < 1; i++ {
	skipGram.Train(input)
	// }

	blackFeatures := skipGram.GetWordFeatures(2)
	brownFeatures := skipGram.GetWordFeatures(8)
	satFeatures := skipGram.GetWordFeatures(4)

	blackBrownSimilarity := embedding.CalculateCosineSimilarity(blackFeatures, brownFeatures)
	blackSatSimilarity := embedding.CalculateCosineSimilarity(blackFeatures, satFeatures)
	brownSatSimilarity := embedding.CalculateCosineSimilarity(brownFeatures, satFeatures)

	fmt.Printf("black-brown: %v\n", blackBrownSimilarity)
	fmt.Printf("black-sat: %v\n", blackSatSimilarity)
	fmt.Printf("brown-sat: %v\n", brownSatSimilarity)

	// fmt.Printf("black-brown: %v\n", embedding.DotProduct(blackFeatures, brownFeatures))
	// fmt.Printf("black-sat: %v\n", embedding.DotProduct(blackFeatures, satFeatures))
	// fmt.Printf("brown-sat: %v\n", embedding.DotProduct(brownFeatures, satFeatures))

	// fmt.Printf("black: %v\n\n", blackFeatures)
	// fmt.Printf("brown: %v\n\n", brownFeatures)
	// fmt.Printf("sat  : %v\n\n", satFeatures)

	t.Errorf("fail")

	// for i := range expected {
	// 	if output[i] != expected[i] {
	// 		t.Errorf("expected: %v\ngot   : %v", expected, output)
	// 		break
	// 	}
	// }

}
