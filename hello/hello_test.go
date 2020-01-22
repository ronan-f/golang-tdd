package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Hello world"

	if got != want {
		t.Errorf("Expected %q but received %q", got, want)
	}
}
