package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dante")

	got := buffer.String()
	want := "Hello, Dante"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TODO: Notes
// What is the & symbol?
