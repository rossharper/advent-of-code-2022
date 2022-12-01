package main

import (
	"testing"
)

func TestHighestCaloriesCarriedByAnElf(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"testOneElf.txt", 600},
		{"testTwoElfHighestFirst.txt", 600},
		{"testThreeElfHighestLast.txt", 3000},
		{"testThreeElfHighestMiddle.txt", 30000},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			ans := highestCaloriesCarriedByAnElf(tt.file)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
