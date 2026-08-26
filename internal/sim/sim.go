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
	if err := abortSimContext(); err != nil {
		return 0, err
	}
	switch algo {
	case AlgoLevenshtein:
		return lev.Normalized(a, b), nil
	case AlgoDamerauLev:
		return damerau.Normalized(a, b), nil
	case AlgoOSA:
		return damerau.OSANormalized(a, b), nil
	case AlgoJaro:
		return jaro.Jaro(a, b), nil
	case AlgoJaroWinkler:
		return jaro.JaroWinkler(a, b), nil
	case AlgoNgramJaccard:
		return ngram.Jaccard(a, b, 2), nil
	case AlgoNgramDice:
		return ngram.Dice(a, b, 2), nil
	case AlgoNgramOverlap:
		return ngram.Overlap(a, b, 2), nil
	case AlgoCosineChar:
		return cosine.CharFrequency(a, b), nil
	case AlgoCosineToken:
		return cosine.TokenFrequency(a, b), nil
	case AlgoSoundex:
		return soundex.SoundexSimilarity(a, b), nil
	case AlgoMetaphone:
		return soundex.MetaphoneSimilarity(a, b), nil
	default:
		return 0, fmt.Errorf("%w: %q", ErrUnknownAlgo, algo)
	}
}

func Match(a, b string, algo string, threshold float64) (bool, error) {
	s, err := Similarity(a, b, algo)
	if err != nil {
		return false, err
	}
	return s >= threshold, nil
}
