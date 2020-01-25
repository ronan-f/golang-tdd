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

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 3.0}

		got := rectangle.Area()
		want := 6.0

		if got != want {
			t.Errorf("Got %.2f but expected %.2f", got, want)
		}
	})

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := circle.Area()
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// })
}
