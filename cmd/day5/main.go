package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part1(input [][]string) string {
	c := CrateLayoutFromStringSlice(input[0])

	for _, line := range input[1] {
		parts := strings.Fields(line)
		count, _ := strconv.Atoi(parts[1])
		src, _ := strconv.Atoi(parts[3])
		dest, _ := strconv.Atoi(parts[5])
		c.Move(count, src, dest)
	}

	return c.Topline()
}

func part2(input [][]string) string {
	c := CrateLayoutFromStringSlice(input[0], WithCraneModelNumber(9001))

	for _, line := range input[1] {
		parts := strings.Fields(line)
		count, _ := strconv.Atoi(parts[1])
		src, _ := strconv.Atoi(parts[3])
		dest, _ := strconv.Atoi(parts[5])
		c.Move(count, src, dest)
	}

	return c.Topline()
}

type CrateLayout struct {
	stacks           [][]byte
	multiCratePickup bool
}

func (c *CrateLayout) Move(count, srcIdx, destIdx int) {
	if c.multiCratePickup {
		srcLen := len(c.stacks[srcIdx])
		crates := c.stacks[srcIdx][srcLen-count:]
		c.stacks[srcIdx] = c.stacks[srcIdx][:srcLen-count]
		c.stacks[destIdx] = append(c.stacks[destIdx], crates...)
		return
	}

	for i := 0; i < count; i++ {
		crate := c.pop(srcIdx)
		c.push(destIdx, crate)
	}
}

func (c *CrateLayout) pop(idx int) byte {
	lastIdx := len(c.stacks[idx]) - 1
	crate := c.stacks[idx][lastIdx]
	c.stacks[idx] = c.stacks[idx][:lastIdx]
	return crate
}

func (c *CrateLayout) push(idx int, crate byte) {
	c.stacks[idx] = append(c.stacks[idx], crate)
}

type CrateLayoutOption func(*CrateLayout) error

func WithCraneModelNumber(modelNum int) CrateLayoutOption {
	multiCratePickup := false
	if modelNum == 9001 {
		multiCratePickup = true
	}
	return func(cl *CrateLayout) error {
		cl.multiCratePickup = multiCratePickup
		return nil
	}
}

func CrateLayoutFromStringSlice(input []string, opts ...CrateLayoutOption) *CrateLayout {
	lastLine := input[len(input)-1]
	parts := strings.Fields(lastLine)
	stackCt, _ := strconv.Atoi(parts[len(parts)-1])

	c := &CrateLayout{
		// it's so much easier to just index everything the same.
		stacks: make([][]byte, stackCt+1),
	}

	for i := range c.stacks {
		c.stacks[i] = make([]byte, 0)
	}

	for i := 1; i <= stackCt; i++ {
		col := (4 * i) - 3
		for y := len(input) - 2; y >= 0; y-- {
			n := input[y][col]
			if n == ' ' {
				break
			}
			c.stacks[i] = append(c.stacks[i], n)
		}
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *CrateLayout) Topline() string {
	var sb strings.Builder

	for i, val := range c.stacks {
		if i == 0 {
			continue
		}
		sb.WriteByte(val[len(val)-1])
	}

	return sb.String()
}
