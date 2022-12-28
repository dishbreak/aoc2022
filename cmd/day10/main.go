package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	acc, x := 0, 1
	cycles := 0
	target := 20

	recordVal := func() {
		acc += target * x
		target += 40
	}

	for _, line := range input {

		switch {
		case line == "noop":
			cycles++
		default:
			if cycles == target-1 || cycles == target-2 {
				recordVal()
			}
			parts := strings.Fields(line)
			arg, _ := strconv.Atoi(parts[1])
			x += arg
			cycles += 2
		}

		if cycles == target {
			recordVal()
		}

		if cycles == 220 {
			return acc
		}
	}

	return acc
}
