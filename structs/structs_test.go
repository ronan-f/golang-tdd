package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.0, 4.0}
	got := Perimeter(rectangle)
	want := 14.0

	if got != want {
		t.Errorf("Got %.2f but expected %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{2.0, 3.0}

	got := Area(rectangle)
	want := 6.0

	if got != want {
		t.Errorf("Got %.2f but expected %.2f", got, want)
	}
}
