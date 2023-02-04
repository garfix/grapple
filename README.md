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

## Transformer implementations

https://www.tensorflow.org/text/tutorials/transformer#create_the_transformer

https://pytorch.org/docs/master/generated/torch.nn.Transformer.html

thanks to https://github.com/huggingface/transformers/issues/4817
