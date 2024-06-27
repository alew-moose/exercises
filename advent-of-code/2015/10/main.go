package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const input = "1321131112"
	fmt.Println("part1:", solvePart1(input))
	fmt.Println("part2:", solvePart2(input))
}

func solvePart1(s string) int {
	for range 40 {
		s = lookAndSay(s)
	}
	return len(s)
}

func solvePart2(s string) int {
	for range 50 {
		s = lookAndSay(s)
	}
	return len(s)
}

func lookAndSay(s string) string {
	var nums []int
	prev := -1
	cnt := 0
	for _, c := range s {
		n := int(c - '0')
		if prev == -1 {
			prev = n
			cnt = 1
			continue
		}
		if n != prev {
			nums = append(nums, cnt, prev)
			prev = n
			cnt = 1
			continue
		}
		cnt++
	}
	nums = append(nums, cnt, prev)

	strs := make([]string, 0, len(nums))
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	return strings.Join(strs, "")
}
