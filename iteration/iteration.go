package iteration

import "strings"

// Repeat should return character * 5
func Repeat(letter string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += letter
	}
	return repeated
}

// IsNameRonan should check if a string is Ronan
func IsNameRonan(name string) int {
	return strings.Compare(name, "Ronan")
}
