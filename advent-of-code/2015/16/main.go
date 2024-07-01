package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type sue struct {
	num        int
	properties map[string]int
}

func main() {
	sues := readInput()
	propertiesRequired := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	fmt.Println("part1:", solvePart1(sues, propertiesRequired))
	fmt.Println("part2:", solvePart2(sues, propertiesRequired))
}

func readInput() []sue {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var ss []sue
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s, err := parseSue(sc.Text())
		if err != nil {
			log.Fatal(err)
		}
		ss = append(ss, s)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return ss
}

var propertyRe = regexp.MustCompile(`(\w+): (\d+)`)

func parseSue(str string) (sue, error) {
	var s sue
	s.properties = make(map[string]int)

	parts := strings.SplitN(str, ":", 2)
	if len(parts) != 2 {
		return sue{}, errors.New("failed to parse")
	}
	nParts := strings.Split(parts[0], " ")
	if len(nParts) != 2 {
		return sue{}, errors.New("failed to parse n part")
	}
	num, err := strconv.Atoi(nParts[1])
	if err != nil {
		return sue{}, err
	}
	s.num = num

	matches := propertyRe.FindAllStringSubmatch(parts[1], -1)
	if len(matches) == 0 {
		return sue{}, errors.New("failed to parse properties")
	}
	for _, match := range matches {
		key, valStr := match[1], match[2]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return sue{}, err
		}
		s.properties[key] = val
	}

	return s, nil
}

func solvePart1(sues []sue, propertiesRequired map[string]int) int {
SUE:
	for _, s := range sues {
		for prKey, prVal := range propertiesRequired {
			val, ok := s.properties[prKey]
			if !ok {
				continue
			}
			if val != prVal {
				continue SUE
			}
		}
		return s.num
	}
	return 0
}

func solvePart2(sues []sue, propertiesRequired map[string]int) int {
SUE:
	for _, s := range sues {
		for prKey, prVal := range propertiesRequired {
			val, ok := s.properties[prKey]
			if !ok {
				continue
			}
			switch prKey {
			case "cats", "trees":
				if val <= prVal {
					continue SUE
				}
			case "pomeranians", "goldfish":
				if val >= prVal {
					continue SUE
				}
			default:
				if val != prVal {
					continue SUE
				}
			}
		}
		return s.num
	}
	return 0
}
