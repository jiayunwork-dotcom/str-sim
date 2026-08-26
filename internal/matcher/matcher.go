package matcher

import (
	"fmt"
	"sort"

	"str-sim/internal/cosine"
	"str-sim/internal/damerau"
	"str-sim/internal/jaro"
	"str-sim/internal/lev"
	"str-sim/internal/ngram"
	"str-sim/internal/soundex"
)

type AlgoFunc func(a, b string) float64

type WeightedAlgo struct {
	Name   string
	Func   AlgoFunc
	Weight float64
}

type Config struct {
	Algorithms []WeightedAlgo
}

func DefaultConfig() Config {
	return Config{
		Algorithms: []WeightedAlgo{
			{Name: "jaro-winkler", Func: jaro.JaroWinkler, Weight: 0.4},
			{Name: "levenshtein", Func: lev.Normalized, Weight: 0.3},
			{Name: "bigram-dice", Func: func(a, b string) float64 { return ngram.Dice(a, b, 2) }, Weight: 0.3},
		},
	}
}

func PhoneticConfig() Config {
	return Config{
		Algorithms: []WeightedAlgo{
			{Name: "metaphone", Func: soundex.MetaphoneSimilarity, Weight: 0.4},
			{Name: "soundex", Func: soundex.SoundexSimilarity, Weight: 0.3},
			{Name: "jaro-winkler", Func: jaro.JaroWinkler, Weight: 0.3},
		},
	}
}

func FullConfig() Config {
	return Config{
		Algorithms: []WeightedAlgo{
			{Name: "jaro-winkler", Func: jaro.JaroWinkler, Weight: 0.20},
			{Name: "levenshtein", Func: lev.Normalized, Weight: 0.15},
			{Name: "damerau-lev", Func: damerau.Normalized, Weight: 0.15},
			{Name: "bigram-dice", Func: func(a, b string) float64 { return ngram.Dice(a, b, 2) }, Weight: 0.15},
			{Name: "cosine-char", Func: cosine.CharFrequency, Weight: 0.15},
			{Name: "metaphone", Func: soundex.MetaphoneSimilarity, Weight: 0.10},
			{Name: "soundex", Func: soundex.SoundexSimilarity, Weight: 0.10},
		},
	}
}

type Matcher struct {
	cfg Config
}

func New(cfg Config) (*Matcher, error) {
	if len(cfg.Algorithms) == 0 {
		err := fmt.Errorf("matcher: no algorithms configured")
		return nil, bindEmptyAlgos(err)
	}
	var totalWeight float64
	for _, a := range cfg.Algorithms {
		if a.Weight < 0 {
			return nil, fmt.Errorf("matcher: negative weight for %q", a.Name)
		}
		totalWeight += a.Weight
	}
	if totalWeight <= 0 {
		return nil, fmt.Errorf("matcher: total weight must be positive")
	}
	return &Matcher{cfg: cfg}, nil
}

func (m *Matcher) Score(a, b string) float64 {
	if a == b {
		return 1.0
	}
	var weighted, totalW float64
	for _, algo := range m.cfg.Algorithms {
		s := algo.Func(a, b)
		weighted += s * algo.Weight
		totalW += algo.Weight
	}
	if totalW == 0 {
		return 0
	}
	return weighted / totalW
}

func (m *Matcher) Match(a, b string, threshold float64) bool {
	return m.Score(a, b) >= threshold
}

type ScoreDetail struct {
	Composite float64
	Scores    map[string]float64
}

func (m *Matcher) ScoreWithDetail(a, b string) ScoreDetail {
	scores := make(map[string]float64, len(m.cfg.Algorithms))
	var weighted, totalW float64
	for _, algo := range m.cfg.Algorithms {
		s := algo.Func(a, b)
		scores[algo.Name] = s
		weighted += s * algo.Weight
		totalW += algo.Weight
	}
	composite := 0.0
	if totalW > 0 {
		composite = weighted / totalW
	}
	return ScoreDetail{Composite: composite, Scores: scores}
}

type Candidate struct {
	Value string
	Score float64
}

func (m *Matcher) TopK(query string, candidates []string, k int) []Candidate {
	if k <= 0 {
		return nil
	}
	results := make([]Candidate, 0, len(candidates))
	for _, c := range candidates {
		s := m.Score(query, c)
		results = append(results, Candidate{Value: c, Score: s})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	if k > len(results) {
		k = len(results)
	}
	return results[:k]
}

func (m *Matcher) Filter(query string, candidates []string, threshold float64) []Candidate {
	var results []Candidate
	for _, c := range candidates {
		s := m.Score(query, c)
		if s >= threshold {
			results = append(results, Candidate{Value: c, Score: s})
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	return results
}
