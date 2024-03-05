package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dante")
	want := "Hello, Dante!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
