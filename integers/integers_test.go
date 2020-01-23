package integers

import "testing"

func TestAdd(t *testing.T) {
	got := add(1, 1)
	want := 2

	if got != want {
		t.Errorf("Expected %d but received %d", want, got)
	}
}
