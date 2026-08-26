package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) != len(rb) {
		return 0, fmt.Errorf("hamming: strings must be equal length (got %d and %d)", len(ra), len(rb))
	}
	d := 0
	for i := range ra {
		if ra[i] != rb[i] {
			d++
		}
	}
	return d, nil
}

func Normalized(a, b string) (float64, error) {
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) != len(rb) {
		return 0, fmt.Errorf("hamming: strings must be equal length (got %d and %d)", len(ra), len(rb))
	}
	if len(ra) == 0 {
		return 1.0, nil
	}
	d, _ := Distance(a, b)
	return 1.0 - float64(d)/float64(len(ra)), nil
}

func BitDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("hamming: byte slices must be equal length (got %d and %d)", len(a), len(b))
	}
	count := 0
	for i := range a {
		xor := a[i] ^ b[i]
		count += popcount(xor)
	}
	return count, nil
}

func popcount(b byte) int {
	count := 0
	for b != 0 {
		count += int(b & 1)
		b >>= 1
	}
	return count
}

func PaddedDistance(a, b string, pad rune) int {
	ra := []rune(a)
	rb := []rune(b)
	maxLen := len(ra)
	if len(rb) > maxLen {
		maxLen = len(rb)
	}
	for len(ra) < maxLen {
		ra = append(ra, pad)
	}
	for len(rb) < maxLen {
		rb = append(rb, pad)
	}
	d := 0
	for i := range ra {
		if ra[i] != rb[i] {
			d++
		}
	}
	return d
}

func PaddedNormalized(a, b string, pad rune) float64 {
	ra := []rune(a)
	rb := []rune(b)
	maxLen := len(ra)
	if len(rb) > maxLen {
		maxLen = len(rb)
	}
	if maxLen == 0 {
		return 1.0
	}
	d := PaddedDistance(a, b, pad)
	return 1.0 - float64(d)/float64(maxLen)
}
