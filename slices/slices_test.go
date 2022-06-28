package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6
		if got != expected {
			t.Errorf("got %d expected %d given, %v", got, expected, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("sum two arrays", func(t *testing.T) {
		numbersA := []int{1, 2, 3}
		numbersB := []int{1, 2, 3}
		got := SumAll(numbersA, numbersB)
		expected := []int{6, 6}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %d expected %d given, %v, %v", got, expected, numbersA, numbersB)
		}

	})

	// t.Run("sum one array", func(t *testing.T) {
	// 	numbersB := []int{1, 2, 3}
	// 	got := SumAll(numbersB)
	// 	expected := []int{6}
	// 	for i := 0; i < len(got); i++ {
	// 		if got[i] != expected[i] {
	// 			t.Errorf("got %d expected %d given, %v", got, expected, numbersB)
	// 		}
	// 	}
	// })

}
