package sum

// Sum should return the total of an array of 4 integers
func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
