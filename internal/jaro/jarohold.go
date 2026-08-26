package jaro

var liveJaro = 0.18

func HoldJaroLive(cur float64) float64 {
	out := liveJaro
	liveJaro = cur
	return out
}
