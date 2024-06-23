package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
)

type coord struct {
	y, x int
}

func (c *coord) move(dir byte) {
	switch dir {
	case up:
		c.y++
	case down:
		c.y--
	case left:
		c.x--
	case right:
		c.x++
	default:
		panic(fmt.Sprintf("invalid dir: %q", dir))
	}

}

func main() {
	dirs := readInput()
	fmt.Println("part1:", solvePart1(dirs))
	fmt.Println("part2:", solvePart2(dirs))
}

func readInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return input[:len(input)-1]
}

func solvePart1(dirs []byte) int {
	var pos coord
	visited := make(map[coord]bool)
	visited[pos] = true

	for _, dir := range dirs {
		pos.move(dir)
		visited[pos] = true
	}

	visitedCnt := 0
	for range visited {
		visitedCnt++
	}
	return visitedCnt
}

func solvePart2(dirs []byte) int {
	var positions [2]coord
	toMove := 0
	visited := make(map[coord]bool)
	visited[coord{0, 0}] = true

	for _, dir := range dirs {
		positions[toMove].move(dir)
		visited[positions[toMove]] = true
		if toMove == 0 {
			toMove = 1
		} else {
			toMove = 0
		}
	}

	visitedCnt := 0
	for range visited {
		visitedCnt++
	}
	return visitedCnt
}
