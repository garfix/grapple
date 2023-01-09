package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/garfix/grapple/src/embedding"
)

type block struct {
	word1 string
	word2 string
	sim   float64
}

type BySim []block

func (a BySim) Len() int           { return len(a) }
func (a BySim) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySim) Less(i, j int) bool { return a[i].sim < a[j].sim }

func TestSkipGram(t *testing.T) {

	// rand.Seed(time.Now().UnixNano())

	skipGram := embedding.CreateSkipGram(20, 300, 2, 0.5)

	words := map[int]string{
		1:  "the",
		2:  "black",
		3:  "cat",
		4:  "sat",
		5:  "on",
		6:  "couch",
		7:  "and",
		8:  "brown",
		9:  "dog",
		10: "slept",
		11: "rug",
	}

	// the black cat sat on the couch and the brown dog slept on the rug
	input := []int{1, 2, 3, 4, 5, 1, 6, 7, 1, 8, 9, 10, 5, 1, 11}

	for i := 0; i < 10000; i++ {
		skipGram.Train(input)
	}

	blocks := []block{}

	for i := 0; i < 11; i++ {
		for j := i + 1; j < 11; j++ {
			f1 := skipGram.GetWordFeatures(i + 1)
			f2 := skipGram.GetWordFeatures(j + 1)
			sim := embedding.CalculateCosineSimilarity(f1, f2)
			blocks = append(blocks, block{words[i+1], words[j+1], sim})
		}
	}

	sort.Sort(BySim(blocks))

	for i := 0; i < len(blocks); i++ {
		block := blocks[len(blocks)-1-i]
		fmt.Printf("%s - %s: %f\n", block.word1, block.word2, block.sim)
	}

	// blackFeatures := skipGram.GetWordFeatures(2)
	// brownFeatures := skipGram.GetWordFeatures(8)
	// satFeatures := skipGram.GetWordFeatures(4)
	// andFeatures := skipGram.GetWordFeatures(7)
	// onFeatures := skipGram.GetWordFeatures(5)
	// theFeatures := skipGram.GetWordFeatures(1)

	// blackBrownSimilarity := embedding.CalculateCosineSimilarity(blackFeatures, brownFeatures)
	// blackSatSimilarity := embedding.CalculateCosineSimilarity(blackFeatures, satFeatures)
	// brownSatSimilarity := embedding.CalculateCosineSimilarity(brownFeatures, satFeatures)
	// onTheSimilarity := embedding.CalculateCosineSimilarity(onFeatures, theFeatures)

	// fmt.Printf("black-brown: %v\n", blackBrownSimilarity)
	// fmt.Printf("black-sat:   %v\n", blackSatSimilarity)
	// fmt.Printf("brown-sat:   %v\n", brownSatSimilarity)
	// fmt.Printf("on-the:      %v\n", onTheSimilarity)

	// fmt.Println()
	// fmt.Printf("black-brown: %v\n", embedding.DotProduct(blackFeatures, brownFeatures))
	// fmt.Printf("black-sat: %v\n", embedding.DotProduct(blackFeatures, satFeatures))
	// fmt.Printf("brown-sat: %v\n", embedding.DotProduct(brownFeatures, satFeatures))
	// fmt.Printf("on-the:    %v\n", embedding.DotProduct(onFeatures, theFeatures))
	// fmt.Printf("blank-and:    %v\n", embedding.DotProduct(blackFeatures, andFeatures))

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
