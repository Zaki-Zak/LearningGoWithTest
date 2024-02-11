package iteration

func Repeat(character string, reps int) string {
	var repeated string
	for i := 0; i < reps; i++ {
		repeated += character
	}
	return repeated
}
