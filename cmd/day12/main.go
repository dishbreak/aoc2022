package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	g := GridFromInput(input)
	return g.ShortestPathTo(g.ReachStartPoint)
}

func part2(input []string) int {
	g := GridFromInput(input)
	return g.ShortestPathTo(g.ReachZeroElevation)
}
