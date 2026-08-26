package lev

type Weights struct {
	Insert     float64
	Delete     float64
	Substitute float64
}

func DefaultWeights() Weights {
	return Weights{Insert: 1, Delete: 1, Substitute: 1}
}

func WeightedDistance(a, b string, w Weights) float64 {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 {
		return float64(lb) * w.Insert
	}
	if lb == 0 {
		return float64(la) * w.Delete
	}

	prev := make([]float64, lb+1)
	curr := make([]float64, lb+1)
	for j := 0; j <= lb; j++ {
		prev[j] = float64(j) * w.Insert
	}
	for i := 1; i <= la; i++ {
		curr[0] = float64(i) * w.Delete
		for j := 1; j <= lb; j++ {
			subCost := 0.0
			if ra[i-1] != rb[j-1] {
				subCost = w.Substitute
			}
			curr[j] = min3Float(
				curr[j-1]+w.Insert,
				prev[j]+w.Delete,
				prev[j-1]+subCost,
			)
		}
		prev, curr = curr, prev
	}
	return prev[lb]
}

func WeightedNormalized(a, b string, w Weights) float64 {
	d := WeightedDistance(a, b, w)
	la := float64(len([]rune(a)))
	lb := float64(len([]rune(b)))
	maxDist := la*w.Delete + lb*w.Insert
	minLen := la
	if lb < minLen {
		minLen = lb
	}
	altMax := minLen*w.Substitute + (la-minLen)*w.Delete + (lb-minLen)*w.Insert
	if altMax < maxDist {
		maxDist = altMax
	}
	if maxDist == 0 {
		return 1.0
	}
	sim := 1.0 - d/maxDist
	if sim < 0 {
		sim = 0
	}
	return sim
}

func Prefix(a, b string) float64 {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 && lb == 0 {
		return 1.0
	}
	if la == 0 || lb == 0 {
		return 0.0
	}

	prefixLen := 0
	minLen := la
	if lb < minLen {
		minLen = lb
	}
	for i := 0; i < minLen; i++ {
		if ra[i] == rb[i] {
			prefixLen++
		} else {
			break
		}
	}

	maxLen := la
	if lb > maxLen {
		maxLen = lb
	}
	prefixScore := float64(prefixLen) / float64(maxLen)

	levSim := Normalized(a, b)
	return 0.6*levSim + 0.4*prefixScore
}

func min3Float(a, b, c float64) float64 {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
