package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func TestIsNameRonan(t *testing.T) {
	isNameRonan := IsNameRonan("Ronan")
	expected := 0

	if isNameRonan != expected {
		t.Errorf("Expected %d but got %d", expected, isNameRonan)
	}
}

func TestReplaceAllRonans(t *testing.T) {
	replace := ReplaceAllRonans("Ronan, Ronan, Ronan", "Aidan")
	expected := "Aidan, Aidan, Aidan"

	if replace != expected {
		t.Errorf("Expected %q but got %q", expected, replace)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	repeat := Repeat("x", 10)
	fmt.Println(repeat)
	// Output: xxxxxxxxxx
}
