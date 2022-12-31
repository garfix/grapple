package embedding

// https://en.wikipedia.org/wiki/Dot_product
func dotProduct(v1 []float64, v2 []float64) float64 {
	product := 0.0

	for i := 0; i < len(v1); i++ {
		product += v1[i] * v2[i]
	}

	return product
}
