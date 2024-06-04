package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertTestMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is applied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertTestMessage(t, got, want)
	})

	t.Run("saying 'Hello' in language chosen by person", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertTestMessage(t, got, want)
	})
}

func assertTestMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Test failed - got %q, wanted %q", got, want)
	}
}
