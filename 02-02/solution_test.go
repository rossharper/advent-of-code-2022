package main

import (
	"testing"
)

func TestIndividualRounds(t *testing.T) {
	var tests = []struct {
		round string
		want  uint
	}{
		{"A Y", 4},
		{"B X", 1},
		{"C Z", 7},
		{"A X", 3},
		{"A Z", 8},
		{"B Y", 5},
		{"B Z", 9},
		{"C Y", 6},
		{"C X", 2},
	}

	for _, tt := range tests {
		t.Run(tt.round, func(t *testing.T) {
			ans := individualRound(tt.round)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestSolution(t *testing.T) {
	var tests = []struct {
		file string
		want uint
	}{
		{"test.txt", 12},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			ans := solution(tt.file)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
