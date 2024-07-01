package main

import "fmt"

type ingredient struct {
	title                                           string
	capacity, durability, flavor, texture, calories int
}

func main() {
	ingredients := []ingredient{
		{title: "Sugar", capacity: 3, durability: 0, flavor: 0, texture: -3, calories: 2},
		{title: "Sprinkles", capacity: -3, durability: 3, flavor: 0, texture: 0, calories: 9},
		{title: "Candy", capacity: -1, durability: 0, flavor: 4, texture: 0, calories: 1},
		{title: "Chocolate", capacity: 0, durability: 0, flavor: -2, texture: 2, calories: 8},
	}

	fmt.Println("part1:", solvePart1(ingredients))
	fmt.Println("part2:", solvePart2(ingredients))
}

func solvePart1(ingredients []ingredient) int {
	amounts := make([]int, 4)
	maxScore := 0
	variants(amounts, 100, func() {
		maxScore = max(maxScore, score(ingredients, amounts))
	})
	return maxScore
}

func solvePart2(ingredients []ingredient) int {
	amounts := make([]int, 4)
	maxScore := 0
	variants(amounts, 100, func() {
		if calories(ingredients, amounts) == 500 {
			maxScore = max(maxScore, score(ingredients, amounts))
		}
	})
	return maxScore
}

func score(ingredients []ingredient, amounts []int) int {
	var capacity, durability, flavor, texture int
	for i, amount := range amounts {
		capacity += ingredients[i].capacity * amount
		durability += ingredients[i].durability * amount
		flavor += ingredients[i].flavor * amount
		texture += ingredients[i].texture * amount
	}
	capacity = max(capacity, 0)
	durability = max(durability, 0)
	flavor = max(flavor, 0)
	texture = max(texture, 0)
	return capacity * durability * flavor * texture
}

func calories(ingredients []ingredient, amounts []int) int {
	c := 0
	for i, amount := range amounts {
		c += ingredients[i].calories * amount
	}
	return c
}

func variants(v []int, left int, cb func()) {
	if left == 0 && len(v) == 0 {
		cb()
	}
	if len(v) == 0 {
		return
	}
	for n := 0; n <= left; n++ {
		v[0] = n
		variants(v[1:], left-n, cb)
	}
}
