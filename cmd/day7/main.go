package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
}

func part1(input io.Reader) int {
	s := bufio.NewScanner(input)
	hits := make(chan int)
	go recurse(s, hits)

	acc := 0
	for v := range hits {
		acc += v
	}

	return acc
}

func recurse(s *bufio.Scanner, hits chan<- int) int {
	size := 0
	defer func() {
		if size <= 100000 {
			hits <- size
		}
	}()

	isRoot := false

	for s.Scan() {
		l := s.Text()
		p := strings.Fields(l)
		switch p[0] {
		case "$":
			if p[1] == "ls" {
				continue
			}

			if p[2] == "/" {
				isRoot = true
				continue
			}

			if p[2] == ".." {
				return size
			}

			size += recurse(s, hits)
		case "dir":
			continue
		default:
			fsize, err := strconv.Atoi(p[0])
			if err != nil {
				panic(err)
			}
			size += fsize
		}
	}

	if isRoot {
		close(hits)
	}

	return size
}
