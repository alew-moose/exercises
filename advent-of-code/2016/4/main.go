package main

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type room struct {
	name     string
	sectorID int
	checksum string
}

func main() {
	rooms := readInput()
	fmt.Println("part1:", solvePart1(rooms))
	fmt.Println("part2:", solvePart2(rooms))
}

func readInput() []room {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var rooms []room
	for s.Scan() {
		r, err := parseRoom(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, r)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return rooms
}

var roomRe = regexp.MustCompile(`^([a-z-]+)-(\d+)\[([a-z]+)\]$`)

func parseRoom(s string) (room, error) {
	m := roomRe.FindStringSubmatch(s)
	if m == nil {
		return room{}, errors.New("no match")
	}
	sectorID, err := strconv.Atoi(m[2])
	if err != nil {
		return room{}, err
	}
	r := room{
		name:     m[1],
		sectorID: sectorID,
		checksum: m[3],
	}
	return r, nil
}

func checksum(s string) string {
	type charCount struct {
		char  byte
		count int
	}
	count := make(map[byte]int)
	for _, c := range s {
		if c == '-' {
			continue
		}
		count[byte(c)]++
	}
	var charCounts []charCount
	for ch, cnt := range count {
		charCounts = append(charCounts, charCount{char: ch, count: cnt})
	}
	slices.SortFunc(charCounts, func(a, b charCount) int {
		if a.count > b.count {
			return -1
		}
		if a.count < b.count {
			return 1
		}
		return cmp.Compare(a.char, b.char)
	})
	checksumBytes := make([]byte, 0, len(charCounts))
	for _, cc := range charCounts[:5] {
		checksumBytes = append(checksumBytes, cc.char)
	}
	return string(checksumBytes)
}

func solvePart1(rooms []room) int {
	sectorIDsSum := 0
	for _, r := range rooms {
		if r.checksum == checksum(r.name) {
			sectorIDsSum += r.sectorID
		}
	}
	return sectorIDsSum
}

func solvePart2(rooms []room) int {
	for _, r := range rooms {
		if decrypt(r.name, r.sectorID) == "northpole object storage" {
			return r.sectorID
		}
	}
	return 0
}

func decrypt(s string, n int) string {
	b := make([]byte, len(s))
	for i, c := range s {
		if c == '-' {
			b[i] = ' '
			continue
		}

		b[i] = byte((int(c)-'a'+n)%26 + 'a')
	}
	return string(b)
}
