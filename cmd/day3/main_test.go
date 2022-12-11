package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonItem(t *testing.T) {
	type testCase struct {
		input  [3]string
		result int
	}

	testCases := []testCase{
		{
			input: [3]string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			result: 18,
		},
		{
			input: [3]string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			result: 52,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.result, commonItem(tc.input))
		})
	}
}
