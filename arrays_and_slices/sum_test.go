package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers,[1,3,4,3,4]", func(t *testing.T) {
		numbers := []int{1, 3, 4, 3, 4}

		got := SumNumbers(numbers)

		want := 15

		if got != want {
			t.Errorf("got %d want %d and given %v", got, want, numbers)
		}
	})
	t.Run("Sum of 0", func(t *testing.T) {
		given := []int{0}
		got := SumNumbers(given)

		want := 0

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, given)
		}
	})
}
