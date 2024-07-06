package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	lines := readInput()
	fmt.Println("part1:", solvePart1(lines))
	fmt.Println("part2:", solvePart2(lines))
}

func readInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

type charCount struct {
	char  byte
	count int
}

func solve(lines []string, cmpFunc func(_, _ charCount) int) string {
	answer := make([]byte, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		cnt := make(map[byte]int)
		for l := range lines {
			cnt[lines[l][i]]++
		}
		var charCounts []charCount
		for ch, n := range cnt {
			charCounts = append(charCounts, charCount{char: ch, count: n})
		}
		slices.SortFunc(charCounts, cmpFunc)
		answer[i] = charCounts[0].char
	}
	return string(answer)
}

func solvePart1(lines []string) string {
	return solve(lines, func(c1, c2 charCount) int {
		if c1.count > c2.count {
			return -1
		}
		if c1.count < c2.count {
			return 1
		}
		return cmp.Compare(c1.char, c2.char)
	})
}

func solvePart2(lines []string) string {
	return solve(lines, func(c1, c2 charCount) int {
		if c1.count < c2.count {
			return -1
		}
		if c1.count > c2.count {
			return 1
		}
		return cmp.Compare(c1.char, c2.char)
	})
}
