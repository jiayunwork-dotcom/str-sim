package damerau

import (
	"math"
	"testing"
)

func TestOSADistanceIdentical(t *testing.T) {
	if OSADistance("hello", "hello") != 0 {
		t.Fatal("identical strings should have distance 0")
	}
}

func TestOSADistanceTransposition(t *testing.T) {
	if d := OSADistance("ab", "ba"); d != 1 {
		t.Fatalf("OSA('ab','ba') = %d, expected 1", d)
	}
}

func TestOSADistanceVsLevenshtein(t *testing.T) {
	d := OSADistance("ca", "abc")
	if d > 3 {
		t.Fatalf("OSA('ca','abc') = %d, should be <= 3", d)
	}
}

func TestOSARestriction(t *testing.T) {
	osa := OSADistance("CA", "ABC")
	dl := Distance("CA", "ABC")
	if dl > osa {
		t.Fatalf("DL=%d should be <= OSA=%d", dl, osa)
	}
}

func TestDistanceIdentical(t *testing.T) {
	if Distance("test", "test") != 0 {
		t.Fatal("identical strings should have distance 0")
	}
}

func TestDistanceTransposition(t *testing.T) {
	if d := Distance("ab", "ba"); d != 1 {
		t.Fatalf("DL('ab','ba') = %d, expected 1", d)
	}
}

func TestDistanceEmpty(t *testing.T) {
	if d := Distance("", "abc"); d != 3 {
		t.Fatalf("DL('','abc') = %d, expected 3", d)
	}
	if d := Distance("abc", ""); d != 3 {
		t.Fatalf("DL('abc','') = %d, expected 3", d)
	}
}

func TestDistanceSymmetry(t *testing.T) {
	d1 := Distance("kitten", "sitting")
	d2 := Distance("sitting", "kitten")
	if d1 != d2 {
		t.Fatalf("DL not symmetric: %d != %d", d1, d2)
	}
}

func TestDistanceKnownValues(t *testing.T) {
	d := Distance("kitten", "sitting")
	if d != 3 {
		t.Fatalf("DL('kitten','sitting') = %d, expected 3", d)
	}
}

func TestNormalizedRange(t *testing.T) {
	n := Normalized("hello", "world")
	if n < 0 || n > 1 {
		t.Fatalf("Normalized=%v out of [0,1]", n)
	}
}

func TestNormalizedIdentical(t *testing.T) {
	if math.Abs(Normalized("abc", "abc")-1.0) > 1e-9 {
		t.Fatal("identical strings should have normalized=1.0")
	}
}

func TestOSANormalizedSymmetry(t *testing.T) {
	n1 := OSANormalized("abc", "bac")
	n2 := OSANormalized("bac", "abc")
	if math.Abs(n1-n2) > 1e-9 {
		t.Fatalf("OSANormalized not symmetric: %v != %v", n1, n2)
	}
}

func TestDistanceUnicode(t *testing.T) {
	d := Distance("café", "cafe")
	if d != 1 {
		t.Fatalf("DL('café','cafe') = %d, expected 1", d)
	}
}
