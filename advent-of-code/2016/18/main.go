package main

import "fmt"

func main() {
	firstRow := ".^^^^^.^^^..^^^^^...^.^..^^^.^^....^.^...^^^...^^^^..^...^...^^.^.^.......^..^^...^.^.^^..^^^^^...^."
	fmt.Println("part 1:", solvePart1(firstRow))
	fmt.Println("part 2:", solvePart2(firstRow))
}

const (
	TileSafe = '.'
	TileTrap = '^'
)

func solvePart1(firstRow string) int {
	rowsCnt := 40
	m := makeMap(firstRow, rowsCnt)
	return countTiles(m, TileSafe)
}

func solvePart2(firstRow string) int {
	rowsCnt := 400000
	m := makeMap(firstRow, rowsCnt)
	return countTiles(m, TileSafe)
}

func makeMap(firstRow string, rowsCnt int) [][]byte {
	m := make([][]byte, rowsCnt)
	m[0] = []byte(firstRow)
	for y := 1; y < rowsCnt; y++ {
		m[y] = make([]byte, len(firstRow))
		for x := 0; x < len(firstRow); x++ {
			if isTrap(m[y-1], x) {
				m[y][x] = TileTrap
			} else {
				m[y][x] = TileSafe
			}
		}
	}
	return m
}

func isTrap(prevRow []byte, x int) bool {
	var leftTrap, centerTrap, rightTrap bool
	if x-1 >= 0 {
		leftTrap = prevRow[x-1] == TileTrap
	}
	if x+1 < len(prevRow) {
		rightTrap = prevRow[x+1] == TileTrap
	}
	centerTrap = prevRow[x] == TileTrap

	if leftTrap && centerTrap && !rightTrap {
		return true
	}
	if rightTrap && centerTrap && !leftTrap {
		return true
	}
	if leftTrap && !centerTrap && !rightTrap {
		return true
	}
	if rightTrap && !centerTrap && !leftTrap {
		return true
	}
	return false
}

func countTiles(m [][]byte, tile byte) int {
	cnt := 0
	for y := range len(m) {
		for x := range len(m[y]) {
			if m[y][x] == tile {
				cnt++
			}
		}
	}
	return cnt
}
