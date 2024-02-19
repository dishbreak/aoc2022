package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	input := "[[4,47],4,4,4]"
	r := strings.NewReader(input)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	s.Scan()

	result := parse(s)
	assert.Equal(t, 47, result[0].([]interface{})[1].(int))
	assert.Equal(t, 4, result[2].(int))
}

func TestInOrder(t *testing.T) {
	type testCase struct {
		a, b   string
		result int
	}

	testCases := []testCase{
		// == Pair 1 ==
		// - Compare [1,1,3,1,1] vs [1,1,5,1,1]
		//   - Compare 1 vs 1
		//   - Compare 1 vs 1
		//   - Compare 3 vs 5
		//     - Left side is smaller, so inputs are in the right order
		{
			a:      "[1,1,3,1,1]",
			b:      "[1,1,5,1,1]",
			result: InOrder,
		},
		// == Pair 2 ==
		// - Compare [[1],[2,3,4]] vs [[1],4]
		//   - Compare [1] vs [1]
		//     - Compare 1 vs 1
		//   - Compare [2,3,4] vs 4
		//     - Mixed types; convert right to [4] and retry comparison
		//     - Compare [2,3,4] vs [4]
		//       - Compare 2 vs 4
		//         - Left side is smaller, so inputs are in the right order
		{
			a:      "[[1],[2,3,4]]",
			b:      "[[1],4]",
			result: InOrder,
		},
		// == Pair 3 ==
		// - Compare [9] vs [[8,7,6]]
		//   - Compare 9 vs [8,7,6]
		//     - Mixed types; convert left to [9] and retry comparison
		//     - Compare [9] vs [8,7,6]
		//       - Compare 9 vs 8
		//         - Right side is smaller, so inputs are not in the right order
		{
			a:      "[9]",
			b:      "[[8,7,6]]",
			result: OutOfOrder,
		},
		// == Pair 4 ==
		// - Compare [[4,4],4,4] vs [[4,4],4,4,4]
		//   - Compare [4,4] vs [4,4]
		//     - Compare 4 vs 4
		//     - Compare 4 vs 4
		//   - Compare 4 vs 4
		//   - Compare 4 vs 4
		//   - Left side ran out of items, so inputs are in the right order
		{
			a:      "[[4,4],4,4]",
			b:      "[[4,4],4,4,4]",
			result: InOrder,
		},
		// == Pair 5 ==
		// - Compare [7,7,7,7] vs [7,7,7]
		//   - Compare 7 vs 7
		//   - Compare 7 vs 7
		//   - Compare 7 vs 7
		//   - Right side ran out of items, so inputs are not in the right order
		{
			a:      "[7,7,7,7]",
			b:      "[7,7,7]",
			result: OutOfOrder,
		},

		// == Pair 6 ==
		// - Compare [] vs [3]
		//   - Left side ran out of items, so inputs are in the right order
		{
			a:      "[]",
			b:      "[3]",
			result: InOrder,
		},
		// == Pair 7 ==
		// - Compare [[[]]] vs [[]]
		//   - Compare [[]] vs []
		//     - Right side ran out of items, so inputs are not in the right order
		{
			a:      "[[[]]]",
			b:      "[[]]",
			result: OutOfOrder,
		},
		// == Pair 8 ==
		// - Compare [1,[2,[3,[4,[5,6,7]]]],8,9] vs [1,[2,[3,[4,[5,6,0]]]],8,9]
		//   - Compare 1 vs 1
		//   - Compare [2,[3,[4,[5,6,7]]]] vs [2,[3,[4,[5,6,0]]]]
		//     - Compare 2 vs 2
		//     - Compare [3,[4,[5,6,7]]] vs [3,[4,[5,6,0]]]
		//       - Compare 3 vs 3
		//       - Compare [4,[5,6,7]] vs [4,[5,6,0]]
		//         - Compare 4 vs 4
		//         - Compare [5,6,7] vs [5,6,0]
		//           - Compare 5 vs 5
		//           - Compare 6 vs 6
		//           - Compare 7 vs 0
		//             - Right side is smaller, so inputs are not in the right order
		{
			a:      "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			b:      "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			result: OutOfOrder,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case", i), func(t *testing.T) {
			result := inOrder(load(tc.a), load(tc.b))
			assert.Equal(t, tc.result, result)
		})
	}
}

func load(s string) interface{} {
	r := strings.NewReader(s)
	i := bufio.NewScanner(r)
	i.Split(bufio.ScanRunes)

	return parse(i)
}

func TestIsDivisor(t *testing.T) {
	type testCase struct {
		input  string
		result bool
	}

	testCases := []testCase{
		{
			input:  "[[2]]",
			result: true,
		},
		{
			input:  "7",
			result: false,
		},
		{
			input:  "[5,6]",
			result: false,
		},
		{
			input:  "[[2,4]]",
			result: false,
		},
		{
			input:  "[[[3]]]",
			result: false,
		},
		{
			input:  "[[3]]",
			result: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case", i), func(t *testing.T) {
			pkt := load(tc.input)
			assert.Equal(t, tc.result, isDivider(pkt, 2))
		})
	}
}

func TestPart2(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`
	r := strings.NewReader(input)
	assert.Equal(t, 140, part2(r))
}
