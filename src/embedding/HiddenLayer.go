package embedding

import "math/rand"

type HiddenLayer struct {
	weights [][]float64
	values  []float64
}

func CreateHiddenLayer(wordCount int, featureCount int) *HiddenLayer {

	weights := make([][]float64, wordCount)
	for i := range weights {
		weights[i] = make([]float64, featureCount)
	}

	// initialize weights with random numbers between 0 and 0.1
	for i := range weights {
		for j := range weights[i] {
			weights[i][j] = rand.Float64() / 10.0
		}
	}

	return &HiddenLayer{
		weights: weights,
		values:  make([]float64, featureCount),
	}
}
