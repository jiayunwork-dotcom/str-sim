package hamming

func SimHash(s string, n int) uint64 {
	if n < 1 {
		n = 2
	}
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}

	var sums [64]int
	for i := 0; i <= len(runes)-n; i++ {
		gram := string(runes[i : i+n])
		h := fnv1a64(gram)
		for bit := 0; bit < 64; bit++ {
			if h&(1<<uint(bit)) != 0 {
				sums[bit]++
			} else {
				sums[bit]--
			}
		}
	}

	var hash uint64
	for bit := 0; bit < 64; bit++ {
		if sums[bit] > 0 {
			hash |= 1 << uint(bit)
		}
	}
	return hash
}

func SimHashSimilarity(a, b string, n int) float64 {
	ha := SimHash(a, n)
	hb := SimHash(b, n)
	if ha == hb {
		return 1.0
	}
	xor := ha ^ hb
	dist := popcount64(xor)
	return 1.0 - float64(dist)/64.0
}

func popcount64(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func fnv1a64(s string) uint64 {
	const (
		offset64 = 14695981039346656037
		prime64  = 1099511628211
	)
	h := uint64(offset64)
	for _, c := range s {
		h ^= uint64(c)
		h *= prime64
	}
	return h
}

func MinHash(a, b string, n int, k int) float64 {
	if a == b {
		return 1.0
	}
	if k < 1 {
		k = 64
	}

	gramsA := ngramSet(a, n)
	gramsB := ngramSet(b, n)
	if len(gramsA) == 0 && len(gramsB) == 0 {
		return 1.0
	}
	if len(gramsA) == 0 || len(gramsB) == 0 {
		return 0.0
	}

	matches := 0
	for i := 0; i < k; i++ {
		minA := minHash(gramsA, uint64(i))
		minB := minHash(gramsB, uint64(i))
		if minA == minB {
			matches++
		}
	}
	return float64(matches) / float64(k)
}

func ngramSet(s string, n int) []string {
	if n < 1 {
		n = 2
	}
	runes := []rune(s)
	if len(runes) < n {
		if len(runes) > 0 {
			return []string{string(runes)}
		}
		return nil
	}
	var grams []string
	seen := make(map[string]bool)
	for i := 0; i <= len(runes)-n; i++ {
		g := string(runes[i : i+n])
		if !seen[g] {
			grams = append(grams, g)
			seen[g] = true
		}
	}
	return grams
}

func minHash(grams []string, seed uint64) uint64 {
	var minVal uint64 = ^uint64(0)
	for _, g := range grams {
		h := fnv1a64(g) ^ seed
		h = h*6364136223846793005 + 1442695040888963407
		if h < minVal {
			minVal = h
		}
	}
	return minVal
}
