package iteration

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	given := Repeat("a", 5)
	expected := "aaaaa"

	if given != expected {
		t.Errorf("Expected: '%s', Received: '%s'", expected, given);
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i<b.N; i++ {
		Repeat("a", 5);
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
