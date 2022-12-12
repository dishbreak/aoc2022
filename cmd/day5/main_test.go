package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingCrates(t *testing.T) {
	var input = []string{
		"[D]        ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}
	c := CrateLayoutFromStringSlice(input)
	assert.Equal(t, []byte{'Z', 'N', 'D'}, c.stacks[1])
}

func TestMove(t *testing.T) {
	var input = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}
	c := CrateLayoutFromStringSlice(input)
	c.Move(1, 2, 1)
	c.Move(3, 1, 3)
	c.Move(2, 2, 1)
	c.Move(1, 1, 2)

	assert.Equal(t, "PDNZ", string(c.stacks[3]))
	assert.Equal(t, "CMZ", c.Topline())
}

func TestPart1(t *testing.T) {
	input := [][]string{
		{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
		},
		{
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		},
	}

	assert.Equal(t, "CMZ", part1(input))
}
