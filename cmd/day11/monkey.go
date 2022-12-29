package main

import (
	"strconv"
	"strings"
)

type operator func(int64) int64

func Square() operator {
	return func(i int64) int64 {
		return i * i
	}
}

func AddWith(operand int64) operator {
	return func(i int64) int64 {
		return i + operand
	}
}

func MultBy(operand int64) operator {
	return func(i int64) int64 {
		return i * operand
	}
}

type Monkey struct {
	items     []int64
	op        operator
	ifTrue    *Monkey
	ifFalse   *Monkey
	trueIdx   int
	falseIdx  int
	modulo    int64
	inspected int64
	worried   bool
}

type MonkeyOption func(*Monkey)

func WithNoCalming() MonkeyOption {
	return func(m *Monkey) {
		m.worried = true
	}
}

func (m *Monkey) Catches(item int64) {
	m.items = append(m.items, item)
}

func (m *Monkey) Inspect() {
	for len(m.items) > 0 {
		m.inspected++
		i := m.items[0]
		m.items = m.items[1:]

		i = m.op(i)

		// what, me, worry?
		if !m.worried {
			i = i / 3
		}

		if i%m.modulo == 0 {
			m.ifTrue.Catches(i)
			continue
		}
		m.ifFalse.Catches(i)
	}
}

func (m *Monkey) Connect(monkeys []*Monkey) {
	m.ifTrue = monkeys[m.trueIdx]
	m.ifFalse = monkeys[m.falseIdx]
}

func toOperator(input string) operator {
	if strings.HasSuffix(input, "old * old") {
		return Square()
	}

	parts := strings.Fields(input)
	scratch, _ := strconv.Atoi(parts[5])
	operand := int64(scratch)
	switch parts[4] {
	case "*":
		return MultBy(operand)
	default:
		return AddWith(operand)
	}
}

func toModulo(input string) int64 {
	parts := strings.Fields(input)
	modulo, _ := strconv.Atoi(parts[3])
	return int64(modulo)
}

func toThrowIdx(input string) int {
	parts := strings.Fields(input)
	throwIdx, _ := strconv.Atoi(parts[5])
	return throwIdx
}

func MonkeyFromInput(input []string, opts ...MonkeyOption) *Monkey {
	m := &Monkey{
		items: make([]int64, int64(0)),
	}

	// parse items
	// '  Starting items: 61, 85'
	for _, s := range strings.Split(strings.TrimPrefix(input[1], "  Starting items: "), ", ") {
		i, _ := strconv.Atoi(s)
		m.items = append(m.items, int64(i))
	}

	// parse operation
	m.op = toOperator(input[2])

	// parse modulo
	m.modulo = toModulo(input[3])

	// get true/false Idx
	m.trueIdx = toThrowIdx(input[4])
	m.falseIdx = toThrowIdx(input[5])

	for _, opt := range opts {
		opt(m)
	}

	return m
}
