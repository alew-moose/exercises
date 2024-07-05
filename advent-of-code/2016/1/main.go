package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

type turn int
type direction int

const (
	leftTurn turn = iota
	rightTurn
)

const (
	up direction = iota
	right
	down
	left
)

type move struct {
	turn turn
	dist int
}

type coord struct {
	x, y int
}

func distance(c coord) int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func makeTurn(dir direction, t turn) direction {
	if t == leftTurn {
		dir--
	} else {
		dir++
	}
	if dir > left {
		return up
	}
	if dir < up {
		return left
	}
	return dir
}

func makeMove(pos coord, dir direction, m move) (coord, direction) {
	dir = makeTurn(dir, m.turn)
	switch dir {
	case up:
		pos.y -= m.dist
	case right:
		pos.x += m.dist
	case down:
		pos.y += m.dist
	case left:
		pos.x -= m.dist
	}
	return pos, dir
}

func main() {
	moves := readInput()
	fmt.Println("part1:", solvePart1(moves))
	fmt.Println("part2:", solvePart2(moves))
}

func readInput() []move {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b = b[:len(b)-1] // trim newline
	ds := bytes.Split(b, []byte(", "))
	var moves []move
	for _, d := range ds {
		var m move
		if d[0] == 'L' {
			m.turn = leftTurn
		} else {
			m.turn = rightTurn
		}
		dist, err := strconv.Atoi(string(d[1:]))
		if err != nil {
			log.Fatal(err)
		}
		m.dist = dist
		moves = append(moves, m)
	}
	return moves
}

func solvePart1(moves []move) int {
	var pos coord
	dir := up
	for _, m := range moves {
		pos, dir = makeMove(pos, dir, m)
	}
	return distance(pos)
}

func steps(s, e int, cb func(int)) {
	if s < e {
		for x := s; x <= e; x++ {
			cb(x)
		}
	} else {
		for x := s; x >= e; x-- {
			cb(x)
		}
	}
}

func coordSteps(cs, ce coord, cb func(coord)) {
	steps(cs.x, ce.x, func(x int) {
		steps(cs.y, ce.y, func(y int) {
			cb(coord{x: x, y: y})
		})
	})
}

func solvePart2(moves []move) (dist int) {
	var pos coord
	dir := up
	visited := make(map[coord]struct{})
	visited[pos] = struct{}{}

	defer func() {
		r := recover()
		if d, ok := r.(int); ok {
			dist = d
		}
	}()

	for _, m := range moves {
		newPos, newDir := makeMove(pos, dir, m)
		first := true
		coordSteps(pos, newPos, func(p coord) {
			if first {
				first = false
				return
			}
			if _, ok := visited[p]; ok {
				panic(distance(p))
			}
			visited[p] = struct{}{}
		})
		pos = newPos
		dir = newDir
	}

	return 0
}
