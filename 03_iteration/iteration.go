package iteration

// repeats string n times and returns result as string
func Repeat(str string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += str
	}
	return result
}
