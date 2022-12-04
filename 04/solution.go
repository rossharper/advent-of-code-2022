package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			if contained(line) {
				sum++
			}
		}
	}

	f.Close()

	return sum
}

func contained(input string) bool {
	pairs := strings.Split(input, ",")
	left := pairs[0]
	right := pairs[1]

	return containedWithin(left, right) || containedWithin(right, left)
}

func containedWithin(s1 string, s2 string) bool {
	left := strings.Split(s1, "-")
	right := strings.Split(s2, "-")
	leftStart, _ := strconv.Atoi(left[0])
	leftEnd, _ := strconv.Atoi(left[1])
	rightStart, _ := strconv.Atoi(right[0])
	rightEnd, _ := strconv.Atoi(right[1])
	return leftStart >= rightStart && leftEnd <= rightEnd
}
