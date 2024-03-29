package slice

import (
	"reflect"
	"testing"
)

func assertSliceMessage(t testing.TB, actual, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected '%v' but got '%v", expected, actual)
	}
}

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, actual, expected int, numbers []int) {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected '%d' but got '%d' from '%v'", expected, actual, numbers)
		}
	}

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

func TestSumAllTails(t *testing.T) {
	t.Run("single slice", func(t *testing.T) {
		actual := SumAllTails([]int{1, 1, 1})
		expected := []int{2}
		assertSliceMessage(t, actual, expected)
	})

	t.Run("equal length slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		assertSliceMessage(t, actual, expected)
	})

	t.Run("inequal length slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 4, 9})
		expected := []int{2, 13}
		assertSliceMessage(t, actual, expected)
	})

	t.Run("empty slice", func(t *testing.T) {
		actual := SumAllTails([]int{})
		expected := []int{0}
		assertSliceMessage(t, actual, expected)
	})

	t.Run("empty tail", func(t *testing.T) {
		actual := SumAllTails([]int{10})
		expected := []int{0}
		assertSliceMessage(t, actual, expected)
	})
}
