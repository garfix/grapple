package embedding

import (
	"math"
)

// https://en.wikipedia.org/wiki/Dot_product
func dotProduct(v1 []float64, v2 []float64) float64 {
	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i] * v2[i]
	}
	return sum
}

func CalculateCosineSimilarity(v1 []float64, v2 []float64) float64 {
	return dotProduct(v1, v2) / (calculateMagnitude(v1) * calculateMagnitude(v2))
}

func calculateMagnitude(v []float64) float64 {
	sum := 0.0
	for _, element := range v {
		sum += element * element
	}
	return math.Sqrt(sum)
}
