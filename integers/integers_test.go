package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	want := 2

	if got != want {
		t.Errorf("Expected %d but received %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
