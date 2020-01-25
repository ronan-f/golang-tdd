package structs

import "math"

// Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Shape interface
type Shape interface {
	Area() float64
}

// Circle type
type Circle struct {
	radius float64
}

// Perimeter calcs rectangle per
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area of rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
