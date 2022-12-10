package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, io.SeekStart)
	fmt.Printf("Part 2: %d\n", part2(f))

}

func part1(input io.Reader) int {
	s := bufio.NewScanner(input)

	max, acc := -1, 0
	for s.Scan() {
		l := s.Text()
		if l == "" {
			if max < acc {
				max = acc
			}
			acc = 0
			continue
		}

		n, _ := strconv.Atoi(l)
		acc += n
	}

	if acc != 0 && max < acc {
		max = acc
	}

	return max
}

func part2(input io.Reader) int {
	s := bufio.NewScanner(input)

	acc := 0
	cts := make([]int, 0)
	for s.Scan() {
		l := s.Text()
		if l == "" {
			cts = append(cts, acc)
			acc = 0
			continue
		}

		n, _ := strconv.Atoi(l)
		acc += n
	}

	if acc != 0 {
		cts = append(cts, acc)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cts)))

	return cts[0] + cts[1] + cts[2]
}
