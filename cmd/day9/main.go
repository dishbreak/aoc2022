package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	fmt.Printf("Part 1: %d\n", part1(input))
}

type Instruction struct {
	Direction image.Point
	Count     int
}

var directions map[string]image.Point = map[string]image.Point{
	"U": image.Pt(0, 1),
	"D": image.Pt(0, -1),
	"L": image.Pt(-1, 0),
	"R": image.Pt(1, 0),
}

func instructionFromString(line string) Instruction {
	parts := strings.Fields(line)
	ct, _ := strconv.Atoi(parts[1])

	return Instruction{
		Direction: directions[parts[0]],
		Count:     ct,
	}
}

func part1(input []string) int {
	r := &Rope{
		tailSpots: make(map[image.Point]bool),
	}

	for _, line := range input {
		inst := instructionFromString(line)
		for i := 0; i < inst.Count; i++ {
			r.Move(inst.Direction)
		}
	}

	return r.TailHits()
}
