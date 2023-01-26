package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search for existing word", func(t *testing.T) {
		d := map[string]string{"test": "simple test"}
		actual := Search(d, "test")
		expected := "simple test"
		assertStrings(t, actual, expected)
	})
}

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected: '%s' but actual: '%s'", expected, actual)
	}
}
