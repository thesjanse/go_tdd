package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("say hello to a person", func(t *testing.T) {
		got := Hello("Jane", "")
		want := "Hello, Jane"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Paul", "French")
		want := "Bonjour, Paul"
		assertCorrectMessage(t, got, want)
	})
}
