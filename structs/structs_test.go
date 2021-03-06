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
		shape   Shape
		want    float64
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
