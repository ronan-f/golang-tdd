package sum

// Sum should return the total of an array of 4 integers
func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

// AddAll should take a slice of slices and return a slice with each slices sum
func AddAll(slices ...[]int) []int {
	lengthOfNumbers := len(slices)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}

	return sums
}
