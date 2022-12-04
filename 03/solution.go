package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partOne := solution("input.txt")
	partTwo := solutionTwo("input.txt")
	fmt.Println(strconv.Itoa(partOne))
	fmt.Println(strconv.Itoa(partTwo))
}

func solution(filename string) int {
	f, err := os.Open(filename)
	utils.Check(err)

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			sum += int(findPriority(line))
		}
	}

	f.Close()

	return sum
}

func solutionTwo(filename string) int {
	f, err := os.Open(filename)
	utils.Check(err)

	scanner := bufio.NewScanner(f)

	sum := 0
	linesRead := 0
	var lines = make([]string, 3)
	for scanner.Scan() {
		lines[linesRead] = scanner.Text()

		if lines[linesRead] != "" {
			linesRead++
			if linesRead%3 == 0 {
				sum += findBadgePriority(lines)
				linesRead = 0
			}
		}
	}

	f.Close()

	return sum
}

func findPriority(items string) int {
	ans := findCommon(items)
	return priority(ans)
}

func findBadgePriority(rucksacks []string) int {
	ans := findBadge(rucksacks)
	return priority(ans)
}

func findCommon(items string) int {
	compLen := len(items) / 2
	var left = utils.IntSet{}
	var right = utils.IntSet{}
	for i, c := range items {
		if i < compLen {
			left.Add(int(c))
		} else {
			right.Add(int(c))
		}
	}
	intersected := left.Intersect(right)
	for k := range intersected {
		return k
	}
	return -1
}

func priority(c int) int {
	var priority = 0
	if c >= 'a' {
		priority = c - 'a'
	} else {
		priority = (c - 'A') + 26
	}
	return priority + 1
}

func findBadge(rucksacks []string) int {
	var intersected = utils.IntSet{}
	for ruckIdx, s := range rucksacks {
		var rucksack = utils.IntSet{}
		for _, c := range s {
			rucksack.Add(int(c))
		}
		if ruckIdx == 0 {
			intersected = rucksack
		} else {
			intersected = intersected.Intersect(rucksack)
		}
	}
	for k := range intersected {
		return k
	}
	return -1
}
