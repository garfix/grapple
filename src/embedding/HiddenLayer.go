package embedding

import "math/rand"

type HiddenLayer struct {
	matrix [][]float64
}

func CreateHiddenLayer(wordCount int, featureCount int) *HiddenLayer {

	matrix := make([][]float64, wordCount)
	for i := range matrix {
		matrix[i] = make([]float64, featureCount)
	}

	// initialize with random numbers between 0 and 1
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = rand.Float64()
		}
	}

	return &HiddenLayer{
		matrix: matrix,
	}
}
