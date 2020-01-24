package iteration

// Repeat should return character * 5
func Repeat(letter string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + letter
	}
	return repeated
}
