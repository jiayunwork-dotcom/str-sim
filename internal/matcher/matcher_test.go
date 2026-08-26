package matcher

import (
	"math"
	"testing"
)

func TestNewMatcherDefault(t *testing.T) {
	m, err := New(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}
	if m == nil {
		t.Fatal("expected non-nil matcher")
	}
}

func TestNewMatcherNoAlgos(t *testing.T) {
	_, err := New(Config{})
	if err == nil {
		t.Fatal("expected error for empty config")
	}
}

func TestScoreIdentical(t *testing.T) {
	m, _ := New(DefaultConfig())
	if m.Score("hello", "hello") != 1.0 {
		t.Fatal("identical strings should score 1.0")
	}
}

func TestScoreSymmetry(t *testing.T) {
	m, _ := New(DefaultConfig())
	s1 := m.Score("kitten", "sitting")
	s2 := m.Score("sitting", "kitten")
	if math.Abs(s1-s2) > 1e-9 {
		t.Fatalf("not symmetric: %v != %v", s1, s2)
	}
}

func TestScoreRange(t *testing.T) {
	m, _ := New(DefaultConfig())
	s := m.Score("hello", "world")
	if s < 0 || s > 1 {
		t.Fatalf("score=%v out of [0,1]", s)
	}
}

func TestMatchThreshold(t *testing.T) {
	m, _ := New(DefaultConfig())
	if !m.Match("test", "test", 0.9) {
		t.Fatal("identical should match at 0.9 threshold")
	}
	if m.Match("abc", "xyz", 0.9) {
		t.Fatal("very different should not match at 0.9")
	}
}

func TestScoreWithDetail(t *testing.T) {
	m, _ := New(DefaultConfig())
	d := m.ScoreWithDetail("hello", "hallo")
	if len(d.Scores) != 3 {
		t.Fatalf("expected 3 algo scores, got %d", len(d.Scores))
	}
	if d.Composite <= 0 || d.Composite > 1 {
		t.Fatalf("composite=%v out of expected range", d.Composite)
	}
	for name, s := range d.Scores {
		if s < 0 || s > 1 {
			t.Fatalf("%s score=%v out of [0,1]", name, s)
		}
	}
}

func TestTopK(t *testing.T) {
	m, _ := New(DefaultConfig())
	candidates := []string{"hello", "help", "world", "hero", "helm"}
	results := m.TopK("hell", candidates, 3)
	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	if results[0].Value != "hello" {
		t.Fatalf("top result = %q, expected hello", results[0].Value)
	}
	for i := 1; i < len(results); i++ {
		if results[i].Score > results[i-1].Score {
			t.Fatalf("results not sorted: [%d]=%v > [%d]=%v",
				i, results[i].Score, i-1, results[i-1].Score)
		}
	}
}

func TestFilter(t *testing.T) {
	m, _ := New(DefaultConfig())
	candidates := []string{"hello", "help", "world", "hero", "xyz"}
	results := m.Filter("hello", candidates, 0.7)
	found := false
	for _, r := range results {
		if r.Value == "hello" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("filter should include exact match 'hello'")
	}
	for _, r := range results {
		if r.Value == "xyz" {
			t.Fatalf("'xyz' should not pass 0.7 threshold, score=%v", r.Score)
		}
	}
}

func TestPhoneticConfig(t *testing.T) {
	m, err := New(PhoneticConfig())
	if err != nil {
		t.Fatal(err)
	}
	s := m.Score("Stephen", "Steven")
	if s < 0.5 {
		t.Fatalf("Stephen/Steven phonetic score = %v, expected >= 0.5", s)
	}
}

func TestFullConfig(t *testing.T) {
	m, err := New(FullConfig())
	if err != nil {
		t.Fatal(err)
	}
	s := m.Score("hello", "hallo")
	if s <= 0 || s >= 1 {
		t.Fatalf("FullConfig score=%v, expected in (0,1)", s)
	}
}
