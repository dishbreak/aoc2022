package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

func part1(input string) int {
	c := NewCharRegister()

	for i, b := range input {
		c.Add(byte(b))

		if c.Match() {
			return i + 1
		}
	}

	return -1
}

func part2(input string) int {
	c := NewCharRegister(WithWindowSize(14))

	for i, b := range input {
		c.Add(byte(b))

		if c.Match() {
			return i + 1
		}
	}

	return -1
}
