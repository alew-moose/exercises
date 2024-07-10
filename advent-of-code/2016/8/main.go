package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type moveType int

const (
	moveRect moveType = iota
	moveRotateCol
	moveRotateRow
)

type move struct {
	moveType moveType
	a, b     int
}

func main() {
	moves := readInput()
	var screen [6][50]bool
	for _, m := range moves {
		makeMove(&screen, m)
	}
	litCnt := 0
	for y := range 6 {
		for x := range 50 {
			if screen[y][x] {
				litCnt++
			}
		}
	}
	fmt.Println("part1:", litCnt)
	printScreen(screen)
}

func printScreen(screen [6][50]bool) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			if screen[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func makeMove(screen *[6][50]bool, m move) {
	switch m.moveType {
	case moveRect:
		for y := 0; y < m.b; y++ {
			for x := 0; x < m.a; x++ {
				screen[y][x] = true
			}
		}
	case moveRotateRow:
		var r [50]bool
		for i := range 50 {
			r[i] = screen[m.a][(50-m.b+i)%50]
		}
		screen[m.a] = r
	case moveRotateCol:
		var r [6]bool
		for i := range 6 {
			r[i] = screen[(6-m.b+i)%6][m.a]
		}
		for i := range 6 {
			screen[i][m.a] = r[i]
		}
	}

}

func readInput() []move {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var moves []move
	for s.Scan() {
		m, err := parseMove(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		moves = append(moves, m)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return moves
}

func parseMove(s string) (move, error) {
	fields := strings.Fields(s)
	if fields[0] == "rect" {
		parts := strings.Split(fields[1], "x")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return move{}, err
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return move{}, err
		}
		return move{moveType: moveRect, a: x, b: y}, nil

	}
	parts := strings.Split(fields[2], "=")
	a, err := strconv.Atoi(parts[1])
	if err != nil {
		return move{}, err
	}
	b, err := strconv.Atoi(fields[4])
	if err != nil {
		return move{}, err
	}
	m := move{a: a, b: b}
	if fields[1] == "column" {
		m.moveType = moveRotateCol
	} else {
		m.moveType = moveRotateRow
	}
	return m, nil
}
