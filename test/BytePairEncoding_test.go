package test

import (
	"garfixia.com/transformer/embedding"
	"testing"
)

func TestBytePairEncoding(t *testing.T) {

	strings := []string{
		"Pen Penapple Apple Pen",
	}

	encoding := embedding.CreateBytePairEncoding()
	encoding.Encode(strings, 5)
	t.Error("fail")
}
