package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Ronan")
	want := "Hello Ronan"

	if got != want {
		t.Errorf("Expected %q but received %q", got, want)
	}
}
