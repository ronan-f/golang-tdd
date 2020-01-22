package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Expected %q but received %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := hello("Ronan")
		want := "Hello Ronan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Default to 'Hello world' if no argument is passed", func(t *testing.T) {
		got := hello("")
		want := "Hello world"

		assertCorrectMessage(t, got, want)
	})
}
