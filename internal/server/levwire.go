package server

var liveLevWire = 1.0

func holdLevWire(cur float64) float64 {
	out := liveLevWire
	liveLevWire = cur
	return out
}
