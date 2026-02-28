package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
)

const FloorsCnt = 4

type State struct {
	generators  [FloorsCnt]uint8
	microchips  [FloorsCnt]uint8
	elevatorPos uint8
	elementsCnt uint8
}

func main() {
	// state, shifts, err := getInput("test-input.txt")
	state, shifts, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	_ = shifts

	fmt.Println("part 1:", solvePart1(state))
	fmt.Println("part 2:", solvePart2(state))
}

func printState(s State) {
	fmt.Println("generators:")
	for _, floor := range s.generators {
		fmt.Printf("%08b\n", floor)
	}
	fmt.Println("microchips:")
	for _, floor := range s.microchips {
		fmt.Printf("%08b\n", floor)
	}
	fmt.Println()
}

func solvePart1(state State) int {
	targetState := makeTargetState(state.elementsCnt)
	return minSteps(state, targetState)
}

func solvePart2(state State) int {
	for i := range 2 {
		n := uint8(1 << (state.elementsCnt + uint8(i)))
		state.generators[0] |= n
		state.microchips[0] |= n
	}
	state.elementsCnt += uint8(2)
	targetState := makeTargetState(state.elementsCnt)
	return minSteps(state, targetState)
}

type Step struct {
	step  int
	state State
}

func minSteps(state, targetState State) int {
	seen := make(map[State]struct{})
	stepsQueue := []Step{{step: 0, state: state}}
	for len(stepsQueue) > 0 {
		nextStep := stepsQueue[0]
		stepsQueue = stepsQueue[1:]
		if _, ok := seen[nextStep.state]; !ok {
			seen[nextStep.state] = struct{}{}
			if nextStep.state == targetState {
				return nextStep.step
			} else {
				stepsQueue = append(stepsQueue, possibleSteps(nextStep.state, nextStep.step+1)...)
			}
		}
	}
	return -1
}

func possibleSteps(state State, step int) []Step {
	nextElevatorPositions := make([]uint8, 0, 2)
	if state.elevatorPos > 0 {
		nextElevatorPositions = append(nextElevatorPositions, state.elevatorPos-1)
	}
	if state.elevatorPos < FloorsCnt-1 {
		nextElevatorPositions = append(nextElevatorPositions, state.elevatorPos+1)
	}
	steps := make([]Step, 0, 2)
	moves(state.elementsCnt, state.generators[state.elevatorPos], state.microchips[state.elevatorPos], func(generator, microchip uint8) {
		for _, nextElevatorPos := range nextElevatorPositions {
			s := state
			s.elevatorPos = nextElevatorPos
			s.generators[state.elevatorPos] &^= generator
			s.generators[s.elevatorPos] |= generator
			s.microchips[state.elevatorPos] &^= microchip
			s.microchips[s.elevatorPos] |= microchip
			if isFloorValid(s.generators[state.elevatorPos], s.microchips[state.elevatorPos]) &&
				isFloorValid(s.generators[s.elevatorPos], s.microchips[s.elevatorPos]) {
				steps = append(steps, Step{state: s, step: step})
			}
		}
	})
	return steps
}

func isFloorValid(generators, microchips uint8) bool {
	if microchips == 0 {
		return true
	}
	if generators == 0 {
		return true
	}
	if microchips&^generators == 0 {
		return true
	}
	return false
}

func moves(elementsCnt uint8, generators, microchips uint8, cb func(generator, microchips uint8)) {
	g1 := takeOneBit(generators, elementsCnt)
	m1 := takeOneBit(microchips, elementsCnt)
	g2 := takeTwoBits(generators, elementsCnt)
	m2 := takeTwoBits(microchips, elementsCnt)
	gm := pairs(g1, m1)
	for _, b := range g1 {
		cb(b, 0)
	}
	for _, b := range m1 {
		cb(0, b)
	}
	for _, b := range g2 {
		cb(b, 0)
	}
	for _, b := range m2 {
		cb(0, b)
	}
	for _, p := range gm {
		cb(p[0], p[1])
	}
}

func takeOneBit(bits uint8, bitsCnt uint8) []uint8 {
	var result []uint8
	for i := range bitsCnt {
		var bit uint8 = 1 << i
		if bits&bit > 0 {
			result = append(result, bit)
		}
	}
	return result
}

func takeTwoBits(bits uint8, bitsCnt uint8) []uint8 {
	var result []uint8
	for a := range bitsCnt {
		var bit1 uint8 = 1 << a
		for b := a + 1; b < bitsCnt; b++ {
			var bit2 uint8 = 1 << b
			res := bit1 | bit2
			if bits&res == res {
				result = append(result, res)
			}
		}
	}
	return result
}

func pairs(s1, s2 []uint8) [][2]uint8 {
	result := make([][2]uint8, 0, len(s1)*len(s2))
	for _, e1 := range s1 {
		for _, e2 := range s2 {
			result = append(result, [2]uint8{e1, e2})
		}
	}
	return result
}

func makeTargetState(elementsCnt uint8) State {
	return State{
		elevatorPos: FloorsCnt - 1,
		elementsCnt: elementsCnt,
		microchips:  makeFullTopFloor(elementsCnt),
		generators:  makeFullTopFloor(elementsCnt),
	}
}

func makeFullTopFloor(elementsCnt uint8) [FloorsCnt]uint8 {
	var floor [FloorsCnt]uint8
	floor[FloorsCnt-1] = uint8(^(^uint8(0) << elementsCnt))
	return floor
}

var elementRE = regexp.MustCompile(`(\w+) generator|(\w+)-compatible microchip`)

func getInput(fileName string) (State, map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return State{}, nil, err
	}
	scanner := bufio.NewScanner(file)

	shifts := make(map[string]int)
	currShift := 0
	getShift := func(item string) int {
		if shift, ok := shifts[item]; ok {
			return shift
		} else {
			shifts[item] = currShift
			currShift++
			if currShift >= 9 {
				panic("currShift >= 9")
			}
			return currShift - 1
		}
	}

	state := State{}
	floor := 0

	for scanner.Scan() {
		if floor >= FloorsCnt {
			return State{}, nil, fmt.Errorf("floor >= %d", FloorsCnt)
		}
		line := scanner.Text()
		matches := elementRE.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[1] != "" {
				state.generators[floor] |= 1 << getShift(match[1])
			} else if match[2] != "" {
				state.microchips[floor] |= 1 << getShift(match[2])
			} else {
				return State{}, nil, errors.New("match failed")
			}
		}
		floor++
	}
	if err := scanner.Err(); err != nil {
		return State{}, nil, err
	}

	state.elementsCnt = uint8(len(shifts))

	return state, shifts, nil
}
