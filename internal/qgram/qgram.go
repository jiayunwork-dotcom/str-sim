package qgram

import "math"

func Distance(a, b string, q int) int {
	if q < 1 {
		q = 2
	}
	pa := profile(a, q)
	pb := profile(b, q)

	allKeys := make(map[string]bool)
	for k := range pa {
		allKeys[k] = true
	}
	for k := range pb {
		allKeys[k] = true
	}

	dist := 0
	for k := range allKeys {
		diff := pa[k] - pb[k]
		if diff < 0 {
			diff = -diff
		}
		dist += diff
	}
	return dist
}

func Normalized(a, b string, q int) float64 {
	if a == b {
		return 1.0
	}
	if q < 1 {
		q = 2
	}
	pa := profile(a, q)
	pb := profile(b, q)

	var totalA, totalB int
	for _, c := range pa {
		totalA += c
	}
	for _, c := range pb {
		totalB += c
	}
	total := totalA + totalB
	if total == 0 {
		return 1.0
	}

	dist := 0
	allKeys := make(map[string]bool)
	for k := range pa {
		allKeys[k] = true
	}
	for k := range pb {
		allKeys[k] = true
	}
	for k := range allKeys {
		diff := pa[k] - pb[k]
		if diff < 0 {
			diff = -diff
		}
		dist += diff
	}

	return 1.0 - float64(dist)/float64(total)
}

func Cosine(a, b string, q int) float64 {
	if a == b {
		return 1.0
	}
	if q < 1 {
		q = 2
	}
	pa := profile(a, q)
	pb := profile(b, q)

	var dot, normA, normB float64
	for k, ca := range pa {
		fa := float64(ca)
		normA += fa * fa
		if cb, ok := pb[k]; ok {
			dot += fa * float64(cb)
		}
	}
	for _, cb := range pb {
		fb := float64(cb)
		normB += fb * fb
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

func EditDistanceLowerBound(a, b string, q int) int {
	if q < 1 {
		q = 2
	}
	d := Distance(a, b, q)
	bound := d / (2 * q)
	if d%(2*q) != 0 {
		bound++
	}
	return bound
}

func profile(s string, q int) map[string]int {
	runes := []rune(s)
	p := make(map[string]int)
	if len(runes) < q {
		if len(runes) > 0 {
			p[string(runes)] = 1
		}
		return p
	}
	for i := 0; i <= len(runes)-q; i++ {
		gram := string(runes[i : i+q])
		p[gram]++
	}
	return p
}
