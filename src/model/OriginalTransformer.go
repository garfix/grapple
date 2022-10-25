package model

import "github.com/garfix/grapple/src/embedding"

type OriginalTransformer struct {
}

func CreateOriginalTransformer() *OriginalTransformer {

	embedding.CreateBytePairEncoding()

	transformer := OriginalTransformer{}
	return &transformer
}
