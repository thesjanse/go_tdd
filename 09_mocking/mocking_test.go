package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Test single element print", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		actual := buffer.String()
		expected := `3
2
1
Go!`
		if actual != expected {
			t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
		}
	})

	t.Run("Sleep after every write", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		expected := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(expected, spySleepPrinter.Calls) {
			t.Errorf("Expected: '%v'; actual: '%v'", expected, spySleepPrinter.Calls)
		}
	})
}
