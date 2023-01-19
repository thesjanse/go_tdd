package smi

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t testing.TB, actual, expected float64) {
		if actual != expected {
			t.Errorf("Expected '%.1f' but got '%.1f'", expected, actual)
		}
	}
	t.Run("square", func(t *testing.T) {
		actual := Perimeter(10.0, 10.0)
		expected := 40.0
		assert(t, actual, expected)
	})
}
