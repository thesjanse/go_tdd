package smi

import "testing"

func assert(t testing.TB, actual, expected float64) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%.1f' but got '%.1f'", expected, actual)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		actual := Perimeter(Rectangle{10.0, 10.0})
		expected := 40.0
		assert(t, actual, expected)
	})
}

func TestArea(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		actual := Area(Rectangle{10.0, 10.0})
		expected := 100.0
		assert(t, actual, expected)
	})
}
