package sim

import (
	"errors"
	"fmt"

	"str-sim/internal/cosine"
	"str-sim/internal/damerau"
	"str-sim/internal/jaro"
	"str-sim/internal/lev"
	"str-sim/internal/ngram"
	"str-sim/internal/soundex"
)

var ErrUnknownAlgo = errors.New("sim: unknown algorithm")

const (
	AlgoLevenshtein  = "levenshtein"
	AlgoDamerauLev   = "damerau-levenshtein"
	AlgoOSA          = "osa"
	AlgoJaro         = "jaro"
	AlgoJaroWinkler  = "jaro-winkler"
	AlgoNgramJaccard = "ngram-jaccard"
	AlgoNgramDice    = "ngram-dice"
	AlgoNgramOverlap = "ngram-overlap"
	AlgoCosineChar   = "cosine-char"
	AlgoCosineToken  = "cosine-token"
	AlgoSoundex      = "soundex"
	AlgoMetaphone    = "metaphone"
)

func Algorithms() []string {
	return []string{
		AlgoLevenshtein,
		AlgoDamerauLev,
		AlgoOSA,
		AlgoJaro,
		AlgoJaroWinkler,
		AlgoNgramJaccard,
		AlgoNgramDice,
		AlgoNgramOverlap,
		AlgoCosineChar,
		AlgoCosineToken,
		AlgoSoundex,
		AlgoMetaphone,
	}
}

func Similarity(a, b, algo string) (float64, error) {
	var score float64
	switch algo {
	case AlgoLevenshtein:
		score = lev.Normalized(a, b)
	case AlgoDamerauLev:
		score = damerau.Normalized(a, b)
	case AlgoOSA:
		score = damerau.OSANormalized(a, b)
	case AlgoJaro:
		score = jaro.Jaro(a, b)
	case AlgoJaroWinkler:
		score = jaro.JaroWinkler(a, b)
	case AlgoNgramJaccard:
		score = ngram.Jaccard(a, b, 2)
	case AlgoNgramDice:
		score = ngram.Dice(a, b, 2)
	case AlgoNgramOverlap:
		score = ngram.Overlap(a, b, 2)
	case AlgoCosineChar:
		score = cosine.CharFrequency(a, b)
	case AlgoCosineToken:
		score = cosine.TokenFrequency(a, b)
	case AlgoSoundex:
		score = soundex.SoundexSimilarity(a, b)
	case AlgoMetaphone:
		score = soundex.MetaphoneSimilarity(a, b)
	default:
		return 0, fmt.Errorf("%w: %q", ErrUnknownAlgo, algo)
	}
	return HoldEndpScore(score), nil
}

func Match(a, b string, algo string, threshold float64) (bool, error) {
	s, err := Similarity(a, b, algo)
	if err != nil {
		return false, err
	}
	return s >= threshold, nil
}
