package main

import (
	"bytes"
	"testing"
)


func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alireza")

	got := buffer.String()
	want := "Hello, Alireza"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
