package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("search for existing word", func(t *testing.T) {
		d := Dictionary{"test": "simple test"}
		actual, _ := d.Search("test")
		expected := "simple test"
		assertStrings(t, actual, expected)
	})

	t.Run("search for unknown word", func(t *testing.T) {
		d := Dictionary{"test": "simple test"}
		_, actual := d.Search("unknown")
		expected := ErrorWordNotFound
		assertError(t, actual, expected)
	})

	t.Run("add new word to dictionary", func(t *testing.T) {
		d := Dictionary{"test": "simple test"}
		word := "terra"
		expected := "earth"
		d.AddWord(word, expected)
		assertDefinition(t, d, word, expected)
	})
}

func assertDefinition(t testing.TB, d Dictionary, word, expected string) {
	t.Helper()
	actual, err := d.Search(word)
	if err != nil {
		t.Fatal("Didn't find a word:", err)
	}
	assertStrings(t, actual, expected)
}

func assertError(t testing.TB, actual, expected error) {
	t.Helper()
	if actual != expected {
		t.Errorf("Actual '%q' error, but expected '%q'", actual, expected)
	}
}

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected: '%s' but actual: '%s'", expected, actual)
	}
}
