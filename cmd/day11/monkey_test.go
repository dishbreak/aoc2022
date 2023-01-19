package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func getInputAsSections() [][]string {
	input := make([][]string, 0)
	for _, b := range strings.Split(testInput, "\n\n") {
		input = append(input, strings.Split(b, "\n"))
	}

	return input
}

func monkeyFixture() []*Monkey {
	input := getInputAsSections()

	monkeys := make([]*Monkey, 4)

	for i := range monkeys {
		monkeys[i] = MonkeyFromInput(input[i])
	}

	return monkeys
}
func TestParseMonkey(t *testing.T) {
	monkeys := monkeyFixture()

	assert.Equal(t, []int{54, 65, 75, 74}, monkeys[1].items)
	assert.Equal(t, 38, monkeys[0].op(2))
	assert.Equal(t, 76, monkeys[1].op(70))
	assert.Equal(t, 16, monkeys[2].op(4))
	assert.Equal(t, 13, monkeys[2].modulo)
	assert.Equal(t, 0, monkeys[3].trueIdx)
	assert.Equal(t, 3, monkeys[2].falseIdx)
}
