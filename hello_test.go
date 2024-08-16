package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Rob")
		want := "Hello, Rob"

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})
}
