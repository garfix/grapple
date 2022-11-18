package tokenizer

// Adds endToken to each of tokens
func AddEndToken(tokens []string, endToken string) []string {
	newTokens := []string{}
	for _, token := range tokens {
		newTokens = append(newTokens, token+endToken)
	}
	return newTokens
}

// Adds endToken to each of tokens
func AddBeginToken(tokens []string, beginToken string) []string {
	newTokens := []string{}
	for _, token := range tokens {
		newTokens = append(newTokens, beginToken+token)
	}
	return newTokens
}
