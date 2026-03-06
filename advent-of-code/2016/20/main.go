package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func main() {
	// ranges, err := getInput("test-input.txt")
	ranges, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	slices.SortFunc(ranges, func(r1, r2 Range) int {
		return int(uint64(r1.From) - uint64(r2.From))
	})
	fmt.Println("part 1:", solvePart1(ranges))
	fmt.Println("part 2:", solvePart2(ranges))
}

func solvePart1(ranges []Range) uint32 {
	var minNotBlocked uint32
	for _, r := range ranges {
		if r.From <= minNotBlocked && minNotBlocked <= r.To {
			minNotBlocked = r.To + 1
		}
	}
	return minNotBlocked
}

func solvePart2(ranges []Range) int64 {
	var allowedCnt int64 = 1 << 32
	ranges = mergeRanges(ranges)
	for _, r := range ranges {
		allowedCnt -= int64(r.To) - int64(r.From) + 1
	}
	return allowedCnt
}

func mergeRanges(ranges []Range) []Range {
	merged := make([]Range, 0, len(ranges))
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}
		prevRange := merged[len(merged)-1]
		if prevRange.To >= r.From {
			if r.To <= prevRange.To {
				continue
			}
			merged[len(merged)-1].To = r.To
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

type Range struct {
	From, To uint32
}

func getInput(fileName string) ([]Range, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var ranges []Range
	for {
		var r Range
		_, err := fmt.Fscanf(f, "%d-%d\n", &r.From, &r.To)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, r)
	}
	return ranges, nil
}
