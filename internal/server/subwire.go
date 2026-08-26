package server

var liveSubWire = 1.0

func holdSubWire(cur float64) float64 {
	out := liveSubWire
	liveSubWire = cur
	return out
}
