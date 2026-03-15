package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	disks, err := getInput("test-input.txt")
	// disks, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 1:", solvePart1(disks))
	fmt.Println("part 2:", solvePart2(disks))
}

func solvePart1(ds []*Disk) int {
	disks := make([]*Disk, len(ds))
	copy(disks, ds)
	slices.SortFunc(disks, func(d1, d2 *Disk) int {
		return d2.avail - d1.avail
	})
	pairsCnt := 0
	for i := range disks {
		if disks[i].used == 0 {
			continue
		}
		for k := range disks {
			if k == i {
				continue
			}
			if disks[i].used <= disks[k].avail {
				pairsCnt++
			} else {
				break
			}
		}
	}
	return pairsCnt
}

func solvePart2(disks []*Disk) int {
	grid := makeGrid(disks)
	printGrid(os.Stdout, grid)
	return -1
}

type Node struct {
	used, size int
	target     bool
}

func makeGrid(disks []*Disk) [][]*Node {
	var maxX, maxY int
	for _, d := range disks {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}
	grid := make([][]*Node, maxY+1)
	for y := range maxY + 1 {
		grid[y] = make([]*Node, maxX+1)
	}
	for _, d := range disks {
		grid[d.y][d.x] = &Node{
			used:   d.used,
			size:   d.size,
			target: d.y == 0 && d.x == maxX,
		}
	}
	return grid
}

func printGrid(w io.Writer, grid [][]*Node) {
	for y := range grid {
		for x := range grid[y] {
			w.Write([]byte(fmt.Sprintf("%3d/%-3d", grid[y][x].used, grid[y][x].size)))
			if grid[y][x].target {
				w.Write([]byte("! "))
			} else {
				w.Write([]byte("  "))
			}
		}
		w.Write([]byte("\n"))
	}
}

func gridToString(grid [][]*Node) string {
	var b strings.Builder
	printGrid(&b, grid)
	return b.String()
}

type Disk struct {
	x, y, size, used, avail, usePct int
}

var diskRE = regexp.MustCompile(
	`^/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%$`,
)

func parseDisk(s string) *Disk {
	matches := diskRE.FindAllStringSubmatch(s, -1)
	var disk Disk
	if matches == nil {
		return nil
	}
	for i, ptr := range []*int{&disk.x, &disk.y, &disk.size, &disk.used, &disk.avail, &disk.usePct} {
		var err error
		*ptr, err = strconv.Atoi(matches[0][i+1])
		if err != nil {
			panic(err)
		}
	}
	return &disk
}

func getInput(file string) ([]*Disk, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var disks []*Disk
	for scanner.Scan() {
		s := scanner.Text()
		disk := parseDisk(s)
		if disk == nil {
			return nil, fmt.Errorf("failed to parse %q", s)
		}
		disks = append(disks, disk)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return disks, err
}
