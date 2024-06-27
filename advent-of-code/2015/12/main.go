package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input := readInput()
	fmt.Println("part1:", allNumbersSum(input))
	fmt.Println("part2:", allNumbersSum2(input))

}

func readInput() map[string]any {
	input := make(map[string]any)
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bytes, &input); err != nil {
		log.Fatal(err)
	}
	return input
}

func allNumbersSum(input any) float64 {
	var sum float64
	switch i := input.(type) {
	case map[string]any:
		for _, v := range i {
			sum += allNumbersSum(v)
		}
	case []any:
		for _, n := range i {
			sum += allNumbersSum(n)
		}
	case float64:
		sum = i
	case int:
		sum = float64(i)
	case string:
		sum = 0
	default:
		log.Fatalf("invalid type %T\n", i)
	}
	return sum
}

// ignore objects with "red"
func allNumbersSum2(input any) float64 {
	var sum float64
	switch i := input.(type) {
	case map[string]any:
		for _, v := range i {
			switch s := v.(type) {
			case string:
				if s == "red" {
					return 0
				}
			}
		}
		for _, v := range i {
			sum += allNumbersSum2(v)
		}
	case []any:
		for _, n := range i {
			sum += allNumbersSum2(n)
		}
	case float64:
		sum = i
	case int:
		sum = float64(i)
	case string:
		sum = 0
	default:
		log.Fatalf("invalid type %T\n", i)
	}
	return sum
}
