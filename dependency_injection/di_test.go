package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Josh")

	got := buffer.String()
	want := "Hello, Josh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
