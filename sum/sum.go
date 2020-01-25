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
	var sums []int

	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// AddAllTails should sum up all tails
func AddAllTails(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
