package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	flying = iota
	resting
)

type deer struct {
	name                     string
	speed, flyTime, restTime int
}

func main() {
	deers := readInput()
	time := 2503
	fmt.Println("part1:", solvePart1(deers, time))
	fmt.Println("part2:", solvePart2(deers, time))
}

func readInput() []deer {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var deers []deer
	for s.Scan() {
		var d deer
		var err error
		parts := strings.Split(s.Text(), " ")
		d.name = parts[0]
		d.speed, err = strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal(err)
		}
		d.flyTime, err = strconv.Atoi(parts[6])
		if err != nil {
			log.Fatal(err)
		}
		d.restTime, err = strconv.Atoi(parts[13])
		if err != nil {
			log.Fatal(err)
		}
		deers = append(deers, d)

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return deers
}

func distanceTraveled(d deer, time int) int {
	roundTime := d.flyTime + d.restTime
	roundsCnt := time / roundTime
	distance := roundsCnt * d.flyTime * d.speed
	left := time % roundTime
	if left > d.flyTime {
		left = d.flyTime
	}
	distance += left * d.speed
	return distance
}

func solvePart1(deers []deer, time int) int {
	maxDistanceTraveled := 0
	for _, d := range deers {
		maxDistanceTraveled = max(maxDistanceTraveled, distanceTraveled(d, time))
	}
	return maxDistanceTraveled
}

func solvePart2(deers []deer, time int) int {
	state := make(map[string]int)
	timesLeft := make(map[string]int)
	distances := make(map[string]int)
	points := make(map[string]int)

	for _, d := range deers {
		timesLeft[d.name] = d.flyTime
	}

	for range time {
		for _, d := range deers {
			if state[d.name] == flying {
				distances[d.name] += d.speed
			}

			timesLeft[d.name]--

			if timesLeft[d.name] == 0 {
				if state[d.name] == flying {
					state[d.name] = resting
					timesLeft[d.name] = d.restTime
				} else {
					state[d.name] = flying
					timesLeft[d.name] = d.flyTime
				}
			}
		}

		maxDistance := 0
		for _, distance := range distances {
			maxDistance = max(maxDistance, distance)
		}
		for name, distance := range distances {
			if distance == maxDistance {
				points[name]++
			}
		}
	}

	maxPoints := 0
	for _, p := range points {
		maxPoints = max(maxPoints, p)
	}
	return maxPoints
}
