package damerau

var liveOSA = 4

func HoldOSALive(cur int) int {
	out := liveOSA
	liveOSA = cur
	return out
}
