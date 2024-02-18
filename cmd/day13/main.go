package main

import (
	"bufio"
	"strconv"
)

func parse(i *bufio.Scanner) (p []interface{}) {
	var buf int
	var stored bool
	var stop bool
	for !stop && i.Scan() {
		c := i.Text()
		switch c {
		case "[":
			p = append(p, parse(i))
		case "]":
			stop = true
		case ",":
			if stored {
				stored = false
				p = append(p, buf)
				buf = 0
			}
		default:
			stored = true
			digit, _ := strconv.Atoi(i.Text())
			buf = 10*buf + digit
		}
	}

	if stored {
		stored = false
		p = append(p, buf)
		buf = 0
	}

	return p
}

const (
	InOrder    = 1
	OutOfOrder = -1
	Continue   = 0
)

func inOrder(a, b interface{}) int {
	aVal, aInt := a.(int)
	bVal, bInt := b.(int)

	if aInt && bInt {
		switch {
		case aVal < bVal:
			return InOrder
		case aVal == bVal:
			return Continue
		case aVal > bVal:
			return OutOfOrder
		}
	}

	aSlc, aOk := a.([]interface{})
	bSlc, bOk := b.([]interface{})

	if !aOk {
		aSlc = []interface{}{aVal}
	}

	if !bOk {
		bSlc = []interface{}{bVal}
	}

	ct := len(aSlc)
	if bLen := len(bSlc); bLen < ct {
		ct = bLen
	}

	for i := 0; i < ct; i++ {
		result := inOrder(aSlc[i], bSlc[i])
		if result != Continue {
			return result
		}
	}

	// if we got here, we were dealing with two lists
	// and no item pair in the lists made a decision.

	if aLen, bLen := len(bSlc), len(aSlc); aLen < bLen { // b ran out of items first
		return OutOfOrder
	} else if aLen > bLen {
		return InOrder
	}

	return Continue // unclear, keep checking

}
