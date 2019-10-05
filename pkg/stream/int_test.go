package stream

import (
	"reflect"
	"testing"
)

func square(x int) int {
	return x * x
}

func isEven(x int) bool {
	return x%2 == 0
}

func TestInt_Collect(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		result := OfInts().Collect()
		if len(result) != 0 {
			t.Errorf("s.Collect() = %v, want []", result)
		}
	})

	t.Run("noop", func(t *testing.T) {
		values := []int{0, 1, 23, 9001}
		result := OfInts(values...).Collect()
		if !reflect.DeepEqual(result, values) {
			t.Errorf("s.Collect() = %v, want %v", result, values)
		}
	})

	t.Run("square", func(t *testing.T) {
		values := []int{0, 1, 23}
		want := []int{0, 1, 529}
		result := OfInts(values...).Map(square).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})

	t.Run("filter", func(t *testing.T) {
		values := []int{0, 1, 2, 34, 5}
		want := []int{0, 2, 34}
		result := OfInts(values...).Filter(isEven).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})

	t.Run("filter square", func(t *testing.T) {
		values := []int{0, 1, 2, 3, 4, 5}
		want := []int{0, 4, 16}
		result := OfInts(values...).Filter(isEven).Map(square).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})

	t.Run("square filter", func(t *testing.T) {
		values := []int{0, 1, 2, 3, 4, 5}
		want := []int{0, 4, 16}
		result := OfInts(values...).Map(square).Filter(isEven).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})
}
