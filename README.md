# Grapple

Learning to build a transformer, in Go.

## Pre-tokenization

Turn a sentence into an array of words, and add spacing-tokens.

I implemented a simple custom tokenizer and a simple custom add-token-to-begin function.

## Tokenization

There are three types of tokenizers:

- byte-pair encoding
- word-piece
- unigram

I implemented byte-pair encoding

## Embedding

Turns a word into an n-dimensional vector, intended to capture semantic similarity

Embedders:

- skip-gram
- continuous bag of words (CBOW)

I will implement skip-gram
