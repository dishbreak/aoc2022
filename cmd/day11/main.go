package main

import (
	"fmt"
	"sort"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day11.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input [][]string) int64 {
	return playGame(input, 20)
}

func playGame(input [][]string, rounds int, opts ...MonkeyOption) int64 {
	monkeys := make([]*Monkey, len(input))

	for i, block := range input {
		monkeys[i] = MonkeyFromInput(block, opts...)
	}

	for _, m := range monkeys {
		m.Connect(monkeys)
	}

	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			m.Inspect()
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}

func makeMonkeys(input [][]string) []*Monkey {
	monkeys := make([]*Monkey, len(input))

	for i, block := range input {
		monkeys[i] = MonkeyFromInput(block)
	}

	for _, m := range monkeys {
		m.Connect(monkeys)
	}

	return monkeys
}
