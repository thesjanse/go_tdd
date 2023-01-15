package slice

import "testing"

func assertCorrectMessage(t testing.TB, actual, expected int, numbers [5]int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%d' but got '%d' from '%v'", expected, actual, numbers)
	}
}

func TestSum(t *testing.T) {
	t.Run("Non-empty list", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expected := 15
		assertCorrectMessage(t, actual, expected, numbers)
	})
}
