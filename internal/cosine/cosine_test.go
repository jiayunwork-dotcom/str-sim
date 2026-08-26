package cosine

import (
	"math"
	"testing"
)

func approx(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}

func TestCharFrequencyIdentical(t *testing.T) {
	if CharFrequency("hello", "hello") != 1.0 {
		t.Fatal("identical strings should return 1.0")
	}
}

func TestCharFrequencyDisjoint(t *testing.T) {
	s := CharFrequency("aaa", "zzz")
	if s != 0.0 {
		t.Fatalf("disjoint chars should return 0.0, got %v", s)
	}
}

func TestCharFrequencySymmetry(t *testing.T) {
	s1 := CharFrequency("hello", "world")
	s2 := CharFrequency("world", "hello")
	if !approx(s1, s2, 1e-9) {
		t.Fatalf("not symmetric: %v != %v", s1, s2)
	}
}

func TestCharFrequencyPartial(t *testing.T) {
	s := CharFrequency("abc", "abd")
	if s <= 0 || s >= 1 {
		t.Fatalf("expected partial similarity in (0,1), got %v", s)
	}
}

func TestTokenFrequencyIdentical(t *testing.T) {
	if TokenFrequency("hello world", "hello world") != 1.0 {
		t.Fatal("identical should return 1.0")
	}
}

func TestTokenFrequencyPartial(t *testing.T) {
	s := TokenFrequency("the quick brown fox", "the slow brown dog")
	if s <= 0 || s >= 1 {
		t.Fatalf("expected partial similarity, got %v", s)
	}
}

func TestTokenFrequencyCaseInsensitive(t *testing.T) {
	s1 := TokenFrequency("Hello World", "hello world")
	if !approx(s1, 1.0, 1e-9) {
		t.Fatalf("case should not matter, got %v", s1)
	}
}

func TestTFIDFBasic(t *testing.T) {
	corpus := []string{
		"the cat sat on the mat",
		"the dog sat on the log",
		"cats and dogs are friends",
	}
	s := TFIDF("the cat", "the dog", corpus)
	if s <= 0 || s >= 1 {
		t.Fatalf("expected partial TFIDF similarity, got %v", s)
	}
}

func TestTFIDFIdentical(t *testing.T) {
	s := TFIDF("hello world", "hello world", nil)
	if s != 1.0 {
		t.Fatalf("identical should return 1.0, got %v", s)
	}
}

func TestTFIDFRareTermBoost(t *testing.T) {
	corpus := []string{
		"the cat sat",
		"the dog sat",
		"the bird flew",
		"the fish swam",
	}
	sCat := TFIDF("the cat", "a cat", corpus)
	sThe := TFIDF("the cat", "the dog", corpus)
	if sCat <= sThe {
		t.Fatalf("sharing rare term 'cat' (%v) should score higher than common 'the' (%v)", sCat, sThe)
	}
}

func TestTokenizeHandlesPunctuation(t *testing.T) {
	tokens := tokenize("hello, world! foo-bar")
	expected := []string{"hello", "world", "foo", "bar"}
	if len(tokens) != len(expected) {
		t.Fatalf("tokenize got %v, expected %v", tokens, expected)
	}
	for i, tok := range tokens {
		if tok != expected[i] {
			t.Fatalf("token[%d]=%q, expected %q", i, tok, expected[i])
		}
	}
}

func TestEmptyStrings(t *testing.T) {
	if CharFrequency("", "") != 1.0 {
		t.Fatal("both empty should return 1.0")
	}
	if CharFrequency("", "abc") != 0.0 {
		t.Fatal("one empty should return 0.0")
	}
}
