package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(10, 15)
	fmt.Println(sum)
	// Output: 25
}

func assertCorrectMessage(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, actual)
	}
}

func TestAdder(t *testing.T) {
	t.Run("Add different numbers", func(t *testing.T) {
		actual := Add(5, 9)
		expected := 14
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("Add identical numbers", func(t *testing.T) {
		actual := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, actual, expected)
	})
}
