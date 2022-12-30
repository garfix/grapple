package model

import "github.com/garfix/grapple/src/encoding"

type OriginalTransformer struct {
}

func CreateOriginalTransformer() *OriginalTransformer {

	encoding.CreateBytePairEncoding()

	transformer := OriginalTransformer{}
	return &transformer
}
