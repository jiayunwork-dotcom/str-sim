package soundex

import (
	"strings"
	"unicode"
)

func Soundex(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "0000"
	}

	upper := strings.ToUpper(s)
	runes := []rune(upper)

	code := []rune{runes[0]}
	lastDigit := soundexDigit(runes[0])

	for i := 1; i < len(runes) && len(code) < 4; i++ {
		r := runes[i]
		if !unicode.IsLetter(r) {
			lastDigit = '0'
			continue
		}
		d := soundexDigit(r)
		if d == '0' {
			if r != 'H' && r != 'W' {
				lastDigit = '0'
			}
			continue
		}
		if d != lastDigit {
			code = append(code, d)
			lastDigit = d
		}
	}

	for len(code) < 4 {
		code = append(code, '0')
	}
	return string(code[:4])
}

func soundexDigit(r rune) rune {
	switch r {
	case 'B', 'F', 'P', 'V':
		return '1'
	case 'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z':
		return '2'
	case 'D', 'T':
		return '3'
	case 'L':
		return '4'
	case 'M', 'N':
		return '5'
	case 'R':
		return '6'
	default:
		return '0'
	}
}

func SoundexSimilarity(a, b string) float64 {
	ca := Soundex(a)
	cb := Soundex(b)
	if ca == cb {
		return 1.0
	}
	matches := 0
	for i := 0; i < 4; i++ {
		if ca[i] == cb[i] {
			matches++
		}
	}
	return float64(matches) / 4.0
}

func Metaphone(s string) string {
	if len(s) == 0 {
		return ""
	}

	upper := strings.ToUpper(strings.TrimSpace(s))
	runes := []rune(upper)
	n := len(runes)

	if n >= 2 {
		switch string(runes[:2]) {
		case "AE", "GN", "KN", "PN", "WR":
			runes = runes[1:]
			n--
		}
	}

	var result []rune
	for i := 0; i < n && len(result) < 6; i++ {
		c := runes[i]
		if !unicode.IsLetter(c) {
			continue
		}

		if i > 0 && c == runes[i-1] && c != 'C' {
			continue
		}

		switch c {
		case 'A', 'E', 'I', 'O', 'U':
			if i == 0 {
				result = append(result, c)
			}
		case 'B':
			if !(i == n-1 && i > 0 && runes[i-1] == 'M') {
				result = append(result, 'B')
			}
		case 'C':
			if i+1 < n && (runes[i+1] == 'I' || runes[i+1] == 'E' || runes[i+1] == 'Y') {
				result = append(result, 'S')
			} else {
				result = append(result, 'K')
			}
		case 'D':
			if i+1 < n && runes[i+1] == 'G' {
				if i+2 < n && (runes[i+2] == 'I' || runes[i+2] == 'E' || runes[i+2] == 'Y') {
					result = append(result, 'J')
				} else {
					result = append(result, 'T')
				}
			} else {
				result = append(result, 'T')
			}
		case 'F':
			result = append(result, 'F')
		case 'G':
			silent := false
			if i+1 < n && runes[i+1] == 'H' && i+2 < n && !isVowel(runes[i+2]) {
				silent = true
			}
			if i == n-1 && (i > 0 && runes[i-1] == 'I' || i > 0 && runes[i-1] == 'N') {
				silent = true
			}
			if !silent {
				if i+1 < n && (runes[i+1] == 'I' || runes[i+1] == 'E' || runes[i+1] == 'Y') {
					result = append(result, 'J')
				} else {
					result = append(result, 'K')
				}
			}
		case 'H':
			if i+1 < n && isVowel(runes[i+1]) {
				if i == 0 || !isVowel(runes[i-1]) {
					result = append(result, 'H')
				}
			}
		case 'J':
			result = append(result, 'J')
		case 'K':
			if i == 0 || runes[i-1] != 'C' {
				result = append(result, 'K')
			}
		case 'L':
			result = append(result, 'L')
		case 'M':
			result = append(result, 'M')
		case 'N':
			result = append(result, 'N')
		case 'P':
			if i+1 < n && runes[i+1] == 'H' {
				result = append(result, 'F')
				i++
			} else {
				result = append(result, 'P')
			}
		case 'Q':
			result = append(result, 'K')
		case 'R':
			result = append(result, 'R')
		case 'S':
			if i+1 < n && runes[i+1] == 'H' {
				result = append(result, 'X')
				i++
			} else if i+2 < n && runes[i+1] == 'I' && (runes[i+2] == 'O' || runes[i+2] == 'A') {
				result = append(result, 'X')
				i += 2
			} else {
				result = append(result, 'S')
			}
		case 'T':
			if i+1 < n && runes[i+1] == 'H' {
				result = append(result, '0')
				i++
			} else if i+2 < n && runes[i+1] == 'I' && (runes[i+2] == 'O' || runes[i+2] == 'A') {
				result = append(result, 'X')
				i += 2
			} else {
				result = append(result, 'T')
			}
		case 'V':
			result = append(result, 'F')
		case 'W', 'Y':
			if i+1 < n && isVowel(runes[i+1]) {
				result = append(result, c)
			}
		case 'X':
			result = append(result, 'K')
			result = append(result, 'S')
		case 'Z':
			result = append(result, 'S')
		}
	}
	return string(result)
}

func MetaphoneSimilarity(a, b string) float64 {
	ma := Metaphone(a)
	mb := Metaphone(b)
	if ma == mb {
		return 1.0
	}
	if len(ma) == 0 || len(mb) == 0 {
		return 0.0
	}
	lcs := longestCommonSubseq(ma, mb)
	maxLen := len(ma)
	if len(mb) > maxLen {
		maxLen = len(mb)
	}
	return float64(lcs) / float64(maxLen)
}

func longestCommonSubseq(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	prev := make([]int, lb+1)
	for i := 1; i <= la; i++ {
		curr := make([]int, lb+1)
		for j := 1; j <= lb; j++ {
			if ra[i-1] == rb[j-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max2(prev[j], curr[j-1])
			}
		}
		prev = curr
	}
	return prev[lb]
}

func isVowel(r rune) bool {
	switch r {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
