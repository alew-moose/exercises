package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b[:len(b)-1])

	fmt.Println("part1:", decompressedLength1(str))
	fmt.Println("part2:", decompressedLength2(str))
}

func decompressedLength1(s string) int {
	length := 0
	for len(s) > 0 {
		if o := strings.IndexByte(s, '('); o != -1 {
			length += o
			c := strings.IndexByte(s, ')')
			if c == -1 {
				panic("')' not found")
			}
			marker := s[o+1 : c]
			numStrs := strings.Split(marker, "x")
			if len(numStrs) != 2 {
				panic("invalid marker")
			}
			markerLength, err := strconv.Atoi(numStrs[0])
			if err != nil {
				panic(err)
			}
			markerRepeat, err := strconv.Atoi(numStrs[1])
			if err != nil {
				panic(err)
			}
			length += markerLength * markerRepeat
			s = s[o+len(marker)+2+markerLength:]
		} else {
			length += len(s)
			break
		}
	}
	return length
}

func decompressedLength2(s string) int {
	length := 0
	for len(s) > 0 {
		if s[0] == '(' {
			c := strings.IndexByte(s, ')')
			if c == -1 {
				panic("')' not found")
			}
			marker := s[1:c]
			numStrs := strings.Split(marker, "x")
			if len(numStrs) != 2 {
				panic("invalid marker")
			}
			markerLength, err := strconv.Atoi(numStrs[0])
			if err != nil {
				panic(err)
			}
			markerRepeat, err := strconv.Atoi(numStrs[1])
			if err != nil {
				panic(err)
			}

			length += decompressedLength2(s[len(marker)+2:len(marker)+2+markerLength]) * markerRepeat
			s = s[len(marker)+2+markerLength:]
		} else if o := strings.IndexByte(s, '('); o != -1 {
			length += o
			s = s[o:]
		} else {
			length += len(s)
			break
		}
	}
	return length
}
