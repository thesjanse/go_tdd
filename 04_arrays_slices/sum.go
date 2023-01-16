package slice

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	result := make([]int, 0)
	for _, slice := range slices {
		result = append(result, Sum(slice))
	}
	return result
}
