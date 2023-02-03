package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("Test single element print", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer)

		actual := buffer.String()
		expected := "3"
		if actual != expected {
			t.Errorf("Actual: '%s'; expected: '%s'.", actual, expected)
		}
	})
}
