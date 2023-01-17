package slice

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		result = append(result, Sum(slice))
	}
	return result
}

func SumAllTails(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		if len(slice) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(slice[1:]))
		}
	}
	return result
}
