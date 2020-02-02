package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	got := dictionary.Search("test")
	want := "this is a test"

	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
