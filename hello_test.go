package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rob")
	want := "Hello, Rob"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
