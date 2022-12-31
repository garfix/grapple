package embedding

import "math"

type SkipGram struct {
	input        *InputLayer
	hidden       *HiddenLayer
	output       *OutputLayer
	wordCount    int
	featureCount int
	windowSize   int
}

func CreateSkipGram(wordCount int, featureCount int, windowSize int) *SkipGram {
	return &SkipGram{
		input:        CreateInputLayer(wordCount),
		hidden:       CreateHiddenLayer(wordCount, featureCount),
		output:       CreateOutputLayer(wordCount, featureCount),
		featureCount: featureCount,
		wordCount:    wordCount,
		windowSize:   windowSize,
	}
}

func (sg *SkipGram) Train(wordIndexes []int) {

	for i, wordIndex1 := range wordIndexes {
		for j := i - sg.windowSize; j <= i+sg.windowSize; j++ {
			if j == i {
				continue
			}
			if j < 0 {
				continue
			}
			if j >= len(wordIndexes) {
				continue
			}

			wordIndex2 := wordIndexes[j]

			sg.trainPair(wordIndex1, wordIndex2)

		}

	}
}

// http://mccormickml.com/2016/04/19/word2vec-tutorial-the-skip-gram-model/
// https://medium.datadriveninvestor.com/word2vec-skip-gram-model-explained-383fa6ddc4ae
func (sg *SkipGram) trainPair(wordIndex1 int, wordIndex2 int) {
	// set input layer
	sg.input.Set(wordIndex1)

	// get current hidden features for word
	hiddenValues := sg.hidden.matrix[sg.input.wordIndex]

	// set output layer
	outputValues := make([]float64, sg.wordCount)
	summedValue := 0.0
	for i := 0; i < sg.wordCount; i++ {
		product := dotProduct(hiddenValues, sg.output.weights[sg.input.wordIndex])
		value := math.Exp(product)
		summedValue += value
		outputValues[i] = value
	}
	for i := 0; i < sg.wordCount; i++ {
		// softmax: exp(x) / sum (exp(x))
		sg.output.outputVector[i] = outputValues[i] / summedValue
	}
	for i := 0; i < sg.wordCount; i++ {
		target := 0.0
		if i == wordIndex2 {
			target = 1.0
		}
		activation := sg.output.outputVector[i]

		// calculate loss
		// loss := crossEntropyLoss(expected, actual)
		error := target - activation

		// back-propagation
		delta := error * activation * (1.0 - activation)

	}
}
