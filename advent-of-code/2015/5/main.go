package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	strs := readInput()

	var niceCnt1, niceCnt2 int
	for _, s := range strs {
		if isNice1(s) {
			niceCnt1++
		}
		if isNice2(s) {
			niceCnt2++
		}
	}

	fmt.Println("part1:", niceCnt1)
	fmt.Println("part2:", niceCnt2)
}

func readInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	var input []string
	for s.Scan() {
		input = append(input, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func isNice1(s string) bool {
	vowelsCnt := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowelsCnt++
		}
	}
	if vowelsCnt < 3 {
		return false
	}

	twiceInARow := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			twiceInARow = true
			break
		}
	}
	if !twiceInARow {
		return false
	}

	for _, x := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, x) {
			return false
		}
	}

	return true
}

func isNice2(s string) bool {
	twoPairs := false
	for i := 0; i < len(s)-3; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			twoPairs = true
			break
		}
	}
	if !twoPairs {
		return false
	}

	repeat := false
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			repeat = true
			break
		}
	}
	if !repeat {
		return false
	}

	return true
}
