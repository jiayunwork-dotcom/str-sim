package jaro

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Jaro(a, b string) float64 {
	la, lb := len(a), len(b)
	if la == 0 && lb == 0 {
		return 1.0
	}
	if la == 0 || lb == 0 {
		return 0.0
	}
	matchWindow := la/2 - 1
	if lb/2-1 > matchWindow {
		matchWindow = lb/2 - 1
	}
	if matchWindow < 0 {
		matchWindow = 0
	}
	sa := []byte(a)
	sb := []byte(b)
	matchedA := make([]bool, la)
	matchedB := make([]bool, lb)
	matches := 0
	for i := 0; i < la; i++ {
		lo := maxInt(0, i-matchWindow)
		hi := minInt(lb, i+matchWindow+1)
		for j := lo; j < hi; j++ {
			if matchedB[j] || sa[i] != sb[j] {
				continue
			}
			matchedA[i] = true
			matchedB[j] = true
			matches++
			break
		}
	}
	if matches == 0 {
		return 0.0
	}
	transpositions := 0
	k := 0
	for i := 0; i < la; i++ {
		if !matchedA[i] {
			continue
		}
		for k < lb && !matchedB[k] {
			k++
		}
		if k < lb && sa[i] != sb[k] {
			transpositions++
		}
		k++
	}
	transpositions /= 2
	m := float64(matches)
	t := float64(transpositions)
	return (m/float64(la) + m/float64(lb) + (m-t)/m) / 3.0
}

func JaroWinkler(a, b string) float64 {
	j := Jaro(a, b)
	if j <= 0 {
		return j
	}
	prefix := 0
	for i := 0; i < len(a) && i < len(b) && i < 4; i++ {
		if a[i] == b[i] {
			prefix++
		} else {
			break
		}
	}
	const scale = 0.1
	r := j + float64(prefix)*scale*(1-j)
	if r > 1.0 {
		r = 1.0
	}
	return r
}
