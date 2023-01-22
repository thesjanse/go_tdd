package smi

import "testing"

func assert(t testing.TB, actual, expected float64) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%g' but got '%.g'", expected, actual)
	}
}

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, name string, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Perimeter()
		if actual != expected {
			t.Errorf("%#v expected '%g' but got '%g'", name, expected, actual)
		}
	}
	perimeterTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "square", shape: Rectangle{10.0, 10.0}, expected: 40.0},
		{name: "rectangle", shape: Rectangle{5.0, 10.0}, expected: 30.0},
		{name: "circle", shape: Circle{10.0}, expected: 62.83185307179586},
	}
	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.name, tt.shape, tt.expected)
		})
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, name string, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("Expected '%g' but got '%g'", expected, actual)
		}
	}

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, expected: 100.0},
		{name: "Circle", shape: Circle{10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, expected: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.name, tt.shape, tt.expected)
		})
	}
}
