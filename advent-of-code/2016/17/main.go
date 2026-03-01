package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	// passcode := []byte("ulqzkmiv") // test passcode
	passcode := []byte("mmsxrhfx")
	fmt.Println("part 1:", solvePart1(passcode))
	fmt.Println("part 2:", solvePart2(passcode))
}

func solvePart1(passcode []byte) string {
	pos := Coord{x: 0, y: 0}
	target := Coord{x: 3, y: 3}
	return shortestPath(passcode, pos, target)
}

func shortestPath(passcode []byte, pos, target Coord) string {
	// scanner := bufio.NewScanner(os.Stdin)
	queue := []*Step{&Step{pos: pos}}
	for len(queue) > 0 {
		step := queue[0]
		queue = queue[1:]
		if step.pos == target {
			return string(step.path)
		}
		doorsState := getDoorsState(passcode, step.path)
		// printMap(step, doorsState)
		// fmt.Printf("%d,%d ", step.pos.y, step.pos.x)
		for i, dir := range Directions {
			if !doorsState[i] {
				continue
			}
			y := step.pos.y + dir.dy
			if y < 0 || y >= MapSize {
				continue
			}
			x := step.pos.x + dir.dx
			if x < 0 || x >= MapSize {
				continue
			}
			path := pathAppend(step.path, dir.dir)
			nextStep := &Step{
				pos:  Coord{y: y, x: x},
				path: path,
			}
			// switch i {
			// case 0:
			// 	fmt.Print("^")
			// case 1:
			// 	fmt.Print("v")
			// case 2:
			// 	fmt.Print("<")
			// case 3:
			// 	fmt.Print(">")
			// default:
			// 	panic("invalid i")
			// }
			queue = append(queue, nextStep)
		}
		// fmt.Println()
		// scanner.Scan()
	}
	return ""
}

func solvePart2(passcode []byte) int {
	pos := Coord{x: 0, y: 0}
	target := Coord{x: 3, y: 3}
	return longestPath(passcode, pos, target)
}

func longestPath(passcode []byte, pos, target Coord) int {
	// scanner := bufio.NewScanner(os.Stdin)
	queue := []*Step{&Step{pos: pos}}
	lp := -1
	for len(queue) > 0 {
		step := queue[0]
		queue = queue[1:]
		if step.pos == target {
			if len(step.path) > lp {
				lp = len(step.path)
			}
			continue
		}
		doorsState := getDoorsState(passcode, step.path)
		// printMap(step, doorsState)
		// fmt.Printf("%d,%d ", step.pos.y, step.pos.x)
		for i, dir := range Directions {
			if !doorsState[i] {
				continue
			}
			y := step.pos.y + dir.dy
			if y < 0 || y >= MapSize {
				continue
			}
			x := step.pos.x + dir.dx
			if x < 0 || x >= MapSize {
				continue
			}
			path := pathAppend(step.path, dir.dir)
			nextStep := &Step{
				pos:  Coord{y: y, x: x},
				path: path,
			}
			// switch i {
			// case 0:
			// 	fmt.Print("^")
			// case 1:
			// 	fmt.Print("v")
			// case 2:
			// 	fmt.Print("<")
			// case 3:
			// 	fmt.Print(">")
			// default:
			// 	panic("invalid i")
			// }
			queue = append(queue, nextStep)
		}
		// fmt.Println()
		// scanner.Scan()
	}
	return lp
}

func pathAppend(path []Dir, dir Dir) []Dir {
	nextPath := make([]Dir, len(path)+1)
	copy(nextPath, path)
	nextPath[len(nextPath)-1] = dir
	return nextPath
}

const MapSize = 4

type Coord struct {
	x, y int
}

type Step struct {
	pos  Coord
	path []Dir
}

type Direction struct {
	dir Dir
	dx  int
	dy  int
}

var Directions = [...]Direction{
	{dir: DirUp, dy: -1},
	{dir: DirDown, dy: +1},
	{dir: DirLeft, dx: -1},
	{dir: DirRight, dx: +1},
}

type Dir byte

const (
	DirUp    Dir = 'U'
	DirDown  Dir = 'D'
	DirLeft  Dir = 'L'
	DirRight Dir = 'R'
)

// [up, down, left, right]; false=closed, true=open
func getDoorsState(passcode []byte, path []Dir) [4]bool {
	data := make([]byte, 0, len(passcode)+len(path))
	data = append(data, passcode...)
	for _, dir := range path {
		data = append(data, byte(dir))
	}
	hash := fmt.Sprintf("%x", md5.Sum(data))
	var state [4]bool
	for i := range 4 {
		if hash[i] >= 'b' && hash[i] <= 'f' {
			state[i] = true
		}
	}
	return state
}

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3
)

func printMap(step *Step, doorsState [4]bool) {
	path := make([]byte, len(step.path))
	for i := range path {
		path[i] = byte(step.path[i])
	}
	fmt.Printf("%-3d %s\n", len(step.path), step.path)
	fmt.Println("#########")
	for y := range 4 {
		fmt.Print("#")
		for x := range 4 {
			if step.pos.y == y && step.pos.x == x {
				fmt.Print("o")
			} else {
				fmt.Print(" ")
			}
			if x < 3 {
				if step.pos.y == y && step.pos.x == x {
					if !doorsState[right] {
						fmt.Print("|")
					} else {
						fmt.Print(" ")
					}
				} else if step.pos.y == y && step.pos.x == x+1 {
					if !doorsState[left] {
						fmt.Print("|")
					} else {
						fmt.Print(" ")
					}

				} else {
					fmt.Print("|")
				}
			}
		}
		fmt.Println("#")
		if y < 3 {
			fmt.Print("#")
			for x := range 4 {
				if step.pos.y == y+1 && step.pos.x == x {
					if !doorsState[up] {
						fmt.Print("-")
					} else {
						fmt.Print(" ")
					}
				} else if step.pos.y == y && step.pos.x == x {
					if !doorsState[down] {
						fmt.Print("-")
					} else {
						fmt.Print(" ")
					}
				} else {
					fmt.Print("-")
				}
				if x < 3 {
					fmt.Print("#")
				}
			}
			fmt.Println("#")
		}
	}
	fmt.Println("#########")
}
