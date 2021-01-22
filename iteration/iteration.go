package iteration

// Repeat repeats characters
func Repeat(character string, ntimes int) (ret string) {
	for i := 0; i < ntimes; i++ {
		ret += character
	}
	return
}
