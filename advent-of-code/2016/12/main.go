package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// instrs, err := getInput("test-input.txt")
	instrs, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 1:", solvePart1(instrs))
	fmt.Println("part 2:", solvePart2(instrs))
}

func solvePart1(instrs []Instruction) int {
	p := newProcessor()
	p.Run(instrs)
	return p.registers["a"]
}

func solvePart2(instrs []Instruction) int {
	p := newProcessor()
	p.registers["c"] = 1
	p.Run(instrs)
	return p.registers["a"]
}

type Processor struct {
	pc        int
	registers map[string]int
}

func newProcessor() *Processor {
	return &Processor{
		pc:        0,
		registers: make(map[string]int),
	}
}

func (p *Processor) Run(instrs []Instruction) {
	for p.pc >= 0 && p.pc < len(instrs) {
		instrs[p.pc].Execute(p)
	}
}

type Instruction interface {
	Execute(*Processor)
}

type InstrInc struct {
	register string
}

func (i *InstrInc) Execute(p *Processor) {
	p.registers[i.register]++
	p.pc++
}

type InstrDec struct {
	register string
}

func (i *InstrDec) Execute(p *Processor) {
	p.registers[i.register]--
	p.pc++
}

type InstrCpyVal struct {
	value       int
	registerDst string
}

func (i *InstrCpyVal) Execute(p *Processor) {
	p.registers[i.registerDst] = i.value
	p.pc++
}

type InstrCpyReg struct {
	registerSrc string
	registerDst string
}

func (i *InstrCpyReg) Execute(p *Processor) {
	p.registers[i.registerDst] = p.registers[i.registerSrc]
	p.pc++
}

type InstrJnzVal struct {
	offset    int
	testValue int
}

func (i *InstrJnzVal) Execute(p *Processor) {
	if i.testValue != 0 {
		p.pc += i.offset
	} else {
		p.pc++
	}
}

type InstrJnzReg struct {
	offset       int
	testRegister string
}

func (i *InstrJnzReg) Execute(p *Processor) {
	if p.registers[i.testRegister] != 0 {
		p.pc += i.offset
	} else {
		p.pc++
	}
}

func parseInstr(s string) (Instruction, error) {
	tokens := strings.Split(s, " ")
	switch len(tokens) {
	case 2:
		reg, err := parseReg(tokens[1])
		if err != nil {
			return nil, err
		}
		switch tokens[0] {
		case "inc":
			return &InstrInc{register: reg}, nil
		case "dec":
			return &InstrDec{register: reg}, nil
		default:
			return nil, fmt.Errorf("invalid instruction %q", tokens[0])
		}
	case 3:
		switch tokens[0] {
		case "cpy":
			regDst, err := parseReg(tokens[2])
			if err != nil {
				return nil, err
			}
			if val, err := strconv.Atoi(tokens[1]); err == nil {
				return &InstrCpyVal{registerDst: regDst, value: val}, nil
			}
			if regSrc, err := parseReg(tokens[1]); err == nil {
				return &InstrCpyReg{registerDst: regDst, registerSrc: regSrc}, nil
			}
			return nil, fmt.Errorf("failed to parse %q", tokens[1])
		case "jnz":
			offset, err := strconv.Atoi(tokens[2])
			if err != nil {
				return nil, fmt.Errorf("failed to parse offset %q", tokens[2])
			}
			if testVal, err := strconv.Atoi(tokens[1]); err == nil {
				return &InstrJnzVal{offset: offset, testValue: testVal}, nil
			}
			if testReg, err := parseReg(tokens[1]); err == nil {
				return &InstrJnzReg{offset: offset, testRegister: testReg}, nil
			}
			return nil, fmt.Errorf("failet to parse %q", tokens[1])
		default:
			return nil, fmt.Errorf("invalid instruction %q", tokens[0])
		}
	default:
		return nil, fmt.Errorf("invalid number of tokens: %d", len(tokens))
	}
}

func parseReg(s string) (string, error) {
	switch s {
	case "a", "b", "c", "d":
		return s, nil
	default:
		return "", errors.New("invalid register")
	}
}

func getInput(fileName string) ([]Instruction, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var instrs []Instruction
	for scanner.Scan() {
		instr, err := parseInstr(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to parse instruction %q: %s", scanner.Text(), err)
		}
		instrs = append(instrs, instr)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return instrs, nil
}
