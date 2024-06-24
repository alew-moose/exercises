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
	and = iota
	or
	lshift
	rshift
	not
	assign
)

type instruction struct {
	op         int
	arg1, arg2 operand
	dest       string
}

const (
	value = iota
	reference
)

type operand struct {
	t   int
	val uint16
	ref string
}

func main() {
	instrs := readInput()
	wires := make(map[string]instruction)
	for _, instr := range instrs {
		wires[instr.dest] = instr
	}
	part1result := evalInstruction(wires, "a")
	fmt.Println("part1:", part1result)

	clear(wires)
	for _, instr := range instrs {
		wires[instr.dest] = instr
	}
	wires["b"] = instruction{
		op: assign,
		arg1: operand{
			t:   value,
			val: part1result,
		},
		dest: "b",
	}
	part2result := evalInstruction(wires, "a")
	fmt.Println("part1:", part2result)

}

func evalInstruction(wires map[string]instruction, wire string) uint16 {
	instr := wires[wire]
	switch instr.op {
	case assign:
		return evalOperand(wires, instr.arg1)
	case not:
		return ^evalOperand(wires, instr.arg1)
	case and:
		return evalOperand(wires, instr.arg1) & evalOperand(wires, instr.arg2)
	case or:
		return evalOperand(wires, instr.arg1) | evalOperand(wires, instr.arg2)
	case lshift:
		return evalOperand(wires, instr.arg1) << evalOperand(wires, instr.arg2)
	case rshift:
		return evalOperand(wires, instr.arg1) >> evalOperand(wires, instr.arg2)
	}
	panic("invalid op")
}

func evalOperand(wires map[string]instruction, arg operand) uint16 {
	if arg.t == value {
		return arg.val
	}

	val := evalInstruction(wires, arg.ref)
	wires[arg.ref] = instruction{
		op: assign,
		arg1: operand{
			t:   value,
			val: val,
		},
	}
	return val
}

func readInput() []instruction {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)

	var instrs []instruction
	for s.Scan() {
		instrs = append(instrs, parseInstruction(s.Text()))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return instrs
}

func parseInstruction(s string) instruction {
	parts := strings.Split(s, " -> ")
	if len(parts) != 2 {
		panic("failed to parse instruction")
	}

	dest := parts[1]
	expr := strings.Split(parts[0], " ")

	if len(expr) == 1 {
		return instruction{
			op:   assign,
			arg1: parseOperand(expr[0]),
			dest: dest,
		}
	}

	if expr[0] == "NOT" {
		return instruction{
			op:   not,
			arg1: parseOperand(expr[1]),
			dest: dest,
		}
	}

	var op int
	switch expr[1] {
	case "AND":
		op = and
	case "OR":
		op = or
	case "LSHIFT":
		op = lshift
	case "RSHIFT":
		op = rshift
	default:
		panic("unknown operand " + expr[1])
	}

	return instruction{
		op:   op,
		arg1: parseOperand(expr[0]),
		arg2: parseOperand(expr[2]),
		dest: dest,
	}
}

func parseOperand(s string) operand {
	val, err := strconv.Atoi(s)
	if err != nil {
		return operand{
			t:   reference,
			ref: s,
		}
	}
	return operand{
		t:   value,
		val: uint16(val),
	}
}
