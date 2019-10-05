package strings

import "testing"

func TestCharsDifferCount(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{"empty", "", "", 0},
		{"one", "fghij", "fguij", 1},
		{"two", "abcde", "axcye", 2},
		{"short long", "abc", "abc123", 3},
		{"long short", "abc123", "abc", 3},
		{"short long diff", "xyz", "abc123", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharsDifferCount(tt.a, tt.b); got != tt.want {
				t.Errorf("CharsDiffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
