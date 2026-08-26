package damerau

type Alignment struct {
	AlignedA string
	AlignedB string
	Ops      []EditOp
	Distance int
}

type EditOp struct {
	Type  OpType
	PosA  int
	PosB  int
	CharA rune
	CharB rune
}

type OpType int

const (
	OpMatch OpType = iota
	OpSubstitute
	OpInsert
	OpDelete
	OpTranspose
)

func (o OpType) String() string {
	switch o {
	case OpMatch:
		return "match"
	case OpSubstitute:
		return "sub"
	case OpInsert:
		return "ins"
	case OpDelete:
		return "del"
	case OpTranspose:
		return "trans"
	default:
		return "?"
	}
}

func Align(a, b string) Alignment {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)

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

	var ops []EditOp
	i, j := la, lb
	for i > 0 || j > 0 {
		if i > 1 && j > 1 && ra[i-1] == rb[j-2] && ra[i-2] == rb[j-1] {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			if d[i][j] == d[i-2][j-2]+cost {
				ops = append(ops, EditOp{Type: OpTranspose, PosA: i - 2, PosB: j - 2, CharA: ra[i-2], CharB: rb[j-2]})
				i -= 2
				j -= 2
				continue
			}
		}
		if i > 0 && j > 0 && ra[i-1] == rb[j-1] && d[i][j] == d[i-1][j-1] {
			ops = append(ops, EditOp{Type: OpMatch, PosA: i - 1, PosB: j - 1, CharA: ra[i-1], CharB: rb[j-1]})
			i--
			j--
		} else if i > 0 && j > 0 && d[i][j] == d[i-1][j-1]+1 {
			ops = append(ops, EditOp{Type: OpSubstitute, PosA: i - 1, PosB: j - 1, CharA: ra[i-1], CharB: rb[j-1]})
			i--
			j--
		} else if j > 0 && d[i][j] == d[i][j-1]+1 {
			ops = append(ops, EditOp{Type: OpInsert, PosA: -1, PosB: j - 1, CharB: rb[j-1]})
			j--
		} else if i > 0 && d[i][j] == d[i-1][j]+1 {
			ops = append(ops, EditOp{Type: OpDelete, PosA: i - 1, PosB: -1, CharA: ra[i-1]})
			i--
		} else {
			if i > 0 {
				ops = append(ops, EditOp{Type: OpDelete, PosA: i - 1, PosB: -1, CharA: ra[i-1]})
				i--
			} else {
				ops = append(ops, EditOp{Type: OpInsert, PosA: -1, PosB: j - 1, CharB: rb[j-1]})
				j--
			}
		}
	}

	for left, right := 0, len(ops)-1; left < right; left, right = left+1, right-1 {
		ops[left], ops[right] = ops[right], ops[left]
	}

	var alignA, alignB []rune
	for _, op := range ops {
		switch op.Type {
		case OpMatch, OpSubstitute:
			alignA = append(alignA, op.CharA)
			alignB = append(alignB, op.CharB)
		case OpInsert:
			alignA = append(alignA, '-')
			alignB = append(alignB, op.CharB)
		case OpDelete:
			alignA = append(alignA, op.CharA)
			alignB = append(alignB, '-')
		case OpTranspose:
			alignA = append(alignA, op.CharA)
			alignB = append(alignB, op.CharB)
		}
	}

	return Alignment{
		AlignedA: string(alignA),
		AlignedB: string(alignB),
		Ops:      ops,
		Distance: d[la][lb],
	}
}

func EditScript(a, b string) string {
	al := Align(a, b)
	var script []rune
	for _, op := range al.Ops {
		switch op.Type {
		case OpMatch:
			script = append(script, 'M')
		case OpSubstitute:
			script = append(script, 'S')
		case OpInsert:
			script = append(script, 'I')
		case OpDelete:
			script = append(script, 'D')
		case OpTranspose:
			script = append(script, 'T')
		}
	}
	return string(script)
}
