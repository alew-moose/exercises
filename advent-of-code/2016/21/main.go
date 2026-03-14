package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// commands, err := getInput("test-input.txt")
	commands, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part 1:", solvePart1(commands))
	fmt.Println("part 2:", solvePart2(commands))
}

func solvePart1(commands []Command) string {
	b := []byte("abcdefgh")
	for _, cmd := range commands {
		cmd.Execute(b)
	}
	return string(b)
}

func solvePart2(commands []Command) string {
	b := []byte("fbgdceah")
	reverseCmds := reverseCommands(commands)
	for _, cmd := range reverseCmds {
		cmd.Execute(b)
	}
	return string(b)
}

func reverseCommands(commands []Command) []Command {
	reverseCmds := make([]Command, 0, len(commands))
	for i := len(commands) - 1; i >= 0; i-- {
		reverseCmds = append(reverseCmds, reverseCommand(commands[i]))
	}
	return reverseCmds
}

func reverseCommand(command Command) Command {
	switch cmd := command.(type) {
	case *CmdSwapPosition, *CmdSwapLetter, *CmdReversePositions:
		return cmd
	case *CmdRotateLeft:
		return &CmdRotateRight{cmd.n}
	case *CmdRotateRight:
		return &CmdRotateLeft{cmd.n}
	case *CmdMove:
		return &CmdMove{cmd.b, cmd.a}
	case *CmdRotateLetterPos:
		return &CmdUnrotateLetterPos{cmd.l}
	default:
		panic("unknown command type")
	}
}

type Command interface {
	Execute([]byte)
}

func parseCommand(s string) (Command, error) {
	if cmd := parseCmdSwapPosition(s); cmd != nil {
		return cmd, nil
	}
	if cmd := parseCmdSwapLetter(s); cmd != nil {
		return cmd, nil
	}
	if cmd := parseCmdRotateLeft(s); cmd != nil {
		return cmd, nil
	}
	if cmd := parseCmdRotateRight(s); cmd != nil {
		return cmd, nil
	}
	if cmd := parseCmdRotateLetterPos(s); cmd != nil {
		return cmd, nil
	}
	if cmd := parseCmdReversePositions(s); cmd != nil {
		return cmd, nil
	}
	if cmd := parseCmdMove(s); cmd != nil {
		return cmd, nil
	}
	return nil, fmt.Errorf("failed to parse command %q", s)
}

type CmdSwapPosition struct {
	a, b int
}

func (c *CmdSwapPosition) Execute(s []byte) {
	s[c.a], s[c.b] = s[c.b], s[c.a]
}

var cmdSwapPositionRE = regexp.MustCompile(`^swap position (\d+) with position (\d+)$`)

func parseCmdSwapPosition(s string) *CmdSwapPosition {
	matches := cmdSwapPositionRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdSwapPosition
	var err error
	cmd.a, err = strconv.Atoi(matches[0][1])
	if err != nil {
		return nil
	}
	cmd.b, err = strconv.Atoi(matches[0][2])
	if err != nil {
		return nil
	}
	return &cmd
}

type CmdSwapLetter struct {
	a, b byte
}

func (c *CmdSwapLetter) Execute(s []byte) {
	ai, bi := -1, -1
	for i, l := range s {
		if l == c.a {
			ai = i
		} else if l == c.b {
			bi = i
		}
		switch l {
		case c.a:
			ai = i
		case c.b:
			bi = i
		}
		if ai >= 0 && bi >= 0 {
			break
		}
	}
	if ai < 0 || bi < 0 {
		panic("letter not found")
	}
	s[ai], s[bi] = s[bi], s[ai]
}

var cmdSwapLetterRE = regexp.MustCompile(`^swap letter ([a-z]) with letter ([a-z])$`)

func parseCmdSwapLetter(s string) *CmdSwapLetter {
	matches := cmdSwapLetterRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdSwapLetter
	cmd.a = byte(matches[0][1][0])
	cmd.b = byte(matches[0][2][0])
	return &cmd
}

type CmdRotateLeft struct {
	n int
}

func (c *CmdRotateLeft) Execute(s []byte) {
	steps := c.n % len(s)
	if steps == 0 {
		return
	}
	b := make([]byte, len(s))
	for i := range b {
		b[i] = s[(i+steps)%len(s)]
	}
	copy(s, b)
}

var cmdRotateLeftRE = regexp.MustCompile(`^rotate left (\d+) steps?$`)

func parseCmdRotateLeft(s string) *CmdRotateLeft {
	matches := cmdRotateLeftRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdRotateLeft
	var err error
	cmd.n, err = strconv.Atoi(matches[0][1])
	if err != nil {
		return nil
	}
	return &cmd
}

