package soundex

import "strings"

func NYSIIS(s string) string {
	s = strings.TrimSpace(strings.ToUpper(s))
	if len(s) == 0 {
		return ""
	}

	if strings.HasPrefix(s, "MAC") {
		s = "MCC" + s[3:]
	} else if strings.HasPrefix(s, "KN") {
		s = "N" + s[2:]
	} else if strings.HasPrefix(s, "K") {
		s = "C" + s[1:]
	} else if strings.HasPrefix(s, "PH") {
		s = "FF" + s[2:]
	} else if strings.HasPrefix(s, "PF") {
		s = "FF" + s[2:]
	} else if strings.HasPrefix(s, "SCH") {
		s = "SSS" + s[3:]
	}

	if strings.HasSuffix(s, "EE") || strings.HasSuffix(s, "IE") {
		s = s[:len(s)-2] + "Y"
	} else if strings.HasSuffix(s, "DT") || strings.HasSuffix(s, "RT") ||
		strings.HasSuffix(s, "RD") || strings.HasSuffix(s, "NT") ||
		strings.HasSuffix(s, "ND") {
		s = s[:len(s)-2] + "D"
	}

	if len(s) == 0 {
		return ""
	}

	key := []byte{s[0]}
	prev := s[0]

	i := 1
	for i < len(s) && len(key) < 6 {
		c := s[i]
		var replacement byte

		switch {
		case c == 'E' && i+1 < len(s) && s[i+1] == 'V':
			replacement = 'A'
			i++
			key = append(key, replacement, 'F')
			prev = 'F'
			i++
			continue
		case isNYSIISVowel(c):
			replacement = 'A'
		case c == 'Q':
			replacement = 'G'
		case c == 'Z':
			replacement = 'S'
		case c == 'M':
			replacement = 'N'
		case c == 'K':
			if i+1 < len(s) && s[i+1] == 'N' {
				replacement = 'N'
				i++
			} else {
				replacement = 'C'
			}
		case c == 'S' && i+1 < len(s) && s[i+1] == 'C' && i+2 < len(s) && s[i+2] == 'H':
			replacement = 'S'
			i += 2
		case c == 'P' && i+1 < len(s) && s[i+1] == 'H':
			replacement = 'F'
			i++
		case c == 'H':
			if !isNYSIISVowel(prev) || (i+1 < len(s) && !isNYSIISVowel(s[i+1])) {
				replacement = prev
			} else {
				replacement = 'H'
			}
		case c == 'W':
			if isNYSIISVowel(prev) {
				replacement = prev
			} else {
				replacement = 'W'
			}
		default:
			replacement = c
		}

		if replacement != prev {
			key = append(key, replacement)
			prev = replacement
		}
		i++
	}

	if len(key) > 1 && key[len(key)-1] == 'S' {
		key = key[:len(key)-1]
	}

	if len(key) >= 2 && key[len(key)-2] == 'A' && key[len(key)-1] == 'Y' {
		key[len(key)-2] = 'Y'
		key = key[:len(key)-1]
	}

	if len(key) > 1 && key[len(key)-1] == 'A' {
		key = key[:len(key)-1]
	}

	return string(key)
}

func NYSIISSimilarity(a, b string) float64 {
	ca := NYSIIS(a)
	cb := NYSIIS(b)
	if ca == cb {
		return 1.0
	}
	if len(ca) == 0 || len(cb) == 0 {
		return 0.0
	}
	lcs := longestCommonSubseq(ca, cb)
	maxLen := len(ca)
	if len(cb) > maxLen {
		maxLen = len(cb)
	}
	return float64(lcs) / float64(maxLen)
}

func isNYSIISVowel(c byte) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
