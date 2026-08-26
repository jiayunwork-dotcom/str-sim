package matcher

import (
	"str-sim/internal/cosine"
	"str-sim/internal/damerau"
	"str-sim/internal/hamming"
	"str-sim/internal/jaro"
	"str-sim/internal/lcsubstr"
	"str-sim/internal/lev"
	"str-sim/internal/ngram"
	"str-sim/internal/qgram"
	"str-sim/internal/soundex"
)

func BuiltinAlgos() map[string]AlgoFunc {
	return map[string]AlgoFunc{
		"levenshtein":    lev.Normalized,
		"damerau-lev":    damerau.Normalized,
		"osa":            damerau.OSANormalized,
		"jaro":           jaro.Jaro,
		"jaro-winkler":   jaro.JaroWinkler,
		"bigram-jaccard": func(a, b string) float64 { return ngram.Jaccard(a, b, 2) },
		"bigram-dice":    func(a, b string) float64 { return ngram.Dice(a, b, 2) },
		"bigram-overlap": func(a, b string) float64 { return ngram.Overlap(a, b, 2) },
		"cosine-char":    cosine.CharFrequency,
		"cosine-token":   cosine.TokenFrequency,
		"soundex":        soundex.SoundexSimilarity,
		"metaphone":      soundex.MetaphoneSimilarity,
		"qgram":          func(a, b string) float64 { return qgram.Normalized(a, b, 2) },
		"lcs-substr":     lcsubstr.SubstringSimilarity,
		"lcs-subseq":     lcsubstr.SubsequenceSimilarity,
		"hamming-padded": func(a, b string) float64 { return hamming.PaddedNormalized(a, b, ' ') },
	}
}

func ConfigFromNames(names []string) Config {
	algos := BuiltinAlgos()
	var weighted []WeightedAlgo
	for _, name := range names {
		fn, ok := algos[name]
		if !ok {
			continue
		}
		weighted = append(weighted, WeightedAlgo{Name: name, Func: fn, Weight: 1.0})
	}
	return Config{Algorithms: weighted}
}

func EditDistanceConfig() Config {
	return Config{
		Algorithms: []WeightedAlgo{
			{Name: "levenshtein", Func: lev.Normalized, Weight: 0.35},
			{Name: "damerau-lev", Func: damerau.Normalized, Weight: 0.35},
			{Name: "jaro-winkler", Func: jaro.JaroWinkler, Weight: 0.30},
		},
	}
}

func TokenBasedConfig() Config {
	return Config{
		Algorithms: []WeightedAlgo{
			{Name: "bigram-dice", Func: func(a, b string) float64 { return ngram.Dice(a, b, 2) }, Weight: 0.30},
			{Name: "cosine-token", Func: cosine.TokenFrequency, Weight: 0.35},
			{Name: "lcs-subseq", Func: lcsubstr.SubsequenceSimilarity, Weight: 0.35},
		},
	}
}
