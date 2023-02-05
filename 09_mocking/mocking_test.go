package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Test single element print", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)

		actual := buffer.String()
		expected := `3
2
1
Go!`
		if actual != expected {
			t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("Actual calls: '%d', expected calls: '%d'", spySleeper.Calls, 3)
		}
	})
}
