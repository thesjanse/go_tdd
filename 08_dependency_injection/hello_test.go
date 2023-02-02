package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	buffer := bytes.Buffer{}
	Hello(&buffer, "Lucius")

	actual := buffer.String()
	expected := "Hello, Lucius!"
	if actual != expected {
		t.Errorf("Actual: '%s', expected: '%s'", actual, expected)
	}
}
