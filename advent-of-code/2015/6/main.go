package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	turnOn = iota
	turnOff
	toggle
)

type instruction struct {
	action         int
	x1, y1, x2, y2 int
}

func main() {
	instrs := readInput()
	fmt.Println("part1:", solvePart1(instrs))
	fmt.Println("part2:", solvePart2(instrs))
}

func readInput() []instruction {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)

	var instrs []instruction
	for s.Scan() {
		f := strings.Split(s.Text(), " ")
		var action int
		var coord1Str, coord2Str string
		if f[0] == "toggle" {
			action = toggle
			coord1Str = f[1]
			coord2Str = f[3]
		} else if f[0] == "turn" && f[1] == "on" {
			action = turnOn
			coord1Str = f[2]
			coord2Str = f[4]
		} else if f[0] == "turn" && f[1] == "off" {
			action = turnOff
			coord1Str = f[2]
			coord2Str = f[4]
		} else {
			log.Fatal("invalid input")
		}
		coord1 := strings.Split(coord1Str, ",")
		coord2 := strings.Split(coord2Str, ",")
		x1, err := strconv.Atoi(coord1[0])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(coord1[1])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(coord2[0])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(coord2[1])
		if err != nil {
			log.Fatal(err)
		}
		instrs = append(instrs, instruction{
			action: action,
			x1:     x1,
			y1:     y1,
			x2:     x2,
			y2:     y2,
		})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return instrs
}

func solvePart1(instrs []instruction) int {
	var grid [1000][1000]bool
	for _, instr := range instrs {
		for y := instr.y1; y <= instr.y2; y++ {
			for x := instr.x1; x <= instr.x2; x++ {
				switch instr.action {
				case turnOn:
					grid[y][x] = true
				case turnOff:
					grid[y][x] = false
				case toggle:
					grid[y][x] = !grid[y][x]
				}
			}
		}
	}

	onCnt := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if grid[y][x] {
				onCnt++
			}
		}
	}
	return onCnt
}

func solvePart2(instrs []instruction) int {
	var grid [1000][1000]int
	for _, instr := range instrs {
		for y := instr.y1; y <= instr.y2; y++ {
			for x := instr.x1; x <= instr.x2; x++ {
				switch instr.action {
				case turnOn:
					grid[y][x]++
				case turnOff:
					grid[y][x]--
					if grid[y][x] < 0 {
						grid[y][x] = 0
					}
				case toggle:
					grid[y][x] += 2
				}
			}
		}
	}

	sum := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			sum += grid[y][x]
		}
	}
	return sum
}
