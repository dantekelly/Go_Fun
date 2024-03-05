package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("Dante", "en")
		want := "Hello, Dante!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Dante", "es")
		want := "Hola, Dante!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, Wrld!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
