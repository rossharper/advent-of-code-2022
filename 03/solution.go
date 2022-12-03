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

func findPriority(items string) int32 {
	_, ans := findCommon(items)
	return ans
}

func findCommon(items string) (int32, int32) {
	var itemTypeCounts [52]byte
	compLen := len(items) / 2
	for i, c := range items {

		var itemIndex int32 = 0
		if c >= 'a' {
			itemIndex = c - 'a'
		} else {
			itemIndex = (c - 'A') + 26
		}

		if i < compLen {
			itemTypeCounts[itemIndex]++
		} else {
			if itemTypeCounts[itemIndex] > 0 {
				return c, itemIndex + 1
			}
		}
	}
	return -1, -1
}
