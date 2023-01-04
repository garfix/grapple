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

func (sg *SkipGram) trainPair(inputWordIndex int, outputWordIndex int) {

	sg.updateInputValues(inputWordIndex)
	sg.updateHiddenValues()
	sg.updateOutputValues()
	sg.propagateBackOutputErrors(outputWordIndex)
	sg.propagateBackHiddenErrors(outputWordIndex)
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
		product := DotProduct(hiddenValues, sg.output.weights[i])
		value := math.Exp(product)
		summedValue += value
		sg.output.values[i] = value
	}
	for i := 0; i < sg.wordCount; i++ {
		// softmax: exp(x) / sum (exp(x))
		sg.output.values[i] /= summedValue
	}
}

func (sg *SkipGram) propagateBackOutputErrors(ouputWordIndex int) {
	for wordIndex := 0; wordIndex < sg.wordCount; wordIndex++ {

		delta := sg.calculateDelta(wordIndex, ouputWordIndex)

		for featureIndex := 0; featureIndex < sg.featureCount; featureIndex++ {
			weight := sg.output.weights[wordIndex][featureIndex]
			hiddenValue := sg.hidden.values[featureIndex]
			newWeight := weight - sg.learningRate*(hiddenValue*delta)
			sg.output.weights[wordIndex][featureIndex] = newWeight
		}
	}
}

func (sg *SkipGram) propagateBackHiddenErrors(outputWordIndex int) {
	for wordIndex := 0; wordIndex < sg.wordCount; wordIndex++ {

		delta := sg.calculateDelta(wordIndex, outputWordIndex)
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

func (sg *SkipGram) calculateDelta(wordIndex int, outputWordIndex int) float64 {
	actual := 0.0
	if wordIndex == outputWordIndex {
		actual = 1.0
	}
	predicted := sg.output.values[wordIndex]
	delta := predicted - actual
	return delta
}

func (sg *SkipGram) GetWordFeatures(wordIndex int) []float64 {
	return sg.hidden.weights[wordIndex]
}
