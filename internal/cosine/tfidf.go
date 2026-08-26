package cosine

import "math"

func BM25(a, b string, corpus []string, k1, bParam float64) float64 {
	if a == b {
		return 1.0
	}
	if k1 <= 0 {
		k1 = 1.2
	}
	if bParam < 0 || bParam > 1 {
		bParam = 0.75
	}

	tokensA := tokenize(a)
	tokensB := tokenize(b)

	if len(tokensA) == 0 || len(tokensB) == 0 {
		return 0.0
	}

	allDocs := append(corpus, a, b)
	nDocs := float64(len(allDocs))

	var totalLen float64
	for _, doc := range allDocs {
		totalLen += float64(len(tokenize(doc)))
	}
	avgDL := totalLen / nDocs

	df := make(map[string]int)
	for _, doc := range allDocs {
		seen := make(map[string]bool)
		for _, t := range tokenize(doc) {
			if !seen[t] {
				df[t]++
				seen[t] = true
			}
		}
	}

	idf := func(term string) float64 {
		d := float64(df[term])
		return math.Log((nDocs - d + 0.5) / (d + 0.5))
	}

	tfB := make(map[string]float64)
	for _, t := range tokensB {
		tfB[t]++
	}
	dlB := float64(len(tokensB))

	var score float64
	seen := make(map[string]bool)
	for _, t := range tokensA {
		if seen[t] {
			continue
		}
		seen[t] = true
		tf := tfB[t]
		if tf == 0 {
			continue
		}
		idfVal := idf(t)
		if idfVal < 0 {
			idfVal = 0
		}
		denom := tf + k1*(1-bParam+bParam*dlB/avgDL)
		score += idfVal * (tf * (k1 + 1)) / denom
	}

	var selfScore float64
	seen2 := make(map[string]bool)
	for _, t := range tokensB {
		if seen2[t] {
			continue
		}
		seen2[t] = true
		tf := tfB[t]
		idfVal := idf(t)
		if idfVal < 0 {
			idfVal = 0
		}
		denom := tf + k1*(1-bParam+bParam*dlB/avgDL)
		selfScore += idfVal * (tf * (k1 + 1)) / denom
	}

	if selfScore == 0 {
		return 0.0
	}
	sim := score / selfScore
	if sim > 1.0 {
		sim = 1.0
	}
	if sim < 0 {
		sim = 0
	}
	return sim
}

func JaccardOnTokens(a, b string) float64 {
	if a == b {
		return 1.0
	}
	ta := tokenSet(a)
	tb := tokenSet(b)
	if len(ta) == 0 && len(tb) == 0 {
		return 1.0
	}

	inter := 0
	for t := range ta {
		if tb[t] {
			inter++
		}
	}
	union := len(ta)
	for t := range tb {
		if !ta[t] {
			union++
		}
	}
	if union == 0 {
		return 0.0
	}
	return float64(inter) / float64(union)
}

func tokenSet(s string) map[string]bool {
	set := make(map[string]bool)
	for _, t := range tokenize(s) {
		set[t] = true
	}
	return set
}
