package embedding

type InputLayer struct {
	wordIndex int
}

func CreateInputLayer(wordCount int) *InputLayer {
	return &InputLayer{
		wordIndex: 0,
	}
}

func (in *InputLayer) Set(wordIndex int) {
	in.wordIndex = wordIndex
}
