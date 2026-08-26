package soundex

import (
	"math"
	"testing"
)

func TestSoundexRobert(t *testing.T) {
	if code := Soundex("Robert"); code != "R163" {
		t.Fatalf("Soundex('Robert') = %q, expected R163", code)
	}
}

func TestSoundexRupert(t *testing.T) {
	if code := Soundex("Rupert"); code != "R163" {
		t.Fatalf("Soundex('Rupert') = %q, expected R163", code)
	}
}

func TestSoundexAshcraft(t *testing.T) {
	code := Soundex("Ashcraft")
	if code != "A261" {
		t.Fatalf("Soundex('Ashcraft') = %q, expected A261", code)
	}
}

func TestSoundexEmpty(t *testing.T) {
	if code := Soundex(""); code != "0000" {
		t.Fatalf("Soundex('') = %q, expected 0000", code)
	}
}

func TestSoundexSimilarNames(t *testing.T) {
	if Soundex("Smith") != Soundex("Smyth") {
		t.Fatalf("Smith=%q Smyth=%q should match", Soundex("Smith"), Soundex("Smyth"))
	}
}

func TestSoundexSimilarityIdentical(t *testing.T) {
	if SoundexSimilarity("John", "John") != 1.0 {
		t.Fatal("identical should be 1.0")
	}
}

func TestSoundexSimilarityDifferent(t *testing.T) {
	s := SoundexSimilarity("John", "Mary")
	if s >= 1.0 {
		t.Fatalf("different names should not be 1.0, got %v", s)
	}
}

func TestMetaphoneBasic(t *testing.T) {
	m := Metaphone("Smith")
	if m != "SM0" {
		t.Fatalf("Metaphone('Smith') = %q, expected SM0", m)
	}
}

func TestMetaphoneKnowledge(t *testing.T) {
	m := Metaphone("Knowledge")
	if len(m) == 0 {
		t.Fatal("Metaphone('Knowledge') should not be empty")
	}
	if m[0] != 'N' {
		t.Fatalf("Metaphone('Knowledge') = %q, expected to start with N", m)
	}
}

func TestMetaphonePhone(t *testing.T) {
	m := Metaphone("Phone")
	if len(m) == 0 || m[0] != 'F' {
		t.Fatalf("Metaphone('Phone') = %q, expected to start with F", m)
	}
}

func TestMetaphoneSimilarityIdentical(t *testing.T) {
	if MetaphoneSimilarity("test", "test") != 1.0 {
		t.Fatal("identical should return 1.0")
	}
}

func TestMetaphoneSimilarityPhonetic(t *testing.T) {
	s := MetaphoneSimilarity("Stephen", "Steven")
	if s < 0.5 {
		t.Fatalf("Stephen/Steven phonetic similarity = %v, expected >= 0.5", s)
	}
}

func TestMetaphoneSimilaritySymmetry(t *testing.T) {
	s1 := MetaphoneSimilarity("cat", "kat")
	s2 := MetaphoneSimilarity("kat", "cat")
	if math.Abs(s1-s2) > 1e-9 {
		t.Fatalf("not symmetric: %v != %v", s1, s2)
	}
}
