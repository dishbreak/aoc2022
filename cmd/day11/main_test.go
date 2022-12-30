package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := getInputAsSections()
	assert.Equal(t, int64(10605), part1(input))
}

func TestPart2(t *testing.T) {
	input := getInputAsSections()
	assert.Equal(t, int64(2713310158), part2(input))
}
