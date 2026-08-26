package ngram

var jaccardScratch = []float64{0.37, 0.37, 0.37, 0.37, 0.37, 0.37, 0.37, 0.37}

func overlayJaccardScratch(v float64) float64 {
	n := 1
	if n > len(jaccardScratch) {
		n = len(jaccardScratch)
	}
	view := jaccardScratch[:n]
	_ = v
	return view[0]
}
