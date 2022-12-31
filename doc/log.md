## 2022-12-30

Started working on the skip-gram neural network. It needs to be trained on the words in the vocabulary and produces the values in the hidden layer. It's these values that will be used by the transformer.

## 2022-11-20 

I am using the book "Transformers for Natural Language Processing" to learn about transformers. It doesn't provide much details about the implementation, but gives enough hints to start searching.

I finished a BPE (Byte-Pair Encoding) based tokenizer, based on Python-code by Hugging Face. Porting from Python to Go is simple, except for the fact that Go doesn't store the order of items in a map. I solved this by keeping an extra array of items to accompany the map.

I noticed that there are several ways to approach tokenization and that naming is not very clear. I am assuming tokenization based on characters, not bytes, for the moment.

Next up: embedding. The book "chooses" the skip-gram architecture of word2vec.
