package iteration

import "testing"

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("Five char repeat", func(t *testing.T) {
		actual := Repeat("z", 5)
		expected := "zzzzz"
		assertCorrectMessage(t, actual, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 5)
	}
}
