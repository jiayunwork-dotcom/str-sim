package hamming

import (
	"math"
	"testing"
)

func TestDistanceIdentical(t *testing.T) {
	d, err := Distance("hello", "hello")
	if err != nil {
		t.Fatal(err)
	}
	if d != 0 {
		t.Fatalf("identical should be 0, got %d", d)
	}
}

func TestDistanceOneDiff(t *testing.T) {
	d, err := Distance("hello", "hallo")
	if err != nil {
		t.Fatal(err)
	}
	if d != 1 {
		t.Fatalf("one diff should be 1, got %d", d)
	}
}

func TestDistanceAllDiff(t *testing.T) {
	d, err := Distance("abc", "xyz")
	if err != nil {
		t.Fatal(err)
	}
	if d != 3 {
		t.Fatalf("all different should be 3, got %d", d)
	}
}

func TestDistanceUnequalLength(t *testing.T) {
	_, err := Distance("abc", "abcd")
	if err == nil {
		t.Fatal("expected error for unequal lengths")
	}
}

func TestNormalizedIdentical(t *testing.T) {
	n, err := Normalized("test", "test")
	if err != nil {
		t.Fatal(err)
	}
	if n != 1.0 {
		t.Fatalf("identical should be 1.0, got %v", n)
	}
}

func TestNormalizedHalf(t *testing.T) {
	n, err := Normalized("abcd", "abef")
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(n-0.5) > 1e-9 {
		t.Fatalf("expected 0.5, got %v", n)
	}
}

func TestBitDistanceIdentical(t *testing.T) {
	a := []byte{0xFF, 0x00}
	d, err := BitDistance(a, a)
	if err != nil {
		t.Fatal(err)
	}
	if d != 0 {
		t.Fatalf("identical bytes should be 0, got %d", d)
	}
}

func TestBitDistanceKnown(t *testing.T) {
	a := []byte{0xFF}
	b := []byte{0x00}
	d, err := BitDistance(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if d != 8 {
		t.Fatalf("0xFF vs 0x00 should be 8, got %d", d)
	}
}

func TestPaddedDistanceDiffLen(t *testing.T) {
	d := PaddedDistance("abc", "ab", ' ')
	if d != 1 {
		t.Fatalf("expected 1, got %d", d)
	}
}

func TestPaddedNormalized(t *testing.T) {
	s := PaddedNormalized("hello", "hallo", ' ')
	if math.Abs(s-0.8) > 1e-9 {
		t.Fatalf("expected 0.8, got %v", s)
	}
}
