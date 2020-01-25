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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
