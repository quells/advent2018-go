package main

import (
	"errors"
	"testing"

	"github.com/quells/advent2018-go/pkg/file"
	"github.com/quells/advent2018-go/pkg/strconv"
	"github.com/quells/advent2018-go/pkg/stream"
)

func Test01(t *testing.T) {
	f, err := file.Open("inputs/01.txt")
	if err != nil {
		t.Fatalf("could not open input file: %v", err)
	}
	defer f.Close()

	lines, err := file.ReadLines(f)
	if err != nil {
		t.Fatalf("could not read lines: %v", err)
	}
	values := stream.OfStrings(lines...).ToInts(strconv.MustParseSignedInt)

	t.Run("A", func(t *testing.T) {
		sum := values.Sum()
		want := 486
		if sum != want {
			t.Errorf("sum = %v, want %v", sum, want)
		}
	})

	t.Run("B", func(t *testing.T) {
		seen := make(map[int]bool)
		repeated := values.Repeated().Reduce(0, func(a, b int) (x int, err error) {
			x = a + b
			haveSeen := seen[x]
			if haveSeen {
				err = errors.New("break")
			}
			seen[x] = true
			return
		})
		want := 69285
		if repeated != want {
			t.Errorf("repeated = %v, want %v", repeated, want)
		}
	})
}
