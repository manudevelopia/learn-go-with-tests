package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Manu")
	want := "Hello, Manu"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
