package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	testCases := []testCase{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, part1(tc.input))
		})
	}
}

func TestPart2(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	testCases := []testCase{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, part2(tc.input))
		})
	}
}
