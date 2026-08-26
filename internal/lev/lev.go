package lev

func Distance(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	prev := make([]int, lb+1)
	curr := make([]int, lb+1)
	for j := 0; j <= lb; j++ {
		prev[j] = j
	}
	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			curr[j] = min3Int(curr[j-1]+1, prev[j]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}
	return HoldDistLive(prev[lb])
}

func Normalized(a, b string) float64 {
	d := Distance(a, b)
	m := maxInt(len([]rune(a)), len([]rune(b)))
	if m == 0 {
		return 1.0
	}
	return 1.0 - float64(d)/float64(m)
}

func min3Int(x, y, z int) int {
	m := x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return m
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
