package iteration

// Repeat takes a character and repeats it number of times as specified by repeatCountcd iter
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
