package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// discs, err := getInput("test-input.txt")
	discs, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 1:", solvePart1(discs))
	fmt.Println("part 2:", solvePart2(discs))
}

func solvePart1(discs []*Disc) int {

	// return -1
	return findOffset(discs)
}

func solvePart2(discs []*Disc) int {
	discs = append(discs, &Disc{pos: 0, period: 11})
	return findOffset(discs)
}

func findOffset(discs []*Disc) int {
	// Должно быть более эффективное решение. Но это работает достаточно быстро
OFFSET:
	for offset := 0; ; offset++ {
		for i, d := range discs {
			discPos := (offset + i + 1 + d.pos) % d.period
			if discPos != 0 {
				continue OFFSET
			}
		}
		return offset
	}
}

type Disc struct {
	pos    int
	period int
}

func getInput(fileName string) ([]*Disc, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var discs []*Disc
	for scanner.Scan() {
		disc, err := parseDisc(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q", scanner.Text())
		}
		discs = append(discs, disc)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return discs, nil
}

var discRE = regexp.MustCompile(`^Disc #\d+ has (\d+) positions; at time=0, it is at position (\d+)\.$`)

func parseDisc(s string) (*Disc, error) {
	matches := discRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil, errors.New("no match")
	}
	len, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse len %q", matches[0][1])
	}
	pos, err := strconv.Atoi(matches[0][2])
	if err != nil {
		return nil, fmt.Errorf("failed to parse pos %q", matches[0][2])
	}
	return &Disc{pos: pos, period: len}, nil
}

/*
Discs:

main.DiscOffset{offset:0, period:5}
main.DiscOffset{offset:1, period:2}

find o:
	d0 offsets: 0, 5, 10, 15, ...
	d1 offsets: 1, 3, 5, 7

	offset = 0 + 5x | x = 1
	offset = 1 + 2y | y = 2

	x, y: int >= 0


	0: (0 + offset) % 5 == 0
	1: (1 + offset) % 2 == 0

	offset = 0
	0 % 5 = 0
	1 % 2 = 1

	offset = 1
	1 % 5 = 1
	2 % 2 = 0

	(0 + x) % 5 == 0
	(1 + x) % 2 == 0

	x % 5 == 0       | x = 0, 5, 10, ... | x = 0 + 5n
	(1 + x) % 2 == 0 | x = 1, 3, 5, ...  | x = 1 + 2k

	0, step 5
	1, step 2

	(offset - 1) % 2 == 0
*/

/*
d0:  main.DiscOffset{offset:11, period:13}
d1:  main.DiscOffset{offset:7, period:19}
d2:  main.DiscOffset{offset:1, period:3}
d3:  main.DiscOffset{offset:2, period:7}
d4:  main.DiscOffset{offset:2, period:5}
d5:  main.DiscOffset{offset:6, period:17}

	(o + 11) % 13 = 0 | 2, 15,      ...
	(o + 7)  % 19 = 0 | 12, 31,     ...
	(o + 1)  %  3 = 0 | 2, 5, 8,    ...
	(o + 2)  %  7 = 0 | 5, 12, 19,  ...
	(o + 2)  %  5 = 0 | 3, 8, 13,   ...
	(o + 6)  % 17 = 0 | 11, 28,     ...

376777

	factor 376777      = 53 7109
	factor 376777 - 11 = 2 13 43 337
	factor 376777 -  7 = 2 3 5 19 661
	factor 376777 -  1 = 2 2 2 3 3 5233
	factor 376777 -  2 = 5 5 7 2153
	factor 376777 -  6 = 17 37 599
*/

/*
type DiscOffset struct {
	offset int
	period int
}

func discsOffsets(discs []*Disc) []DiscOffset {
	var offsets []DiscOffset
	for i, d := range discs {
		for offset := 0; ; offset++ {
			discPos := (offset + i + 1 + d.pos) % d.period
			if discPos == 0 {
				offsets = append(offsets, DiscOffset{offset: offset, period: d.period})
				break
			}

		}
	}
	return offsets
}
*/
