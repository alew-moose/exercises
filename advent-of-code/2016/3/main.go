package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sizes := readInput()
	fmt.Println("part1:", solvePart1(sizes))
	fmt.Println("part2:", solvePart2(sizes))
}

func readInput() [][3]int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sizes [][3]int
	for {
		var size [3]int
		n, err := fmt.Fscanf(f, "%d %d %d\n", &size[0], &size[1], &size[2])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n != 3 {
			log.Fatal("n != 3")
		}
		sizes = append(sizes, size)
	}
	return sizes
}

func isPossibleTriangle(s [3]int) bool {
	return s[0]+s[1] > s[2] && s[0]+s[2] > s[1] && s[1]+s[2] > s[0]
}

func solvePart1(sizes [][3]int) int {
	possibleTrianglesCnt := 0
	for _, size := range sizes {
		if isPossibleTriangle(size) {
			possibleTrianglesCnt++
		}
	}
	return possibleTrianglesCnt
}

func solvePart2(sizes [][3]int) int {
	possibleTrianglesCnt := 0
	for i := 0; i < len(sizes); i += 3 {
		for k := 0; k < 3; k++ {
			triangle := [3]int{sizes[i][k], sizes[i+1][k], sizes[i+2][k]}
			if isPossibleTriangle(triangle) {
				possibleTrianglesCnt++
			}
		}
	}
	return possibleTrianglesCnt
}
