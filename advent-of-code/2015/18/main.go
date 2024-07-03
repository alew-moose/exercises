package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const boardSize = 100

func main() {
	board := readInput()
	fmt.Println("part1:", solvePart1(board))
	fmt.Println("part2:", solvePart2(board))
}

func readInput() [boardSize][boardSize]bool {
	var board [boardSize][boardSize]bool
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	for y := 0; s.Scan(); y++ {
		for x, c := range s.Text() {
			if c == '#' {
				board[y][x] = true
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return board
}

func solvePart1(board [boardSize][boardSize]bool) int {
	for range 100 {
		step(&board)
	}
	return boardOnCnt(&board)
}

func step(board *[boardSize][boardSize]bool) {
	var newBoard [boardSize][boardSize]bool
	for y := range boardSize {
		for x := range boardSize {
			onCnt := neighborOnCnt(board, y, x)
			if board[y][x] {
				if onCnt == 2 || onCnt == 3 {
					newBoard[y][x] = true
				}
			} else {
				if onCnt == 3 {
					newBoard[y][x] = true
				}
			}
		}
	}
	*board = newBoard
}

func neighborOnCnt(board *[boardSize][boardSize]bool, y, x int) int {
	onCnt := 0
	for ny := y - 1; ny <= y+1; ny++ {
		for nx := x - 1; nx <= x+1; nx++ {
			if ny < 0 || ny >= boardSize || nx < 0 || nx >= boardSize || ny == y && nx == x {
				continue
			}
			if board[ny][nx] {
				onCnt++
			}
		}
	}
	return onCnt
}

func boardOnCnt(board *[boardSize][boardSize]bool) int {
	onCnt := 0
	for y := range boardSize {
		for x := range boardSize {
			if board[y][x] {
				onCnt++
			}
		}
	}
	return onCnt
}

func solvePart2(board [boardSize][boardSize]bool) int {
	board[0][0] = true
	board[0][boardSize-1] = true
	board[boardSize-1][0] = true
	board[boardSize-1][boardSize-1] = true
	for range 100 {
		step2(&board)
	}
	return boardOnCnt(&board)
}

func step2(board *[boardSize][boardSize]bool) {
	var newBoard [boardSize][boardSize]bool
	newBoard[0][0] = true
	newBoard[0][boardSize-1] = true
	newBoard[boardSize-1][0] = true
	newBoard[boardSize-1][boardSize-1] = true
	for y := range boardSize {
		for x := range boardSize {
			if y == 0 && x == 0 || y == 0 && x == boardSize-1 || y == boardSize-1 && x == 0 || y == boardSize-1 && x == boardSize-1 {
				continue
			}
			onCnt := neighborOnCnt(board, y, x)
			if board[y][x] {
				if onCnt == 2 || onCnt == 3 {
					newBoard[y][x] = true
				}
			} else {
				if onCnt == 3 {
					newBoard[y][x] = true
				}
			}
		}
	}
	*board = newBoard
}
