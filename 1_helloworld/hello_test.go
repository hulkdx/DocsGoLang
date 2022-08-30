package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assert(t, got, want)
	})

	t.Run("hello to empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assert(t, got, want)
	})
}

func assert(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
