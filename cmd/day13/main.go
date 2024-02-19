package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)
	var lines [2]string
	var i int
	var acc int
	idx := 1

	for s.Scan() {
		lines[i] = s.Text()
		i++

		if i == 2 {
			i = 0
			left, right := parse(makeScanner(lines[0])), parse(makeScanner(lines[1]))
			if inOrder(left, right) == InOrder {
				acc += idx
			}
			idx++
			s.Scan() // grab the newline
		}
	}
	return acc
}

func makeScanner(s string) *bufio.Scanner {
	r := strings.NewReader(s)
	i := bufio.NewScanner(r)
	i.Split(bufio.ScanRunes)
	return i
}

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

func isDivider(pkt interface{}, divisor int) bool {
	outerSlc, ok := pkt.([]interface{})
	if !ok {
		return false
	}
	if len(outerSlc) != 1 {
		return false
	}

	innerSlc, ok := outerSlc[0].([]interface{})
	if !ok {
		return false
	}
	if len(innerSlc) != 1 {
		return false
	}

	innerSlc, ok = innerSlc[0].([]interface{})
	if !ok {
		return false
	}
	if len(innerSlc) != 1 {
		return false
	}

	innerVal, ok := innerSlc[0].(int)
	if !ok {
		return false
	}

	return innerVal == divisor
}

func part2(r io.Reader) int {
	var pkts []interface{}
	s := bufio.NewScanner(r)

	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		pkts = append(pkts, parse(makeScanner(s.Text())))
	}

	for _, div := range []int{2, 6} {
		pkts = append(pkts, parse(makeScanner(fmt.Sprintf("[[%d]]", div))))
	}

	sort.Slice(pkts, func(i, j int) bool {
		return inOrder(pkts[i], pkts[j]) == InOrder
	})

	divTgt := 2

	result := 1
	for i, pkt := range pkts {
		if isDivider(pkt, divTgt) {
			result *= (i + 1)
			if divTgt == 6 {
				break
			}
			divTgt = 6
		}
	}

	return result
}
