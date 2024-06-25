package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	strs := readInput()
	fmt.Println("part1:", solvePart1(strs))
	fmt.Println("part2:", solvePart2(strs))
}

func readInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var strs []string
	for s.Scan() {
		strs = append(strs, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return strs
}

func solvePart1(strs []string) int {
	lenSum := 0
	for _, str := range strs {
		strLen := len(str)
		str = strings.ReplaceAll(str, `\"`, "x")
		str = strings.ReplaceAll(str, `\\`, "x")
		re := regexp.MustCompile(`\\x\w\w`)
		str = re.ReplaceAllString(str, "x")
		strLen2 := len(str) - 2
		lenSum += strLen - strLen2
	}
	return lenSum
}

func solvePart2(strs []string) int {
	lenSum := 0
	for _, str := range strs {
		strLen1 := len(str)
		strLen2 := len(str) + 2
		for _, c := range str {
			if c == '\\' || c == '"' {
				strLen2++
			}
		}
		lenSum += strLen2 - strLen1
	}
	return lenSum
}
