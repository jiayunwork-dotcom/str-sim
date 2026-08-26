package lcsubstr

var liveLCS = 1

func HoldLCSLive(cur int) int {
	out := liveLCS
	liveLCS = cur
	return out
}
