package jaro

import (
	"math"
	"testing"
)

func TestJaroIdentical(t *testing.T) {
	if v := Jaro("abc", "abc"); v != 1.0 {
		t.Errorf("Jaro(abc,abc) = %v, want 1.0", v)
	}
}

func TestJaroEmpty(t *testing.T) {
	if v := Jaro("", ""); v != 1.0 {
		t.Errorf("Jaro(\"\",\"\") = %v, want 1.0", v)
	}
	if v := Jaro("a", ""); v != 0.0 {
		t.Errorf("Jaro(a,\"\") = %v, want 0.0", v)
	}
}

func TestJaroSymmetric(t *testing.T) {
	a, b := "martha", "marhta"
	if x, y := Jaro(a, b), Jaro(b, a); x != y {
		t.Errorf("Jaro not symmetric: %v vs %v", x, y)
	}
}

func TestJaroWinklerPrefixBoost(t *testing.T) {
	a, b := "prefix", "prefixes"
	j := Jaro(a, b)
	w := JaroWinkler(a, b)
	if w <= j {
		t.Errorf("JaroWinkler (%v) should exceed Jaro (%v) for shared prefix", w, j)
	}
	if w < 0 || w > 1 {
		t.Errorf("JaroWinkler out of [0,1]: %v", w)
	}
}

func TestJaroWinklerIdentical(t *testing.T) {
	if v := JaroWinkler("abc", "abc"); math.Abs(v-1.0) > 1e-9 {
		t.Errorf("JaroWinkler(abc,abc) = %v, want 1.0", v)
	}
}
