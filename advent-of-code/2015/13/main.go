package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	points := readInput()
	fmt.Println("part1:", solve(points))
	points["self"] = make(map[string]int)
	fmt.Println("part2:", solve(points))
}

func solve(points map[string]map[string]int) int {
	names := make([]string, 0, len(points))
	for name := range points {
		names = append(names, name)
	}

	havePoints := false
	maxPoints := 0
	permutations(names, func() {
		p := countPoints(names, points)
		if !havePoints || p > maxPoints {
			havePoints = true
			maxPoints = p
		}
	})

	return maxPoints
}

func readInput() map[string]map[string]int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	input := make(map[string]map[string]int)
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		name1, name2 := parts[0], parts[10]
		name2 = name2[:len(name2)-1]
		points, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal(err)
		}
		if parts[2] == "lose" {
			points = -points
		}
		if _, ok := input[name1]; !ok {
			input[name1] = make(map[string]int)
		}
		input[name1][name2] = points

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func permutations(items []string, cb func()) {
	if len(items) == 0 {
		cb()
		return
	}
	for i := range items {
		items[0], items[i] = items[i], items[0]
		permutations(items[1:], cb)
	}
}

func countPoints(names []string, points map[string]map[string]int) int {
	p := 0
	for i, name := range names {
		l := i - 1
		if l < 0 {
			l = len(names) - 1
		}
		r := i + 1
		if r >= len(names) {
			r = 0
		}

		p += points[name][names[l]]
		p += points[name][names[r]]
	}
	return p
}
