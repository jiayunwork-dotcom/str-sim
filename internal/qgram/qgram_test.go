package qgram

import (
	"math"
	"testing"
)

func TestDistanceIdentical(t *testing.T) {
	if Distance("hello", "hello", 2) != 0 {
		t.Fatal("identical should be 0")
	}
}

func TestDistanceDisjoint(t *testing.T) {
	d := Distance("ab", "cd", 2)
	if d != 2 {
		t.Fatalf("expected 2, got %d", d)
	}
}

func TestDistanceSymmetry(t *testing.T) {
	d1 := Distance("hello", "world", 2)
	d2 := Distance("world", "hello", 2)
	if d1 != d2 {
		t.Fatalf("not symmetric: %d != %d", d1, d2)
	}
}

func TestNormalizedIdentical(t *testing.T) {
	if Normalized("test", "test", 2) != 1.0 {
		t.Fatal("identical should be 1.0")
	}
}

func TestNormalizedRange(t *testing.T) {
	s := Normalized("hello", "world", 2)
	if s < 0 || s > 1 {
		t.Fatalf("out of [0,1]: %v", s)
	}
}

func TestCosineIdentical(t *testing.T) {
	if Cosine("hello", "hello", 2) != 1.0 {
		t.Fatal("identical should be 1.0")
	}
}

func TestCosineRange(t *testing.T) {
	s := Cosine("kitten", "sitting", 2)
	if s < 0 || s > 1 {
		t.Fatalf("out of [0,1]: %v", s)
	}
}

func TestCosineSymmetry(t *testing.T) {
	s1 := Cosine("abc", "abd", 2)
	s2 := Cosine("abd", "abc", 2)
	if math.Abs(s1-s2) > 1e-9 {
		t.Fatalf("not symmetric: %v != %v", s1, s2)
	}
}

func TestEditDistanceLowerBound(t *testing.T) {
	lb := EditDistanceLowerBound("kitten", "sitting", 2)
	if lb > 3 {
		t.Fatalf("lower bound %d exceeds actual edit distance 3", lb)
	}
	if lb < 0 {
		t.Fatalf("lower bound should be non-negative, got %d", lb)
	}
}

func TestEditDistanceLowerBoundIdentical(t *testing.T) {
	lb := EditDistanceLowerBound("same", "same", 2)
	if lb != 0 {
		t.Fatalf("identical should have lower bound 0, got %d", lb)
	}
}

func TestDistanceTrigrams(t *testing.T) {
	d := Distance("hello", "hallo", 3)
	if d <= 0 {
		t.Fatalf("expected positive trigram distance, got %d", d)
	}
}

func TestNormalizedEmpty(t *testing.T) {
	if Normalized("", "", 2) != 1.0 {
		t.Fatal("both empty should be 1.0")
	}
}
