package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	containers := readInput()
	fmt.Println("part1:", combinations(containers, 150))
	fmt.Println("part2:", solvePart2(containers, 150))
}

func readInput() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var containers []int
	for s.Scan() {
		str := s.Text()
		c, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		containers = append(containers, c)

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return containers
}

func combinations(containers []int, toFill int) int {
	if toFill == 0 {
		return 1
	}
	if len(containers) == 0 || toFill < 0 {
		return 0
	}
	return combinations(containers[1:], toFill-containers[0]) + combinations(containers[1:], toFill)
}

func combinations2(containers []int, toFill int, used int) (int, int) {
	if toFill == 0 {
		return 1, used
	}
	if len(containers) == 0 || toFill < 0 {
		return 0, 0
	}
	count1, used1 := combinations2(containers[1:], toFill-containers[0], used+1)
	count2, used2 := combinations2(containers[1:], toFill, used)
	if used1 != 0 {
		if used2 != 0 {
			if used1 < used2 {
				return count1, used1
			} else if used2 < used1 {
				return count2, used2
			} else {
				return count1 + count2, used1
			}
		}
		return count1, used1
	}
	if used2 != 0 {
		return count2, used2
	}
	return 0, 0
}

func solvePart2(containers []int, toFill int) int {
	count, _ := combinations2(containers, toFill, 0)
	return count
}
