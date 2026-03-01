package main

import "fmt"

func main() {
	// favNum := 10                // test favorite number
	// target := Coord{x: 7, y: 4} // test target
	favNum := 1364
	target := Coord{x: 31, y: 39}
	fmt.Println("part 1:", solvePart1(favNum, target))
	fmt.Println("part 2:", solvePart2(favNum))
}

func solvePart1(favNum int, target Coord) int {
	m := newMap(favNum)
	start := Coord{x: 1, y: 1}
	return shortestPathLen(m, start, target)
}

func shortestPathLen(m *Map, start, target Coord) int {
	visited := make(map[Coord]struct{})
	queue := []Step{{pos: start, steps: 0}}
	for len(queue) > 0 {
		step := queue[0]
		queue = queue[1:]
		if step.pos == target {
			return step.steps
		}
		visited[step.pos] = struct{}{}
		queue = append(queue, nextSteps(step, m, visited)...)
	}
	return -1
}

func solvePart2(favNum int) int {
	m := newMap(favNum)
	start := Coord{x: 1, y: 1}
	maxSteps := 50
	return countLocations(m, start, maxSteps)
}

func countLocations(m *Map, start Coord, maxSteps int) int {
	visited := make(map[Coord]struct{})
	queue := []Step{{pos: start, steps: 0}}
	for len(queue) > 0 {
		step := queue[0]
		queue = queue[1:]
		if step.steps > maxSteps {
			break
		}
		visited[step.pos] = struct{}{}
		queue = append(queue, nextSteps(step, m, visited)...)
	}
	return len(visited)
}

func nextSteps(step Step, m *Map, visited map[Coord]struct{}) []Step {
	steps := make([]Step, 0, 4)
	stepCandidates := [...]Step{
		{steps: step.steps + 1, pos: Coord{x: step.pos.x, y: step.pos.y - 1}},
		{steps: step.steps + 1, pos: Coord{x: step.pos.x + 1, y: step.pos.y}},
		{steps: step.steps + 1, pos: Coord{x: step.pos.x, y: step.pos.y + 1}},
		{steps: step.steps + 1, pos: Coord{x: step.pos.x - 1, y: step.pos.y}},
	}
	for _, nextStep := range stepCandidates {
		if nextStep.pos.x < 0 || nextStep.pos.y < 0 {
			continue
		}
		if _, ok := visited[nextStep.pos]; ok || m.GetField(nextStep.pos) == FieldWall {
			continue
		}
		steps = append(steps, nextStep)
	}
	return steps
}

type Step struct {
	pos   Coord
	steps int
}

type Coord struct {
	x, y int
}

type Field uint8

const (
	FieldOpen Field = 0
	FieldWall Field = 1
)

type Map struct {
	favNum int
	fields map[Coord]Field
}

func newMap(favNum int) *Map {
	return &Map{
		favNum: favNum,
		fields: make(map[Coord]Field),
	}
}

func (m *Map) GetField(c Coord) Field {
	if field, ok := m.fields[c]; ok {
		return field
	}
	field := calculateField(c, m.favNum)
	m.fields[c] = field
	return field
}

func calculateField(c Coord, favNum int) Field {
	// x*x + 3*x + 2*x*y + y + y*y
	n := uint64(c.x*c.x + 3*c.x + 2*c.x*c.y + c.y + c.y*c.y + favNum)
	if popCount(n)%2 == 0 {
		return FieldOpen
	}
	return FieldWall
}

func popCount(n uint64) int {
	pc := 0
	for n > 0 {
		if n&1 == 1 {
			pc++
		}
		n >>= 1
	}
	return pc
}
