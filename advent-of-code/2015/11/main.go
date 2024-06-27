package main

import (
	"fmt"
)

func main() {
	const input = "cqjxjnds"

	answer1 := solve(input)
	answer2 := solve(answer1)
	fmt.Println("part1:", answer1)
	fmt.Println("part2:", answer2)
}

func solve(input string) string {
	s := []byte(input)
	s = increment(s)
	for !isValid(s) {
		s = increment(s)
	}
	return string(s)
}

func increment(s []byte) []byte {
	carry := true
	for i := len(s) - 1; i >= 0 && carry; i-- {
		s[i]++
		if s[i] > 'z' {
			s[i] = 'a'
		} else {
			carry = false
		}
	}
	// don't really need this, the password must be always 8 chars long and shouldn't overflow
	if carry {
		s = append([]byte{'a'}, s...)
	}
	return s
}

func isValid(s []byte) bool {
	hasIncreasing := false
	for i := range len(s) - 2 {
		if s[i] == s[i+1]-1 && s[i+1] == s[i+2]-1 {
			hasIncreasing = true
			break
		}
	}
	if !hasIncreasing {
		return false
	}

	for _, c := range s {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
	}

	pairsCnt := 0
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			pairsCnt++
			i++
			if pairsCnt >= 2 {
				break
			}
		}
		i++
	}
	if pairsCnt < 2 {
		return false
	}

	return true
}
