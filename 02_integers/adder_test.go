package integers

import "testing"

func assertCorrectMessage(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, actual)
	}
}

func TestAdder(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	assertCorrectMessage(t, actual, expected)
}
