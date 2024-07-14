package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instructionType int
type outputType int

const (
	instructionTypeValue instructionType = iota
	instructionTypeBot
)

const (
	outputTypeBot outputType = iota
	outputTypeOutput
)

type instruction struct {
	t      instructionType
	value  int
	bot    int
	lowTo  output
	highTo output
}

type output struct {
	t  outputType
	id int
}

func main() {
	instrs := readInput()
	fmt.Println("part1:", solvePart1(instrs, 17, 61))
	fmt.Println("part1:", solvePart2(instrs, -1, -1))
}

func readInput() []instruction {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var instrs []instruction
	for s.Scan() {
		instr, err := parseInstruction(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		instrs = append(instrs, instr)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return instrs
}

func parseInstruction(s string) (instruction, error) {
	var instr instruction
	var err error
	fields := strings.Fields(s)

	if fields[0] == "value" {
		instr.t = instructionTypeValue
		instr.value, err = strconv.Atoi(fields[1])
		if err != nil {
			return instruction{}, err
		}
		instr.lowTo, err = parseOutput(fields[4], fields[5])
		if err != nil {
			return instruction{}, err
		}
		return instr, nil
	}

	instr.t = instructionTypeBot
	instr.bot, err = strconv.Atoi(fields[1])
	if err != nil {
		return instruction{}, err
	}
	instr.lowTo, err = parseOutput(fields[5], fields[6])
	if err != nil {
		return instruction{}, err
	}
	instr.highTo, err = parseOutput(fields[10], fields[11])
	if err != nil {
		return instruction{}, err
	}

	return instr, nil
}

func parseOutput(typeStr, idStr string) (o output, err error) {
	if typeStr == "bot" {
		o.t = outputTypeBot
	} else {
		o.t = outputTypeOutput
	}
	o.id, err = strconv.Atoi(idStr)
	if err != nil {
		return output{}, err
	}
	return
}

func solvePart1(instrs []instruction, searchLow, searchHigh int) int {
	outputs := make(map[int]int)
	botsValues := make(map[int][]int)
	botsTodo := make(map[int]instruction)

	for _, instr := range instrs {
		if instr.t == instructionTypeValue {
			botID := sendValue(instr.value, instr.lowTo, searchLow, searchHigh, outputs, botsValues, botsTodo)
			if botID != -1 {
				return botID
			}
		} else {
			botID := botSendValue(instr, searchLow, searchHigh, outputs, botsValues, botsTodo)
			if botID != -1 {
				return botID
			}
		}
	}
	return -1
}

func solvePart2(instrs []instruction, searchLow, searchHigh int) int {
	outputs := make(map[int]int)
	botsValues := make(map[int][]int)
	botsTodo := make(map[int]instruction)

	for _, instr := range instrs {
		if instr.t == instructionTypeValue {
			sendValue(instr.value, instr.lowTo, searchLow, searchHigh, outputs, botsValues, botsTodo)
		} else {
			botSendValue(instr, searchLow, searchHigh, outputs, botsValues, botsTodo)
		}
	}
	return outputs[0] * outputs[1] * outputs[2]
}

func sendValue(val int, o output, searchLow, searchHigh int, outputs map[int]int, botsValues map[int][]int, botsTodo map[int]instruction) int {
	if o.t == outputTypeOutput {
		outputs[o.id] = val
	} else {
		botsValues[o.id] = append(botsValues[o.id], val)
		if i, ok := botsTodo[o.id]; ok {
			if botID := botSendValue(i, searchLow, searchHigh, outputs, botsValues, botsTodo); botID != -1 {
				return botID
			}
		}
	}
	return -1
}

func botSendValue(instr instruction, searchLow, searchHigh int, outputs map[int]int, botsValues map[int][]int, botsTodo map[int]instruction) int {
	if len(botsValues[instr.bot]) < 2 {
		botsTodo[instr.bot] = instr
		return -1
	}
	var lowVal, highVal int
	if botsValues[instr.bot][0] < botsValues[instr.bot][1] {
		lowVal = botsValues[instr.bot][0]
		highVal = botsValues[instr.bot][1]
	} else {
		highVal = botsValues[instr.bot][0]
		lowVal = botsValues[instr.bot][1]
	}
	if searchLow != -1 && searchHigh != -1 && lowVal == searchLow && highVal == searchHigh {
		return instr.bot
	}

	if botID := sendValue(lowVal, instr.lowTo, searchLow, searchHigh, outputs, botsValues, botsTodo); botID != -1 {
		return botID
	}
	if botID := sendValue(highVal, instr.highTo, searchLow, searchHigh, outputs, botsValues, botsTodo); botID != -1 {
		return botID
	}

	return -1
}
