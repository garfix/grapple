package data

type TokenFrequency struct {
	tokens    []string
	frequency int
}

func NewTokenFrequency(tokens []string, frequency int) TokenFrequency {
	return TokenFrequency{
		tokens:    tokens,
		frequency: frequency,
	}
}
