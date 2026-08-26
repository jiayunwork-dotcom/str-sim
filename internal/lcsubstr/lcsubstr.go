package lcsubstr

func LongestCommonSubstring(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 || lb == 0 {
		return 0
	}

	maxLen := 0
	prev := make([]int, lb+1)
	for i := 1; i <= la; i++ {
		curr := make([]int, lb+1)
		for j := 1; j <= lb; j++ {
			if ra[i-1] == rb[j-1] {
				curr[j] = prev[j-1] + 1
				if curr[j] > maxLen {
					maxLen = curr[j]
				}
			}
		}
		prev = curr
	}
	return maxLen
}

func SubstringSimilarity(a, b string) float64 {
	if a == b {
		return 1.0
	}
	la := len([]rune(a))
	lb := len([]rune(b))
	maxLen := la
	if lb > maxLen {
		maxLen = lb
	}
	if maxLen == 0 {
		return 1.0
	}
	lcs := LongestCommonSubstring(a, b)
	return float64(lcs) / float64(maxLen)
}

func LongestCommonSubsequence(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 || lb == 0 {
		return 0
	}

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

func SubsequenceSimilarity(a, b string) float64 {
	if a == b {
		return 1.0
	}
	la := len([]rune(a))
	lb := len([]rune(b))
	total := la + lb
	if total == 0 {
		return 1.0
	}
	lcs := LongestCommonSubsequence(a, b)
	return 2.0 * float64(lcs) / float64(total)
}

func SubstringExtract(a, b string) string {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 || lb == 0 {
		return ""
	}

	maxLen := 0
	endIdx := 0
	prev := make([]int, lb+1)
	for i := 1; i <= la; i++ {
		curr := make([]int, lb+1)
		for j := 1; j <= lb; j++ {
			if ra[i-1] == rb[j-1] {
				curr[j] = prev[j-1] + 1
				if curr[j] > maxLen {
					maxLen = curr[j]
					endIdx = i
				}
			}
		}
		prev = curr
	}
	if maxLen == 0 {
		return ""
	}
	return string(ra[endIdx-maxLen : endIdx])
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