type CmdRotateRight struct {
	n int
}

func (c *CmdRotateRight) Execute(s []byte) {
	steps := c.n % len(s)
	if steps == 0 {
		return
	}
	b := make([]byte, len(s))
	for i := range b {
		b[i] = s[(i-steps+len(s))%len(s)]
	}
	copy(s, b)
}

var cmdRotateRightRE = regexp.MustCompile(`^rotate right (\d+) steps?$`)

func parseCmdRotateRight(s string) *CmdRotateRight {
	matches := cmdRotateRightRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdRotateRight
	var err error
	cmd.n, err = strconv.Atoi(matches[0][1])
	if err != nil {
		return nil
	}
	return &cmd
}

type CmdRotateLetterPos struct {
	l byte
}

func (c *CmdRotateLetterPos) Execute(s []byte) {
	pos := bytes.IndexByte(s, c.l)
	if pos < 0 {
		panic("letter not found")
	}
	steps := 1 + pos
	if pos >= 4 {
		steps++
	}
	cmd := CmdRotateRight{steps}
	cmd.Execute(s)
}

var cmdRotateLetterPosRE = regexp.MustCompile(`^rotate based on position of letter ([a-z])$`)

func parseCmdRotateLetterPos(s string) *CmdRotateLetterPos {
	matches := cmdRotateLetterPosRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdRotateLetterPos
	cmd.l = byte(matches[0][1][0])
	return &cmd
}

type CmdReversePositions struct {
	a, b int
}

func (c *CmdReversePositions) Execute(s []byte) {
	i, k := c.a, c.b
	for i < k {
		s[i], s[k] = s[k], s[i]
		i++
		k--
	}
}

type CmdUnrotateLetterPos struct {
	l byte
}

func (c *CmdUnrotateLetterPos) Execute(s []byte) {
	// 01234567
	// 0 -> steps=1; pos=1
	// 1 -> steps=2; pos=3
	// 2 -> steps=3; pos=5
	// 3 -> steps=4; pos=7
	// 4 -> steps=6; pos=2
	// 5 -> steps=7; pos=4
	// 6 -> steps=8; pos=6
	// 7 -> steps=9; pos=0
	pos := bytes.IndexByte(s, c.l)
	if pos < 0 {
		panic("letter not found")
	}
	var cmd Command
	switch pos {
	case 1:
		cmd = &CmdRotateLeft{1}
	case 3:
		cmd = &CmdRotateLeft{2}
	case 5:
		cmd = &CmdRotateLeft{3}
	case 7:
		cmd = &CmdRotateLeft{4}
	case 2:
		cmd = &CmdRotateRight{2}
	case 4:
		cmd = &CmdRotateRight{1}
	case 6:
		return
	case 0:
		cmd = &CmdRotateLeft{1}
	default:
		panic("invalid index")
	}
	cmd.Execute(s)
}

var cmdReversePositionsRE = regexp.MustCompile(`^reverse positions (\d+) through (\d+)$`)

func parseCmdReversePositions(s string) *CmdReversePositions {
	matches := cmdReversePositionsRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdReversePositions
	var err error
	cmd.a, err = strconv.Atoi(matches[0][1])
	if err != nil {
		return nil
	}
	cmd.b, err = strconv.Atoi(matches[0][2])
	if err != nil {
		return nil
	}
	return &cmd
}

type CmdMove struct {
	a, b int
}

func (c *CmdMove) Execute(s []byte) {
	if c.a < c.b {
		b := make([]byte, 0, len(s))
		b = append(b, s[:c.a]...)
		b = append(b, s[c.a+1:c.b+1]...)
		b = append(b, s[c.a])
		copy(s, b)
	} else if c.a > c.b {
		b := make([]byte, 0, len(s))
		b = append(b, s[:c.b]...)
		b = append(b, s[c.a])
		b = append(b, s[c.b:c.a]...)
		b = append(b, s[c.a+1:]...)
		copy(s, b)
	}
}

var cmdMoveRE = regexp.MustCompile(`^move position (\d+) to position (\d+)$`)

func parseCmdMove(s string) *CmdMove {
	matches := cmdMoveRE.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return nil
	}
	var cmd CmdMove
	var err error
	cmd.a, err = strconv.Atoi(matches[0][1])
	if err != nil {
		return nil
	}
	cmd.b, err = strconv.Atoi(matches[0][2])
	if err != nil {
		return nil
	}
	return &cmd
}

func getInput(file string) ([]Command, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var commands []Command
	for scanner.Scan() {
		s := scanner.Text()
		cmd, err := parseCommand(s)
		if err != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return commands, nil
}
