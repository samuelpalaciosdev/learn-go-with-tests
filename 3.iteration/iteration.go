package iteration

func Repeat(character string, count int) string {
	var repeated string

	if count == 0 {
		return character
	}
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
