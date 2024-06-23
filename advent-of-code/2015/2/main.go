package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type box struct {
	h, w, l int // height, width, length
}

func (b *box) surfaceArea() int {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

func (b *box) smallestSideArea() int {
	return min(b.l*b.w, b.w*b.h, b.h*b.l)
}

func (b *box) smallestPerimeter() int {
	return min(2*(b.l+b.w), 2*(b.w+b.h), 2*(b.h+b.l))
}

func (b *box) volume() int {
	return b.h * b.w * b.l
}

func main() {
	boxes := readInput()
	fmt.Println("part1:", solvePart1(boxes))
	fmt.Println("part2:", solvePart2(boxes))
}

func readInput() []box {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	var boxes []box
	for s.Scan() {
		sizes := strings.Split(s.Text(), "x")
		if len(sizes) != 3 {
			log.Fatal("invalid sizes")
		}
		h, err := strconv.Atoi(sizes[0])
		if err != nil {
			log.Fatal(err)
		}
		w, err := strconv.Atoi(sizes[1])
		if err != nil {
			log.Fatal(err)
		}
		l, err := strconv.Atoi(sizes[2])
		if err != nil {
			log.Fatal(err)
		}
		boxes = append(boxes, box{h: h, w: w, l: l})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return boxes
}

func solvePart1(boxes []box) int {
	area := 0
	for _, b := range boxes {
		area += b.surfaceArea() + b.smallestSideArea()
	}
	return area
}

func solvePart2(boxes []box) int {
	length := 0
	for _, b := range boxes {
		length += b.smallestPerimeter() + b.volume()
	}
	return length
}
