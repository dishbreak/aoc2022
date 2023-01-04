package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePacket(t *testing.T) {
	input := "[[5],[1,[[0]]],[],[3,[[9,1],[3,4,10],8,3],6]]"
	p := PacketFromString(input)
	assert.Equal(t, 5, p.items[0].items[0].value)
	assert.Equal(t, input, p.String())
}

func TestLessThan(t *testing.T) {
	type testCase struct {
		left, right string
		result      bool
	}

	testCases := []testCase{
		{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			true,
		},
		{
			"[[1],[2,3,4]]",
			"[[1],4]",
			true,
		},
		{
			"[9]",
			"[[8,7,6]]",
			false,
		},
		{
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
			true,
		},
		{
			"[7,7,7,7]",
			"[7,7,7]",
			false,
		},
		{
			"[]",
			"[3]",
			true,
		},
		{
			"[[[]]]",
			"[[]]",
			false,
		},
		{
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
			false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			pktL, pktR := PacketFromString(tc.left), PacketFromString(tc.right)
			assert.Equal(t, tc.result, LessThan(pktL, pktR))
		})
	}
}
