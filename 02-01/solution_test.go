package main

import (
	"testing"
)

func TestIndividualRounds(t *testing.T) {
	var tests = []struct {
		round string
		want  uint
	}{
		{"A Y", 8},
		{"B X", 1},
		{"C Z", 6},
		{"A X", 4},
		{"A Z", 3},
		{"B Y", 5},
		{"B Z", 9},
		{"C Y", 2},
		{"C X", 7},
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
		{"test.txt", 15},
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
