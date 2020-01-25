package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectSum := func(t *testing.T, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	}

	t.Run("Collection of 4 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		result := Sum(numbers)
		expected := 10

		assertCorrectSum(t, result, expected)
	})

	t.Run("Collection of any length", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

		result := Sum(nums)
		expected := 36

		assertCorrectSum(t, result, expected)

	})
}

func TestSumAll(t *testing.T) {
	result := AddAll([]int{1, 2, 3}, []int{2, 3, 4})
	expected := []int{6, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	}

	t.Run("Correctly sums basic tails", func(t *testing.T) {
		result := AddAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		expected := []int{5, 7}

		checkSums(t, result, expected)

	})

	t.Run("Returns 0 if empty slice passed", func(t *testing.T) {
		result := AddAllTails([]int{}, []int{2, 3, 4})
		expected := []int{0, 7}

		checkSums(t, result, expected)
	})
}
