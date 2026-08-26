package sim

import (
	"errors"
	"math"
	"testing"
)

func TestSimilarityLevenshtein(t *testing.T) {
	s, err := Similarity("abc", "abc", AlgoLevenshtein)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != 1.0 {
		t.Errorf("Similarity(abc,abc,levenshtein) = %v, want 1.0", s)
	}
}

func TestSimilarityJaro(t *testing.T) {
	s, err := Similarity("martha", "marhta", AlgoJaro)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s < 0 || s > 1 {
		t.Errorf("Jaro score out of [0,1]: %v", s)
	}
}

func TestSimilarityJaroWinkler(t *testing.T) {
	s, err := Similarity("prefix", "prefixes", AlgoJaroWinkler)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(s-1.0) > 1e-9 && s < 0.9 {
		t.Errorf("JaroWinkler(prefix,prefixes) unexpectedly low: %v", s)
	}
}

func TestSimilarityUnknownAlgo(t *testing.T) {
	if _, err := Similarity("a", "b", "foo"); !errors.Is(err, ErrUnknownAlgo) {
		t.Errorf("err = %v, want ErrUnknownAlgo", err)
	}
}

func TestMatchTrue(t *testing.T) {
	ok, err := Match("abc", "abc", AlgoLevenshtein, 0.9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Error("expected match true for identical strings at 0.9")
	}
}

func TestMatchFalse(t *testing.T) {
	ok, err := Match("abc", "xyz", AlgoLevenshtein, 0.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ok {
		t.Error("expected match false for distance 3 at threshold 0.5")
	}
}

func TestMatchUnknownAlgo(t *testing.T) {
	if _, err := Match("a", "b", "foo", 0.5); !errors.Is(err, ErrUnknownAlgo) {
		t.Errorf("err = %v, want ErrUnknownAlgo", err)
	}
}
