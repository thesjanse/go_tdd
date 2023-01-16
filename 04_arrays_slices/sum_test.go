package slice

import "testing"

func assertCorrectMessage(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%d' but got '%d' from '%v'", expected, actual, numbers)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of 3 elements", func(t *testing.T) {
		numbers := []int{10, 20, 33}
		actual := Sum(numbers)
		expected := 63
		assertCorrectMessage(t, actual, expected, numbers)
	})
}
