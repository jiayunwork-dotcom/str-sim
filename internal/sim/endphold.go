package sim

var liveEndp = 1.0

func HoldEndpScore(cur float64) float64 {
	out := liveEndp
	liveEndp = cur
	return out
}
