package iteration

// repeats character n times and returns result as string
func Repeat(char string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}
