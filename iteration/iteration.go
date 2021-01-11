package iteration

// Repeat repeats characters
func Repeat(character string) (ret string) {
	for i := 0; i < 5; i++ {
		ret += character
	}
	return
}
