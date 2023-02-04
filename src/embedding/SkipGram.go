package embedding

import (
	"github.com/garfix/grapple/src/nn"
)

type SkipGram struct {
	// input        *InputLayer
	// hidden       *HiddenLayer
	// output       *OutputLayer
	wordCount    int
	featureCount int
	windowSize   int
	// learningRate float64
	nn *nn.ThreeLayerNetwork
}

func CreateSkipGram(wordCount int, featureCount int, windowSize int, learningRate float64) *SkipGram {
	return &SkipGram{
		windowSize: windowSize,
		nn:         nn.CreateThreeLayerNetwork(wordCount, featureCount, wordCount, 0.0, 0.0, learningRate, 0.01),
		// input:        CreateInputLayer(wordCount),
		// hidden:       CreateHiddenLayer(wordCount, featureCount),
		// output:       CreateOutputLayer(wordCount, featureCount),
		featureCount: featureCount,
		wordCount:    wordCount,
		// windowSize:   windowSize,
		// learningRate: learningrate,
	}
}

func (sg *SkipGram) Train(wordIndexes []int) {

	for i, wordIndex1 := range wordIndexes {
		for j := i - sg.windowSize; j <= i+sg.windowSize; j++ {
			if j == i || j < 0 || j >= len(wordIndexes) {
				continue
			}

			wordIndex2 := wordIndexes[j]
			sg.trainPair(wordIndex1, wordIndex2)
		}
	}
}

func (sg *SkipGram) trainPair(inputWordIndex int, outputWordIndex int) {

	inputs := []float64{}
	outputs := []float64{}

	for i := 0; i < sg.wordCount; i++ {
		if i == inputWordIndex {
			inputs = append(inputs, 1.0)
		} else {
			inputs = append(inputs, 0.0)
		}
	}

	for i := 0; i < sg.wordCount; i++ {
		if i == outputWordIndex {
			outputs = append(outputs, 1.0)
		} else {
			outputs = append(outputs, 0.0)
		}
	}

	sg.nn.Train(inputs, outputs)

	// 	sg.updateInputValues(inputWordIndex)
	// 	sg.updateHiddenValues()
	// 	sg.updateOutputValues()
	// 	// sg.propagateBackOutputErrors(outputWordIndex)
	// 	// sg.propagateBackHiddenErrors(outputWordIndex)
	// 	// sg.propagateBackErrors(outputWordIndex)
}

// func (sg *SkipGram) updateInputValues(wordIndex int) {
// 	sg.input.Set(wordIndex)
// }

// func (sg *SkipGram) updateHiddenValues() {
// 	hiddenValues := sg.hidden.weights[sg.input.wordIndex]
// 	sg.hidden.values = hiddenValues
// }

// func (sg *SkipGram) updateOutputValues() {
// 	hiddenValues := sg.hidden.values
// 	summedValue := 0.0
// 	for i := 0; i < sg.wordCount; i++ {
// 		product := DotProduct(hiddenValues, sg.output.weights[i])
// 		value := math.Exp(product)
// 		summedValue += value
// 		sg.output.values[i] = value
// 	}
// 	for i := 0; i < sg.wordCount; i++ {
// 		// softmax: exp(x) / sum (exp(x))
// 		sg.output.values[i] /= summedValue
// 	}
// }

// func (sg *SkipGram) propagateBackErrors(ouputWordIndex int) {

// 	pd_errors_wrt_output_neuron_total_net_input := []float64{}
// 	for wordIndex := 0; wordIndex < sg.wordCount; wordIndex++ {

// 		output := sg.output.values[wordIndex]
// 		error := calculate_pd_error_wrt_total_net_input(output)
// 		pd_errors_wrt_output_neuron_total_net_input = append(pd_errors_wrt_output_neuron_total_net_input, error)
// 	}
// }

// func (sg *SkipGram) calculate_pd_error_wrt_total_net_input(output float64) {

// }

// func (sg *SkipGram) propagateBackOutputErrors(ouputWordIndex int) {
// 	for wordIndex := 0; wordIndex < sg.wordCount; wordIndex++ {

// 		delta := sg.calculateDelta(wordIndex, ouputWordIndex)
// 		t := sg.output.values[wordIndex] * (1.0 - sg.output.values[wordIndex])

// 		for featureIndex := 0; featureIndex < sg.featureCount; featureIndex++ {
// 			weight := sg.output.weights[wordIndex][featureIndex]
// 			hiddenValue := sg.hidden.values[featureIndex]
// 			newWeight := weight - sg.learningRate*(delta*t*hiddenValue)
// 			sg.output.weights[wordIndex][featureIndex] = newWeight
// 		}
// 	}
// }

// func (sg *SkipGram) propagateBackHiddenErrors(outputWordIndex int) {

// 	// handling only the active input, because the calculations for inactive inputs don't make their weights change
// 	inputValue := 1.0
// 	inWordIndex := sg.input.wordIndex

// 	for outWordIndex := 0; outWordIndex < sg.wordCount; outWordIndex++ {

// 		delta := sg.calculateDelta(outWordIndex, outputWordIndex)
// 		t1 := sg.output.values[outWordIndex] * (1.0 - sg.output.values[outWordIndex])

// 		for featureIndex := 0; featureIndex < sg.featureCount; featureIndex++ {

// 			t2 := sg.hidden.values[featureIndex] * (1.0 - sg.hidden.values[featureIndex])

// 			weight := sg.hidden.weights[inWordIndex][featureIndex]
// 			outputWeight := sg.output.weights[outWordIndex][featureIndex]
// 			newWeight := weight - sg.learningRate*(delta*t1*outputWeight*t2*inputValue)
// 			sg.hidden.weights[inWordIndex][featureIndex] = newWeight
// 		}
// 	}
// }

// func (sg *SkipGram) calculateDelta(wordIndex int, outputWordIndex int) float64 {
// 	actual := 0.0
// 	if wordIndex == outputWordIndex {
// 		actual = 1.0
// 	}
// 	predicted := sg.output.values[wordIndex]
// 	delta := predicted - actual
// 	return delta
// }

// func (sg *SkipGram) GetWordFeatures(wordIndex int) []float64 {
// 	return sg.hidden.weights[wordIndex]
// }

func (sg *SkipGram) GetWordFeatures(wordIndex int) []float64 {
	return sg.nn.GetHiddenLayerWeights(wordIndex)
}
