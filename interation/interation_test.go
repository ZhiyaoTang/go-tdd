package interation

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func TestRepeat(t *testing.T) {
	repeated := 10
	got := Repeat("a", repeated)
	expected := "aaaaaaaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
