package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"test.txt", 600},
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
