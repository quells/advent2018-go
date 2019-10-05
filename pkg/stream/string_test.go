package stream

import (
	"reflect"
	"testing"

	"github.com/quells/advent2018-go/pkg/strconv"
)

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func onlySignedInts(s string) bool {
	_, err := strconv.ParseSignedInt(s)
	return err == nil
}

func TestString_Collect(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		result := OfStrings().Collect()
		if len(result) != 0 {
			t.Errorf("s.Collect() = %v, want []", result)
		}
	})

	t.Run("noop", func(t *testing.T) {
		values := []string{"", "a", "bc", "123"}
		result := OfStrings(values...).Collect()
		if !reflect.DeepEqual(result, values) {
			t.Errorf("s.Collect() = %v, want %v", result, values)
		}
	})

	t.Run("reverse", func(t *testing.T) {
		values := []string{"", "a", "bc", "123"}
		want := []string{"", "a", "cb", "321"}
		result := OfStrings(values...).Map(reverse).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})

	t.Run("filter", func(t *testing.T) {
		values := []string{"abc", "123"}
		want := []string{"123"}
		result := OfStrings(values...).Filter(onlySignedInts).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})

	t.Run("filter reverse", func(t *testing.T) {
		values := []string{"", "a", "bc", "123"}
		want := []string{"321"}
		result := OfStrings(values...).Filter(onlySignedInts).Map(reverse).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})

	t.Run("reverse filter", func(t *testing.T) {
		values := []string{"", "a", "bc", "123"}
		want := []string{"321"}
		result := OfStrings(values...).Map(reverse).Filter(onlySignedInts).Collect()
		if !reflect.DeepEqual(result, want) {
			t.Errorf("s.Collect() = %v, want %v", result, want)
		}
	})
}
