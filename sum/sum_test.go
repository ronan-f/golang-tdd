package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := [4]int{1, 2, 3, 4}

	result := Sum(numbers)
	expected := 10

	if result != expected {
		t.Errorf("Expected %d but received %d", expected, result)
	}
}
