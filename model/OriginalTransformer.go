package model

import "garfixia.com/transformer/embedding"

type OriginalTransformer struct {
}

func CreateOriginalTransformer() *OriginalTransformer {

	e := embedding.CreateBytePairEncoding()

	transformer := OriginalTransformer{}
	return &transformer
}
