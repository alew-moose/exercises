package main

import "fmt"

func main() {
	// elvesCnt := 5 // test input
	elvesCnt := 3014603
	fmt.Println("part 1:", solvePart1(elvesCnt))
	fmt.Println("part 2:", solvePart2(elvesCnt))
}

type Elve struct {
	N        int
	Presents int
	Prev     *Elve
	Next     *Elve
}

func makeElves(cnt int) *Elve {
	if cnt < 1 {
		return nil
	}
	head := &Elve{
		N:        1,
		Presents: 1,
		Prev:     nil,
		Next:     nil,
	}
	curr := head
	for n := 2; n <= cnt; n++ {
		next := &Elve{
			N:        n,
			Presents: 1,
			Prev:     curr,
			Next:     nil,
		}
		curr.Next = next
		curr = next
	}
	curr.Next = head
	head.Prev = curr
	return head
}

func nextNElve(elve *Elve, n int) *Elve {
	if n <= 0 {
		panic("invalid n")
	}
	for n > 0 {
		elve = elve.Next
		n--

	}
	return elve
}

func solvePart1(elvesCnt int) int {
	elve := makeElves(elvesCnt)
	for elve.Next != elve {
		elve.Presents += elve.Next.Presents
		elve.Next = elve.Next.Next
		elve = elve.Next
	}
	return elve.N
}

func solvePart2(elvesCnt int) int {
	elve := makeElves(elvesCnt)
	across := nextNElve(elve, elvesCnt/2)
	for elve.Next != elve {
		if elvesCnt%2 == 0 {
			across = across.Next
		}
		beforeAcross := across.Prev
		afterAcross := across.Next
		beforeAcross.Next = afterAcross
		afterAcross.Prev = beforeAcross
		across = across.Next
		elve = elve.Next
		elvesCnt--
	}
	return elve.N
}
