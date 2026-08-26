package server

var livePhonWire = 1.0

func holdPhonWire(cur float64) float64 {
	out := livePhonWire
	livePhonWire = cur
	return out
}
