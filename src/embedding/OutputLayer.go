package embedding

import "math/rand"

type OutputLayer struct {
	weights      [][]float64
	outputVector []float64
}

func CreateOutputLayer(wordCount int, featureCount int) *OutputLayer {

	weights := make([][]float64, wordCount)
	for i := range weights {
		weights[i] = make([]float64, featureCount)
	}

	// initialize with random numbers between 0 and 1
	for i := range weights {
		for j := range weights[i] {
			weights[i][j] = rand.Float64()
		}
	}

	return &OutputLayer{
		weights:      weights,
		outputVector: make([]float64, wordCount),
	}
}
