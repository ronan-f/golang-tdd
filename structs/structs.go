package structs

// Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calcs rectangle per
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area returns area of shape
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
