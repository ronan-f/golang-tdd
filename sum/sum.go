package sum

// Sum should return the total of an array of 4 integers
func Sum(numbers [4]int) int {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += numbers[i]
	}

	return sum
}
