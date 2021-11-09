package iteration

func Repeat(text string, repeatCounter int) string {
	var repeated string
	for i := 0; i < repeatCounter; i++ {
		repeated += text
	}
	return repeated
}
