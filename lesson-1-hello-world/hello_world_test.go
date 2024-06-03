package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("my twin")
	want := "Hello my twin"

	if got != want {
		t.Errorf("Tests failed, Got %q but wanted %q", got, want)
	}
}
