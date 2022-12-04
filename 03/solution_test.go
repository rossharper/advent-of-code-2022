package main

import (
	"testing"
)

func TestFindCommonItems(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p'},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L'},
		{"PmmdzqPrVvPwwTWBwg", 'P'},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v'},
		{"ttgJtRGJQctTZtZT", 't'},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 's'},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ans, _ := findCommon(tt.input)
			if ans != tt.want {
				t.Errorf("got %c, want %c", ans, tt.want)
			}
		})
	}
}

func TestPriorites(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"PmmdzqPrVvPwwTWBwg", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"ttgJtRGJQctTZtZT", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ans := findPriority(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"partOneTest.txt", 157},
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
