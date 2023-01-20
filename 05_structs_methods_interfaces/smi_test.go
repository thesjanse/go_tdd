package smi

import "testing"

func assert(t testing.TB, actual, expected float64) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%g' but got '%.g'", expected, actual)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		actual := Rectangle{10.0, 10.0}.Perimeter()
		expected := 40.0
		assert(t, actual, expected)
	})
}

func TestArea(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		actual := Rectangle{10.0, 10.0}.Area()
		expected := 100.0
		assert(t, actual, expected)
	})

	t.Run("circle", func(t *testing.T) {
		actual := Circle{10.0}.Area()
		expected := 314.1592653589793
		assert(t, actual, expected)
	})
}
