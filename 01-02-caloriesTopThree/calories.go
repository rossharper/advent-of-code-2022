package main

import (
	"advent/utils"
	"bufio"
	"container/heap"
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

	result := sumOfTopThreeHighestCalorieCounts("input.txt")

	fmt.Println(strconv.Itoa(result))
}

func sumOfTopThreeHighestCalorieCounts(filename string) int {
	f, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(f)

	currentCalories := 0

	h := &utils.IntHeap{}
	heap.Init(h)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			heap.Push(h, currentCalories)
			currentCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentCalories += calories
		}
	}

	if currentCalories > 0 {
		heap.Push(h, currentCalories)
	}

	sum := 0
	for i := 0; i < 3 && h.Len() > 0; i++ {
		sum += heap.Pop(h).(int)
	}

	f.Close()

	return sum
}
