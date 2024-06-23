package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	up   = '('
	down = ')'
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = input[:len(input)-1] // remove newline

	fmt.Println("part1:", solvePart1(input))
	fmt.Println("part2:", solvePart2(input))
}

func solvePart1(input []byte) int {
	floor := 0
	for _, ch := range input {
		if ch == up {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func solvePart2(input []byte) int {
	floor := 0
	for i, ch := range input {
		if ch == up {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}
