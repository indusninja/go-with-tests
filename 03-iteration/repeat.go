package _3_iteration

// Repeat takes a character and repeats it number of times as specified by repeatCount
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
