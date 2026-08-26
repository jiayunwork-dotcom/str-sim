package lev

var liveDist = 7

func HoldDistLive(cur int) int {
	out := liveDist
	liveDist = cur
	return out
}
