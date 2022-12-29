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
	fmt.Printf("Part 2: %s\n", part2(input))
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

func part2(input []string) string {
	var register [241]int

	// It should be possible to evaluate the register and CRT in a single loop.
	// I'm not smart enough to do that.
	// I know that I'm going to evaluate 240 cycles, I'm going to do that first.
	pc := 1
	x := 1

	for _, line := range input {
		register[pc] = x

		switch {
		case line == "noop":
			pc++
		default:
			register[pc+1] = x
			register[pc+2] = x
			pc += 2
			parts := strings.Fields(line)
			arg, _ := strconv.Atoi(parts[1])
			x += arg
		}
	}

	// model my screen as an array of bytes.
	var screen [240]byte

	// iterate over each pixel.
	for i := range screen {
		pos := i % 40
		val := register[i+1]

		// if the pixel falls within the bounds of the sprite, write a star.
		if pos >= val-1 && pos <= val+1 {
			screen[i] = '*'
			continue
		}

		// otherwise, write a space. this makes it easier to read later.
		screen[i] = ' '
	}

	var sb strings.Builder

	// use a string builder to turn the screen into a rectangular grid.

	// leading newline helps make it legible in the output
	sb.WriteByte('\n')
	for i := 0; i < 240; i += 40 {
		sb.Write(screen[i : i+40]) // write one row at a time.
		sb.WriteByte('\n')
	}

	return sb.String()
}
