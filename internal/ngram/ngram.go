package ngram

func Profile(s string, n int) map[string]int {
	if n < 1 {
		n = 2
	}
	runes := []rune(s)
	profile := make(map[string]int)
	if len(runes) < n {
		if len(runes) > 0 {
			profile[string(runes)] = 1
		}
		return profile
	}
	for i := 0; i <= len(runes)-n; i++ {
		gram := string(runes[i : i+n])
		profile[gram]++
	}
	return profile
}

func SetProfile(s string, n int) map[string]bool {
	if n < 1 {
		n = 2
	}
	runes := []rune(s)
	set := make(map[string]bool)
	if len(runes) < n {
		if len(runes) > 0 {
			set[string(runes)] = true
		}
		return set
	}
	for i := 0; i <= len(runes)-n; i++ {
		gram := string(runes[i : i+n])
		set[gram] = true
	}
	return set
}

func Jaccard(a, b string, n int) float64 {
	if a == b {
		return 1.0
	}
	sa := SetProfile(a, n)
	sb := SetProfile(b, n)
	if len(sa) == 0 && len(sb) == 0 {
		return 1.0
	}

	inter := intersection(sa, sb)
	union := unionSize(sa, sb)
	if union == 0 {
		return 0.0
	}
	return float64(inter) / float64(union)
}

func Dice(a, b string, n int) float64 {
	if a == b {
		return 1.0
	}
	sa := SetProfile(a, n)
	sb := SetProfile(b, n)
	total := len(sa) + len(sb)
	if total == 0 {
		return 1.0
	}
	inter := intersection(sa, sb)
	return 2.0 * float64(inter) / float64(total)
}

func Overlap(a, b string, n int) float64 {
	if a == b {
		return 1.0
	}
	sa := SetProfile(a, n)
	sb := SetProfile(b, n)
	minSize := len(sa)
	if len(sb) < minSize {
		minSize = len(sb)
	}
	if minSize == 0 {
		return 0.0
	}
	inter := intersection(sa, sb)
	return float64(inter) / float64(minSize)
}

func WeightedJaccard(a, b string, n int) float64 {
	if a == b {
		return 1.0
	}
	pa := Profile(a, n)
	pb := Profile(b, n)
	if len(pa) == 0 && len(pb) == 0 {
		return 1.0
	}

	allKeys := make(map[string]bool)
	for k := range pa {
		allKeys[k] = true
	}
	for k := range pb {
		allKeys[k] = true
	}

	var interSum, unionSum int
	for k := range allKeys {
		ca, cb := pa[k], pb[k]
		interSum += minInt(ca, cb)
		unionSum += maxInt(ca, cb)
	}
	if unionSum == 0 {
		return 0.0
	}
	return float64(interSum) / float64(unionSum)
}

func PaddedJaccard(a, b string, n int, pad rune) float64 {
	return Jaccard(padString(a, n, pad), padString(b, n, pad), n)
}

func PaddedDice(a, b string, n int, pad rune) float64 {
	return Dice(padString(a, n, pad), padString(b, n, pad), n)
}

func padString(s string, n int, pad rune) string {
	if n < 2 {
		return s
	}
	p := make([]rune, n-1)
	for i := range p {
		p[i] = pad
	}
	prefix := string(p)
	return prefix + s + prefix
}

func intersection(a, b map[string]bool) int {
	count := 0
	for k := range a {
		if b[k] {
			count++
		}
	}
	return count
}

func unionSize(a, b map[string]bool) int {
	merged := make(map[string]bool, len(a)+len(b))
	for k := range a {
		merged[k] = true
	}
	for k := range b {
		merged[k] = true
	}
	return len(merged)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
