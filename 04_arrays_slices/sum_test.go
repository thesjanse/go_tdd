package slice

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%d' but got '%d' from '%v'", expected, actual, numbers)
	}
}

func assertSliceMessage(t testing.TB, actual, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v' but got '%v", expected, actual)
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

func TestSumAll(t *testing.T) {
	t.Run("single slice", func(t *testing.T) {
		actual := SumAll([]int{1, 1, 1})
		expected := []int{3}
		assertSliceMessage(t, actual, expected)
	})

	t.Run("multiple slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		assertSliceMessage(t, actual, expected)
	})
}
