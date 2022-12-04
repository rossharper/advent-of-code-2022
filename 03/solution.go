package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	result := solution("input.txt")

	fmt.Println(strconv.Itoa(result))
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

func findPriority(items string) int {
	_, ans := findCommon(items)
	return ans
}

func findCommon(items string) (int, int) {
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
		return k, priority(k)
	}
	return -1, -1
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
