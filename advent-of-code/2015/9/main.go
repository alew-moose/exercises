package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	distance := make(map[string]map[string]int)
	for s.Scan() {
		str := s.Text()
		parts := strings.Split(str, " ")
		from, to, distStr := parts[0], parts[2], parts[4]
		dist, err := strconv.Atoi(distStr)
		if err != nil {
			log.Fatal(err)
		}
		for _, l := range []string{from, to} {
			if _, ok := distance[l]; !ok {
				distance[l] = make(map[string]int)
			}
		}
		distance[from][to] = dist
		distance[to][from] = dist
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", solvePart1(distance))
	fmt.Println("part2:", solvePart2(distance))
}

func solvePart1(distance map[string]map[string]int) int {
	minDist := -1
	for from := range distance {
		d := shortestDistance(distance, make(map[string]bool), from)
		if minDist == -1 || d < minDist {
			minDist = d
		}
	}
	return minDist
}

func shortestDistance(distance map[string]map[string]int, visited map[string]bool, from string) int {
	minDist := -1
	visited[from] = true
	for to := range distance[from] {
		if visited[to] {
			continue
		}
		d := distance[from][to] + shortestDistance(distance, visited, to)
		if minDist == -1 || d < minDist {
			minDist = d
		}
	}
	visited[from] = false
	if minDist == -1 {
		return 0
	}
	return minDist
}

func solvePart2(distance map[string]map[string]int) int {
	maxDist := -1
	for from := range distance {
		d := longestDistance(distance, make(map[string]bool), from)
		if maxDist == -1 || d > maxDist {
			maxDist = d
		}
	}
	return maxDist
}

func longestDistance(distance map[string]map[string]int, visited map[string]bool, from string) int {
	maxDist := -1
	visited[from] = true
	for to := range distance[from] {
		if visited[to] {
			continue
		}
		d := distance[from][to] + longestDistance(distance, visited, to)
		if maxDist == -1 || d > maxDist {
			maxDist = d
		}
	}
	visited[from] = false
	if maxDist == -1 {
		return 0
	}
	return maxDist
}
