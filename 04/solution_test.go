package main

import (
	"testing"
)

func TestContained(t *testing.T) {
	var tests = []struct {
		pairs string
		want  bool
	}{
		{"2-4,6-8", false},
		{"2-3,4-5", false},
		{"5-7,7-9", false},
		{"2-8,3-7", true},
		{"6-6,4-6", true},
		{"2-6,4-8", false},
	}

	for _, tt := range tests {
		t.Run(tt.pairs, func(t *testing.T) {
			ans := contained(tt.pairs)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestSolution(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"test.txt", 2},
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
