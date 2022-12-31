package embedding

import "math"

type SkipGram struct {
	input        *InputLayer
	hidden       *HiddenLayer
	output       *OutputLayer
	wordCount    int
	featureCount int
	windowSize   int
	learningRate float64
}

func CreateSkipGram(wordCount int, featureCount int, windowSize int, learningrate float64) *SkipGram {
	return &SkipGram{
		input:        CreateInputLayer(wordCount),
		hidden:       CreateHiddenLayer(wordCount, featureCount),
		output:       CreateOutputLayer(wordCount, featureCount),
		featureCount: featureCount,
		wordCount:    wordCount,
		windowSize:   windowSize,
		learningRate: learningrate,
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

// http://mccormickml.com/2016/04/19/word2vec-tutorial-the-skip-gram-model/
// https://medium.datadriveninvestor.com/word2vec-skip-gram-model-explained-383fa6ddc4ae
// https://hmkcode.com/ai/backpropagation-step-by-step/
func (sg *SkipGram) trainPair(inputWordIndex int, expectedWordIndex int) {

	sg.updateInputValues(inputWordIndex)
	sg.updateHiddenValues()
	sg.updateOutputValues()
	sg.propagateBackOutputErrors(expectedWordIndex)
	sg.propagateBackHiddenErrors(expectedWordIndex)
}

func (sg *SkipGram) updateInputValues(wordIndex int) {
	sg.input.Set(wordIndex)
}

func (sg *SkipGram) updateHiddenValues() {
	hiddenValues := sg.hidden.weights[sg.input.wordIndex]
	sg.hidden.values = hiddenValues
}

func (sg *SkipGram) updateOutputValues() {
	hiddenValues := sg.hidden.values
	summedValue := 0.0
	for i := 0; i < sg.wordCount; i++ {
		product := dotProduct(hiddenValues, sg.output.weights[sg.input.wordIndex])
		value := math.Exp(product)
		summedValue += value
		sg.output.values[i] = value
	}
	for i := 0; i < sg.wordCount; i++ {
		// softmax: exp(x) / sum (exp(x))
		sg.output.values[i] /= summedValue
	}
}

func (sg *SkipGram) propagateBackOutputErrors(expectedWordIndex int) {
	for wordIndex := 0; wordIndex < sg.wordCount; wordIndex++ {

		delta := sg.calculateDelta(wordIndex, expectedWordIndex)

		for featureIndex := 0; featureIndex < sg.featureCount; featureIndex++ {
			weight := sg.output.weights[wordIndex][featureIndex]
			hiddenValue := sg.hidden.values[featureIndex]
			newWeight := weight - sg.learningRate*(hiddenValue*delta)
			sg.output.weights[wordIndex][featureIndex] = newWeight
		}
	}
}

func (sg *SkipGram) propagateBackHiddenErrors(expectedWordIndex int) {
	for wordIndex := 0; wordIndex < sg.wordCount; wordIndex++ {

		delta := sg.calculateDelta(wordIndex, expectedWordIndex)
		inputValue := 0.0
		if wordIndex == sg.input.wordIndex {
			inputValue = 1.0
		}

		for featureIndex := 0; featureIndex < sg.featureCount; featureIndex++ {
			weight := sg.hidden.weights[wordIndex][featureIndex]
			outputWeight := sg.output.weights[wordIndex][featureIndex]
			newWeight := weight - sg.learningRate*(inputValue*delta*outputWeight)
			sg.hidden.weights[wordIndex][featureIndex] = newWeight
		}
	}
}

func (sg *SkipGram) calculateDelta(wordIndex int, expectedWordIndex int) float64 {
	predicted := 0.0
	if wordIndex == expectedWordIndex {
		predicted = 1.0
	}
	actual := sg.output.values[wordIndex]
	delta := predicted - actual
	return delta
}
