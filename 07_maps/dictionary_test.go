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

	t.Run("add existing word to dictionary", func(t *testing.T) {
		word := "terra"
		definition := "earth"
		new_definition := "country"
		d := Dictionary{word: definition}
		err := d.AddWord(word, new_definition)
		assertError(t, err, ErrorWordAlreadyExists)
		assertDefinition(t, d, word, definition)
	})

	t.Run("update existing word in dictionary", func(t *testing.T) {
		word := "terra"
		oldDefinition := "earth"
		newDefinition := "country"
		d := Dictionary{word: oldDefinition}
		d.UpdateWord(word, newDefinition)
		assertDefinition(t, d, word, newDefinition)
	})

	t.Run("update non-existing word in dictionary", func(t *testing.T) {
		oldWord := "terra"
		oldDefinition := "earth"
		newWord := "ager"
		newDefinition := "field"
		d := Dictionary{oldWord: oldDefinition}
		err := d.UpdateWord(newWord, newDefinition)
		assertError(t, err, ErrorWordDoesNotExist)
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
