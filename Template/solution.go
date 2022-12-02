package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	result := solution("input.txt")

	fmt.Println(strconv.Itoa(result))
}

func solution(filename string) int {
	f, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(f)

	result := 0
	current := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if result < current {
				result = current
			}
			current = 0
		} else {
			newVal, _ := strconv.Atoi(line)
			current += newVal
		}
	}

	if current > 0 {
		if result < current {
			result = current
		}
	}

	f.Close()

	return result
}
