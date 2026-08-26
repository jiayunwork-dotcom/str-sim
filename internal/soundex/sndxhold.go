package soundex

var liveSoundex = "A000"

func HoldSoundexLive(cur string) string {
	out := liveSoundex
	liveSoundex = cur
	return out
}
