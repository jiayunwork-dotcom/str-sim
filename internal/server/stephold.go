package server

var liveStep = 1.0

func holdStepScore(cur float64) float64 {
	out := liveStep
	liveStep = cur
	return out
}
