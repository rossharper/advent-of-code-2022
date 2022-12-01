package main

import (
	"testing"
)

func TestSumOfTopThreeHighestCalorieCounts(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"FirstThree.txt", 1480},
		{"LastThree.txt", 1480},
		{"RandomThree.txt", 1480},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			ans := sumOfTopThreeHighestCalorieCounts(tt.file)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
