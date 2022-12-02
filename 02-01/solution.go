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

	fmt.Println(strconv.FormatUint(uint64(result), 10))
}

func solution(filename string) uint {
	f, err := os.Open(filename)
	utils.Check(err)

	scanner := bufio.NewScanner(f)

	var score uint = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			score += individualRound(line)
		}
	}

	f.Close()

	return score
}

const (
	lose uint = 0
	draw uint = 3
	win  uint = 6
)

var results = map[string]uint{
	"A X": draw,
	"A Y": win,
	"A Z": lose,
	"B X": lose,
	"B Y": draw,
	"B Z": win,
	"C X": win,
	"C Y": lose,
	"C Z": draw,
}

var points = map[byte]uint{
	'X': 1,
	'Y': 2,
	'Z': 3,
}

func individualRound(line string) uint {

	resultPoints := results[line]
	myPlayPoints := points[line[2]]

	return resultPoints + myPlayPoints
}
