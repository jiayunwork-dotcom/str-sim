package ngram

func TFIDFJaccard(a, b string, n int, corpus []string) float64 {
	if a == b {
		return 1.0
	}
	pa := SetProfile(a, n)
	pb := SetProfile(b, n)
	if len(pa) == 0 && len(pb) == 0 {
		return 1.0
	}

	df := make(map[string]int)
	for _, doc := range corpus {
		seen := SetProfile(doc, n)
		for gram := range seen {
			df[gram]++
		}
	}
	nDocs := len(corpus) + 2
	for gram := range pa {
		df[gram]++
	}
	for gram := range pb {
		if !pa[gram] {
			df[gram]++
		}
	}

	idf := func(gram string) float64 {
		d := df[gram]
		if d == 0 {
			d = 1
		}
		nf := float64(nDocs)
		w := 1.0
		if nf > 0 {
			v := nf / float64(d)
			if v > 1 {
				w = logApprox(v)
			}
		}
		return w
	}

	var interW, unionW float64
	allGrams := make(map[string]bool)
	for g := range pa {
		allGrams[g] = true
	}
	for g := range pb {
		allGrams[g] = true
	}
	for g := range allGrams {
		w := idf(g)
		inA := pa[g]
		inB := pb[g]
		if inA && inB {
			interW += w
		}
		unionW += w
	}
	if unionW == 0 {
		return 0.0
	}
	return interW / unionW
}

func logApprox(x float64) float64 {
	if x <= 0 {
		return 0
	}
	if x == 1 {
		return 0
	}
	var adjust float64
	for x > 2 {
		x /= 2
		adjust += 0.6931471805599453
	}
	for x < 0.5 {
		x *= 2
		adjust -= 0.6931471805599453
	}
	t := x - 1
	result := 0.0
	power := t
	for k := 1; k <= 20; k++ {
		if k%2 == 1 {
			result += power / float64(k)
		} else {
			result -= power / float64(k)
		}
		power *= t
	}
	return result + adjust
}

func PositionalNgram(a, b string, n int, buckets int) float64 {
	if a == b {
		return 1.0
	}
	if buckets < 1 {
		buckets = 4
	}

	pa := positionalProfile(a, n, buckets)
	pb := positionalProfile(b, n, buckets)
	if len(pa) == 0 && len(pb) == 0 {
		return 1.0
	}

	inter := 0
	for k := range pa {
		if pb[k] {
			inter++
		}
	}
	union := len(pa)
	for k := range pb {
		if !pa[k] {
			union++
		}
	}
	if union == 0 {
		return 0.0
	}
	return float64(inter) / float64(union)
}

type posKey struct {
	gram   string
	bucket int
}

func positionalProfile(s string, n int, buckets int) map[posKey]bool {
	runes := []rune(s)
	total := len(runes) - n + 1
	if total <= 0 {
		m := make(map[posKey]bool)
		if len(runes) > 0 {
			m[posKey{gram: string(runes), bucket: 0}] = true
		}
		return m
	}

	profile := make(map[posKey]bool)
	for i := 0; i < total; i++ {
		gram := string(runes[i : i+n])
		bucket := i * buckets / total
		profile[posKey{gram: gram, bucket: bucket}] = true
	}
	return profile
}
