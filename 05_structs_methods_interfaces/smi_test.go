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
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("Expected '%g' but got '%g'", expected, actual)
		}
	}

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12.0, 6.0}, 36.0},
	}
	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.expected)
	}
}
