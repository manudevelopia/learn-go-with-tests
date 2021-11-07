package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("Manu", "english")
		want := "Hello, Manu"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("saying hello to people in french", func(t *testing.T) {
		got := Hello("Manu", "french")
		want := "Halo, Manu"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Manu", "spanish")
		want := "Hola, Manu"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("saying hello to people default", func(t *testing.T) {
		got := Hello("Manu", "")
		want := "Hello, Manu"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
