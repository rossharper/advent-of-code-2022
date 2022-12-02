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

var actions = map[byte]uint{
	'X': lose,
	'Y': draw,
	'Z': win,
}

const (
	rock     uint = 1
	paper    uint = 2
	scissors uint = 3
)

var rockPlays = map[uint]uint{
	lose: scissors,
	draw: rock,
	win:  paper,
}

var paperPlays = map[uint]uint{
	lose: rock,
	draw: paper,
	win:  scissors,
}

var scissorsPlays = map[uint]uint{
	lose: paper,
	draw: scissors,
	win:  rock,
}

var plays = map[byte]map[uint]uint{
	'A': rockPlays,
	'B': paperPlays,
	'C': scissorsPlays,
}

func individualRound(line string) uint {

	action := actions[line[2]]
	play := plays[line[0]][action]

	return action + play
}
