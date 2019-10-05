package main

import (
	"testing"

	"github.com/quells/advent2018-go/pkg/file"
	"github.com/quells/advent2018-go/pkg/stream"
	"github.com/quells/advent2018-go/pkg/strings"
)

func Test02(t *testing.T) {
	f, err := file.Open("inputs/02.txt")
	if err != nil {
		t.Fatalf("could not open input file: %v", err)
	}
	defer f.Close()

	lines, err := file.ReadLines(f)
	if err != nil {
		t.Fatalf("could not read lines: %v", err)
	}

	t.Run("A", func(t *testing.T) {
		charCounts := stream.OfStrings(lines...).ToAnys().Map(func(i interface{}) interface{} {
			s := i.(string)
			m := make(map[rune]int)
			for _, r := range s {
				m[r]++
			}
			return m
		})

		hasCharCount := func(count int) stream.AnyFilter {
			return func(i interface{}) bool {
				m := i.(map[rune]int)
				for _, v := range m {
					if v == count {
						return true
					}
				}
				return false
			}
		}

		withTwo := charCounts.Filter(hasCharCount(2)).Collect()
		withThree := charCounts.Filter(hasCharCount(3)).Collect()

		checksum := len(withTwo) * len(withThree)
		want := 5952
		if checksum != want {
			t.Errorf("checksum = %v, want %v", checksum, want)
		}
	})

	t.Run("B", func(t *testing.T) {
		var found bool
	outer:
		for i := 0; i < len(lines)-1; i++ {
			for j := i + 1; j < len(lines); j++ {
				if strings.CharsDifferCount(lines[i], lines[j]) == 1 {
					found = true
					a := lines[i]
					b := lines[j]
					for x := 0; x < len(a); x++ {
						if a[x] != b[x] {
							result := a[:x] + a[x+1:]
							want := "krdmtuqjgwfoevnaboxglzjph"
							if result != want {
								t.Errorf("common letters = %v, want %v", result, want)
							}
							break
						}
					}
					break outer
				}
			}
		}
		if !found {
			t.Error("could not find lines which differed by single character")
		}
	})
}
