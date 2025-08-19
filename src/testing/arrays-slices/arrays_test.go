package arraysslices

import "testing"

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
}
