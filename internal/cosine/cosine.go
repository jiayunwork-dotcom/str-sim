package cosine

import (
	"math"
	"strings"
	"unicode"
)

func CharFrequency(a, b string) float64 {
	if a == b {
		return 1.0
	}
	va := charVector(a)
	vb := charVector(b)
	return cosineSim(va, vb)
}

func TokenFrequency(a, b string) float64 {
	if a == b {
		return 1.0
	}
	ta := tokenize(a)
	tb := tokenize(b)
	va := tokenVector(ta)
	vb := tokenVector(tb)
	return cosineSim(va, vb)
}

func TFIDF(a, b string, corpus []string) float64 {
	if a == b {
		return 1.0
	}

	ta := tokenize(a)
	tb := tokenize(b)

	idf := computeIDF(corpus, ta, tb)

	va := tfidfVector(ta, idf)
	vb := tfidfVector(tb, idf)
	return cosineSim(va, vb)
}

func charVector(s string) map[string]float64 {
	v := make(map[string]float64)
	for _, r := range s {
		v[string(r)]++
	}
	return v
}

func tokenize(s string) []string {
	var tokens []string
	var current []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			current = append(current, r)
		} else {
			if len(current) > 0 {
				tokens = append(tokens, string(current))
				current = current[:0]
			}
		}
	}
	if len(current) > 0 {
		tokens = append(tokens, string(current))
	}
	return tokens
}

func tokenVector(tokens []string) map[string]float64 {
	v := make(map[string]float64)
	for _, t := range tokens {
		v[t]++
	}
	return v
}

func computeIDF(corpus []string, tokensA, tokensB []string) map[string]float64 {
	df := make(map[string]int)
	nDocs := len(corpus) + 2

	for _, doc := range corpus {
		seen := make(map[string]bool)
		for _, t := range tokenize(doc) {
			if !seen[t] {
				df[t]++
				seen[t] = true
			}
		}
	}
	seenA := make(map[string]bool)
	for _, t := range tokensA {
		if !seenA[t] {
			df[t]++
			seenA[t] = true
		}
	}
	seenB := make(map[string]bool)
	for _, t := range tokensB {
		if !seenB[t] {
			df[t]++
			seenB[t] = true
		}
	}

	idf := make(map[string]float64)
	for term, count := range df {
		idf[term] = math.Log(float64(nDocs) / float64(count+1))
	}
	return idf
}

func tfidfVector(tokens []string, idf map[string]float64) map[string]float64 {
	tf := make(map[string]float64)
	for _, t := range tokens {
		tf[t]++
	}
	v := make(map[string]float64, len(tf))
	for term, count := range tf {
		weight := idf[term]
		if weight == 0 {
			weight = 1
		}
		v[term] = count * weight
	}
	return v
}

func cosineSim(a, b map[string]float64) float64 {
	var dot, normA, normB float64
	for k, va := range a {
		normA += va * va
		if vb, ok := b[k]; ok {
			dot += va * vb
		}
	}
	for _, vb := range b {
		normB += vb * vb
	}
	if normA == 0 || normB == 0 {
		return 0.0
	}
	sim := dot / (math.Sqrt(normA) * math.Sqrt(normB))
	if sim > 1.0 {
		sim = 1.0
	}
	return sim
}
