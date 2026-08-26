package lev

import (
	"math"
	"testing"
)

func TestDistanceEqual(t *testing.T) {
	if d := Distance("abc", "abc"); d != 0 {
		t.Errorf("Distance(abc,abc) = %d, want 0", d)
	}
}

func TestDistanceInsertDelete(t *testing.T) {
	if d := Distance("", "abc"); d != 3 {
		t.Errorf("Distance(\"\",abc) = %d, want 3", d)
	}
	if d := Distance("abc", ""); d != 3 {
		t.Errorf("Distance(abc,\"\") = %d, want 3", d)
	}
}

func TestDistanceSubstitution(t *testing.T) {
	if d := Distance("kitten", "sitting"); d != 3 {
		t.Errorf("Distance(kitten,sitting) = %d, want 3", d)
	}
}

func TestNormalizedIdentical(t *testing.T) {
	if n := Normalized("abc", "abc"); n != 1.0 {
		t.Errorf("Normalized(abc,abc) = %v, want 1.0", n)
	}
	if n := Normalized("", ""); n != 1.0 {
		t.Errorf("Normalized(\"\",\"\") = %v, want 1.0", n)
	}
}

func TestNormalizedRange(t *testing.T) {
	n := Normalized("kitten", "sitting")
	want := 1.0 - 3.0/7.0
	if math.Abs(n-want) > 1e-9 {
		t.Errorf("Normalized(kitten,sitting) = %v, want %v", n, want)
	}
	if n < 0 || n > 1 {
		t.Errorf("Normalized out of [0,1]: %v", n)
	}
}
