package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lines := readInput()
	fmt.Println("part1:", solvePart1(lines))
	fmt.Println("part2:", solvePart2(lines))
}

func readInput() [][]byte {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	var lines [][]byte
	for {
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines
}

func solvePart1(lines [][]byte) string {
	keypad := [][]byte{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}
	x, y := 1, 1
	var code []byte
	for _, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				y--
				if y < 0 {
					y = 0
				}
			case 'D':
				y++
				if y > 2 {
					y = 2
				}
			case 'L':
				x--
				if x < 0 {
					x = 0
				}
			case 'R':
				x++
				if x > 2 {
					x = 2
				}
			}
		}
		code = append(code, keypad[y][x])
	}
	return string(code)
}

func solvePart2(lines [][]byte) string {
	keypad := [][]byte{
		{0, 0, '1', 0, 0},
		{0, '2', '3', '4', 0},
		{'5', '6', '7', '8', '9'},
		{0, 'A', 'B', 'C', 0},
		{0, 0, 'D', 0, 0},
	}
	x, y := 0, 2
	var code []byte
	for _, line := range lines {
		for _, c := range line {
			switch c {
			case 'U':
				if y > 0 && keypad[y-1][x] != 0 {
					y--
				}
			case 'D':
				if y < 4 && keypad[y+1][x] != 0 {
					y++
				}
			case 'L':
				if x > 0 && keypad[y][x-1] != 0 {
					x--
				}
			case 'R':
				if x < 4 && keypad[y][x+1] != 0 {
					x++
				}
			}
		}
		code = append(code, keypad[y][x])
	}
	return string(code)
}
