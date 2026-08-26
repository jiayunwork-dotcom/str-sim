package lcsubstr

import (
	"math"
	"testing"
)

func TestLCSSubstringIdentical(t *testing.T) {
	if LongestCommonSubstring("hello", "hello") != 5 {
		t.Fatal("identical should return full length")
	}
}

func TestLCSSubstringPartial(t *testing.T) {
	l := LongestCommonSubstring("abcdef", "xbcdy")
	if l != 3 {
		t.Fatalf("expected 3, got %d", l)
	}
}

func TestLCSSubstringDisjoint(t *testing.T) {
	if LongestCommonSubstring("abc", "xyz") != 0 {
		t.Fatal("disjoint should return 0")
	}
}

func TestLCSSubstringEmpty(t *testing.T) {
	if LongestCommonSubstring("", "hello") != 0 {
		t.Fatal("empty should return 0")
	}
}

func TestSubstringSimilarityIdentical(t *testing.T) {
	if SubstringSimilarity("test", "test") != 1.0 {
		t.Fatal("identical should be 1.0")
	}
}

func TestSubstringSimilarityRange(t *testing.T) {
	s := SubstringSimilarity("hello", "world")
	if s < 0 || s > 1 {
		t.Fatalf("out of range: %v", s)
	}
}

func TestLCSSubsequenceIdentical(t *testing.T) {
	if LongestCommonSubsequence("hello", "hello") != 5 {
		t.Fatal("identical should return full length")
	}
}

func TestLCSSubsequenceKnown(t *testing.T) {
	l := LongestCommonSubsequence("ABCBDAB", "BDCAB")
	if l != 4 {
		t.Fatalf("expected 4, got %d", l)
	}
}

func TestSubsequenceSimilarityIdentical(t *testing.T) {
	if math.Abs(SubsequenceSimilarity("abc", "abc")-1.0) > 1e-9 {
		t.Fatal("identical should be 1.0")
	}
}

func TestSubsequenceSimilaritySymmetry(t *testing.T) {
	s1 := SubsequenceSimilarity("hello", "halo")
	s2 := SubsequenceSimilarity("halo", "hello")
	if math.Abs(s1-s2) > 1e-9 {
		t.Fatalf("not symmetric: %v != %v", s1, s2)
	}
}

func TestSubstringExtractBasic(t *testing.T) {
	s := SubstringExtract("abcdef", "xbcdy")
	if s != "bcd" {
		t.Fatalf("expected 'bcd', got %q", s)
	}
}

func TestSubstringExtractEmpty(t *testing.T) {
	s := SubstringExtract("abc", "xyz")
	if s != "" {
		t.Fatalf("expected empty, got %q", s)
	}
}

func TestSubstringExtractUnicode(t *testing.T) {
	s := SubstringExtract("你好世界再见", "欢迎世界你好")
	if s != "世界" && s != "你好" {
		t.Fatalf("expected '世界' or '你好', got %q", s)
	}
	if len([]rune(s)) != 2 {
		t.Fatalf("expected length 2, got %d", len([]rune(s)))
	}
}
