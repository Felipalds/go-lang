package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := SumArray(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d exp %d, from %v", got, expected, numbers)
		}
	})
	t.Run("slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := SumSlice(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d exp %d, from %v", got, expected, numbers)
		}
	})
	t.Run("sum all", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 4}
		got := SumAllSmart(slice1, slice2)
		expected := []int{3, 7}

		// see slices.Equal
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %d exp %d", got, expected)
		}

	})
}

func BenchmarkSumAllSmart(b *testing.B) {
	for b.Loop() {
		SumAllSmart([]int{1, 2, 3}, []int{4, 5, 6})
	}
}

func BenchmarkSumAllDumb(b *testing.B) {
	for b.Loop() {
		SumAllDumb([]int{1, 2, 3}, []int{4, 5, 6})
	}
}
