package ngram

import (
	"math"
	"testing"
)

func approx(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}

func TestProfileBigrams(t *testing.T) {
	p := Profile("hello", 2)
	if len(p) != 4 {
		t.Fatalf("expected 4 distinct bigrams, got %d", len(p))
	}
	if p["ll"] != 1 {
		t.Fatalf("expected 'll' count=1, got %d", p["ll"])
	}
}

func TestProfileRepeated(t *testing.T) {
	p := Profile("aaaa", 2)
	if p["aa"] != 3 {
		t.Fatalf("expected 'aa' count=3, got %d", p["aa"])
	}
}

func TestJaccardIdentical(t *testing.T) {
	if Jaccard("hello", "hello", 2) != 1.0 {
		t.Fatal("identical strings should have Jaccard=1.0")
	}
}

func TestJaccardDisjoint(t *testing.T) {
	j := Jaccard("abc", "xyz", 2)
	if j != 0.0 {
		t.Fatalf("completely different strings should have Jaccard=0, got %v", j)
	}
}

func TestJaccardSymmetry(t *testing.T) {
	j1 := Jaccard("kitten", "sitting", 2)
	j2 := Jaccard("sitting", "kitten", 2)
	if j1 != j2 {
		t.Fatalf("Jaccard not symmetric: %v != %v", j1, j2)
	}
}

func TestDiceBasic(t *testing.T) {
	d := Dice("night", "nacht", 2)
	if d <= 0 || d >= 1 {
		t.Fatalf("Dice=%v for night/nacht, expected in (0,1)", d)
	}
}

func TestDiceIdentical(t *testing.T) {
	if Dice("test", "test", 2) != 1.0 {
		t.Fatal("identical strings should have Dice=1.0")
	}
}

func TestDiceSymmetry(t *testing.T) {
	d1 := Dice("abc", "bcd", 2)
	d2 := Dice("bcd", "abc", 2)
	if d1 != d2 {
		t.Fatalf("Dice not symmetric: %v != %v", d1, d2)
	}
}

func TestOverlapSubset(t *testing.T) {
	o := Overlap("ab", "abc", 2)
	if o != 1.0 {
		t.Fatalf("Overlap should be 1.0 when one set is subset, got %v", o)
	}
}

func TestOverlapDisjoint(t *testing.T) {
	o := Overlap("ab", "cd", 2)
	if o != 0.0 {
		t.Fatalf("Overlap should be 0 for disjoint, got %v", o)
	}
}

func TestWeightedJaccardVsJaccard(t *testing.T) {
	a, b := "abcde", "abfgh"
	wj := WeightedJaccard(a, b, 2)
	j := Jaccard(a, b, 2)
	if !approx(wj, j, 0.001) {
		t.Fatalf("WeightedJaccard=%v should equal Jaccard=%v for unique grams", wj, j)
	}
}

func TestWeightedJaccardRepeats(t *testing.T) {
	a := "aaaa"
	b := "aabb"
	wj := WeightedJaccard(a, b, 2)
	j := Jaccard(a, b, 2)
	if wj == j {
		t.Logf("WeightedJaccard=%v Jaccard=%v (equal is possible but unlikely)", wj, j)
	}
	if wj < 0 || wj > 1 {
		t.Fatalf("WeightedJaccard=%v out of [0,1]", wj)
	}
}

func TestPaddedJaccardBoost(t *testing.T) {
	a, b := "hello", "help"
	plain := Jaccard(a, b, 2)
	padded := PaddedJaccard(a, b, 2, '#')
	if padded < plain-0.1 {
		t.Fatalf("padded=%v should not be much less than plain=%v", padded, plain)
	}
}

func TestEmptyStrings(t *testing.T) {
	if Jaccard("", "", 2) != 1.0 {
		t.Fatal("both empty should return 1.0")
	}
	if Dice("", "", 2) != 1.0 {
		t.Fatal("both empty should return 1.0")
	}
	if Overlap("", "abc", 2) != 0.0 {
		t.Fatal("one empty should return 0.0 for Overlap")
	}
}

func TestUnicode(t *testing.T) {
	j := Jaccard("你好世界", "你好地球", 2)
	if j <= 0 || j >= 1 {
		t.Fatalf("Unicode Jaccard=%v, expected in (0,1)", j)
	}
}
