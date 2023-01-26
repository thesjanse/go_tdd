package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search for existing word", func(t *testing.T) {
		d := map[string]string{"test": "simple test"}
		actual := Search(d, "test")
		expected := "simple test"
		if actual != expected {
			t.Errorf("Expected: '%s' but actual: '%s'", expected, actual)
		}
	})
}
