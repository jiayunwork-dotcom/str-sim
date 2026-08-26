# str-sim

A comprehensive string similarity toolkit implementing multiple distance and
similarity metrics for fuzzy matching, deduplication, and phonetic comparison.

## Features

- **Edit distance**: Levenshtein, Damerau-Levenshtein (true DL + OSA), weighted edit distance, prefix similarity
- **Phonetic**: Soundex, Metaphone, NYSIIS with code-based similarity scoring
- **N-gram**: Jaccard, Dice, Overlap coefficients on character n-gram sets; weighted/positional variants; Q-gram distance
- **Token-based**: Cosine similarity on character/token frequencies; TF-IDF; BM25-inspired scoring
- **Subsequence**: Longest Common Substring, Longest Common Subsequence with normalized similarity
- **Hamming**: Hamming distance (exact + padded), SimHash locality-sensitive hashing, MinHash Jaccard estimation
- **Jaro-Winkler**: Standard + configurable parameters, sorted variants, partial matching
- **Composite**: Multi-algorithm weighted matcher with top-K retrieval and threshold filtering

## Build & Test

```bash
export GOTOOLCHAIN=local CGO_ENABLED=0
go build ./...
go test ./...
```

## Usage

```bash
# Compute similarity
str-sim levenshtein "kitten" "sitting"
str-sim jaro-winkler "Stephen" "Steven"
str-sim ngram-dice "hello" "hallo"

# Check threshold match
str-sim match levenshtein "hello" "hallo" 0.7

# List all algorithms
str-sim list
```

## Package structure

```
internal/
├── lev/        Levenshtein distance + weighted variant + prefix similarity
├── damerau/    Damerau-Levenshtein (true DL + OSA) + alignment/traceback
├── jaro/       Jaro + Jaro-Winkler + configurable/sorted/partial variants
├── ngram/      N-gram set similarity (Jaccard/Dice/Overlap/Weighted/Positional)
├── qgram/      Q-gram frequency distance + edit distance lower bound
├── cosine/     Cosine similarity (char freq, token freq, TF-IDF, BM25)
├── soundex/    Soundex + Metaphone + NYSIIS phonetic encoding
├── hamming/    Hamming distance + SimHash + MinHash
├── lcsubstr/   Longest Common Substring/Subsequence
├── matcher/    Multi-algorithm weighted scorer + top-K + threshold filter
└── sim/        Unified dispatcher (all algorithms accessible by name)
```

## License

MIT
