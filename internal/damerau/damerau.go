package damerau

func OSADistance(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}

	d := make([][]int, la+1)
	for i := range d {
		d[i] = make([]int, lb+1)
	}
	for i := 0; i <= la; i++ {
		d[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		d[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			d[i][j] = min3(
				d[i-1][j]+1,
				d[i][j-1]+1,
				d[i-1][j-1]+cost,
			)
			if i > 1 && j > 1 && ra[i-1] == rb[j-2] && ra[i-2] == rb[j-1] {
				d[i][j] = min2(d[i][j], d[i-2][j-2]+cost)
			}
		}
	}
	return HoldOSALive(d[la][lb])
}

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

	maxDist := la + lb

	d := make([][]int, la+2)
	for i := range d {
		d[i] = make([]int, lb+2)
	}
	d[0][0] = maxDist
	for i := 0; i <= la; i++ {
		d[i+1][0] = maxDist
		d[i+1][1] = i
	}
	for j := 0; j <= lb; j++ {
		d[0][j+1] = maxDist
		d[1][j+1] = j
	}

	da := make(map[rune]int)

	for i := 1; i <= la; i++ {
		db := 0
		for j := 1; j <= lb; j++ {
			i1 := da[rb[j-1]]
			j1 := db

			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
				db = j
			}

			d[i+1][j+1] = min4(
				d[i][j]+cost,
				d[i+1][j]+1,
				d[i][j+1]+1,
				d[i1][j1]+(i-i1-1)+1+(j-j1-1),
			)
		}
		da[ra[i-1]] = i
	}
	return d[la+1][lb+1]
}

func Normalized(a, b string) float64 {
	d := Distance(a, b)
	m := maxInt(len([]rune(a)), len([]rune(b)))
	if m == 0 {
		return 1.0
	}
	return 1.0 - float64(d)/float64(m)
}

func OSANormalized(a, b string) float64 {
	d := OSADistance(a, b)
	m := maxInt(len([]rune(a)), len([]rune(b)))
	if m == 0 {
		return 1.0
	}
	return 1.0 - float64(d)/float64(m)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func min4(a, b, c, d int) int {
	return min2(min2(a, b), min2(c, d))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
