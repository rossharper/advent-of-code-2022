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

	result := highestCaloriesCarriedByAnElf("input.txt")

	fmt.Println(strconv.Itoa(result))
}

func highestCaloriesCarriedByAnElf(filename string) int {
	f, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(f)

	highestCalories := 0
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			currentCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentCalories += calories
			if currentCalories > highestCalories {
				highestCalories = currentCalories
			}
		}
	}

	f.Close()

	return highestCalories
}
