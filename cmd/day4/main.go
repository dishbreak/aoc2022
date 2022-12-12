package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, io.SeekStart)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	acc := 0
	s := bufio.NewScanner(r)

	for s.Scan() {
		l := s.Text()
		if nestingRanges(l) {
			acc++
		}
	}

	return acc
}

func part2(r io.Reader) int {
	acc := 0
	s := bufio.NewScanner(r)

	for s.Scan() {
		l := s.Text()
		if overlappingRanges(l) {
			acc++
		}
	}

	return acc
}

var rangeExp = regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

type Range struct {
	min, max int
}

func parseRanges(s string) (one, other Range) {
	m := rangeExp.FindStringSubmatch(s)
	one.min, _ = strconv.Atoi(m[1])
	one.max, _ = strconv.Atoi(m[2])
	other.min, _ = strconv.Atoi(m[3])
	other.max, _ = strconv.Atoi(m[4])
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nestingRanges(s string) bool {
	one, other := parseRanges(s)
	return one.contains(other) || other.contains(one)
}

func overlappingRanges(s string) bool {
	one, other := parseRanges(s)

	return max(one.min, other.min) <= min(one.max, other.max)
}

func (r Range) contains(o Range) bool {
	return o.min >= r.min && o.max <= r.max
}
