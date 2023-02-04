package test

import (
	"testing"

	"github.com/garfix/grapple/src/nn"
)

func TestNeuralNetwork(t *testing.T) {
	// Blog post example:

	training_sets := [][][]float64{
		{{0.0, 0.0}, {0.0}},
		{{0.0, 1.0}, {1.0}},
		{{1.0, 0.0}, {1.0}},
		{{1.1, 1.1}, {0.0}},
	}

	// note: hidden layer would need 3 nodes without bias updating
	neuralNetwork2 := nn.CreateThreeLayerNetwork(2, 2, 1, 0.35, 0.6, 0.5, 0.01)
	for i := 0; i < 100000; i++ {
		for j := 0; j < 4; j++ {
			neuralNetwork2.Train(training_sets[j][0], training_sets[j][1])
		}
	}
	a := neuralNetwork2.CalculateTotalError(training_sets)
	if a > 0.001 {
		t.Errorf("Expected: 0.000998, got %v", a)
	}

}
