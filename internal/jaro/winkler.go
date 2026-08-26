package jaro

type WinklerParams struct {
	PrefixScale    float64
	MaxPrefix      int
	BoostThreshold float64
}

func DefaultWinklerParams() WinklerParams {
	return WinklerParams{
		PrefixScale:    0.1,
		MaxPrefix:      4,
		BoostThreshold: 0.7,
	}
}

func JaroWinklerConfigured(a, b string, params WinklerParams) float64 {
	j := Jaro(a, b)
	if j <= 0 {
		return 0
	}
	if j < params.BoostThreshold {
		return j
	}

	maxP := params.MaxPrefix
	if maxP > len(a) {
		maxP = len(a)
	}
	if maxP > len(b) {
		maxP = len(b)
	}
	prefix := 0
	for i := 0; i < maxP; i++ {
		if a[i] == b[i] {
			prefix++
		} else {
			break
		}
	}

	scale := params.PrefixScale
	if scale > 0.25 {
		scale = 0.25
	}

	r := j + float64(prefix)*scale*(1-j)
	if r > 1.0 {
		r = 1.0
	}
	return r
}

func SortedJaro(a, b string) float64 {
	return Jaro(sortRunes(a), sortRunes(b))
}

func SortedJaroWinkler(a, b string) float64 {
	return JaroWinkler(sortRunes(a), sortRunes(b))
}

func sortRunes(s string) string {
	runes := []rune(s)
	for i := 1; i < len(runes); i++ {
		for j := i; j > 0 && runes[j] < runes[j-1]; j-- {
			runes[j], runes[j-1] = runes[j-1], runes[j]
		}
	}
	return string(runes)
}

func PartialJaro(a, b string) float64 {
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) > len(rb) {
		ra, rb = rb, ra
	}
	if len(ra) == 0 {
		if len(rb) == 0 {
			return 1.0
		}
		return 0.0
	}

	short := string(ra)
	bestScore := 0.0
	windowSize := len(ra)

	for i := 0; i <= len(rb)-windowSize; i++ {
		window := string(rb[i : i+windowSize])
		score := Jaro(short, window)
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}
