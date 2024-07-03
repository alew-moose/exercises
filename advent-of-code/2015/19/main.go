package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type replacement struct {
	from, to string
}

type queueElem struct {
	step int
	from string
}

func main() {
	replacements, molecule := readInput()
	fmt.Println("part1:", solvePart1(molecule, replacements))
	// fmt.Println("part2:", solvePart2(molecule, replacements))
}

func readInput() ([]replacement, string) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var replacements []replacement
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		parts := strings.Fields(s.Text())
		if len(parts) != 3 {
			log.Fatal("len parts != 3")
		}
		replacements = append(replacements, replacement{from: parts[0], to: parts[2]})
	}
	if !s.Scan() {
		log.Fatal("line expected")
	}
	return replacements, s.Text()
}

func replaceAll(s string, replacements []replacement, cb func(string)) {
	for _, r := range replacements {
		replaceStr(s, r, cb)
	}
}

func replaceStr(s string, r replacement, cb func(string)) {
	from := 0
	for {
		i := strings.Index(s[from:], r.from)
		if i == -1 {
			break
		}
		newS := s[:i+from] + r.to + s[i+from+len(r.from):]
		cb(newS)
		from += i + len(r.from)
	}
}

func solvePart1(start string, replacements []replacement) int {
	results := make(map[string]struct{})
	replaceAll(start, replacements, func(s string) {
		results[s] = struct{}{}
	})
	cnt := 0
	for range results {
		cnt++
	}
	return cnt
}

// func solvePart2(dest string, replacements []replacement) (steps int) {
// 	seen := make(map[string]struct{})
// 	queue := []queueElem{{step: 1, from: "e"}}
// 	defer func() {
// 		s := recover()
// 		steps = s.(int)
// 	}()
// 	for len(queue) > 0 {
// 		e := queue[0]
// 		queue = queue[1:]
// 		replaceAll(e.from, replacements, func(s string) {
// 			if s == dest {
// 				panic(e.step)
// 			}
// 			if _, ok := seen[s]; ok {
// 				return
// 			}
// 			queue = append(queue, queueElem{step: e.step + 1, from: s})
// 			seen[s] = struct{}{}
// 		})
// 	}

// 	return steps
// }

// func solvePart2(dest string, replacements []replacement) (steps int) {
// 	queue := []queueElem{{step: 1, from: dest}}
// 	seen := make(map[string]struct{})
// 	defer func() {
// 		s := recover()
// 		steps = s.(int)
// 	}()
// 	for i := range replacements {
// 		r := &replacements[i]
// 		r.from, r.to = r.to, r.from
// 	}
// 	for len(queue) > 0 {
// 		e := queue[0]
// 		queue = queue[1:]
// 		replaceAll(e.from, replacements, func(s string) {
// 			if s == "e" {
// 				panic(e.step)
// 			}
// 			if _, ok := seen[s]; ok {
// 				return
// 			}
// 			queue = append(queue, queueElem{step: e.step + 1, from: s})
// 			seen[s] = struct{}{}
// 		})
// 	}
// 	return steps
// }
