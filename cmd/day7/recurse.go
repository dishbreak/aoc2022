package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func getDirSizes(s *bufio.Scanner) []int {
	results := make([]int, 0)
	hits := make(chan int)

	go recurse(s, hits)

	for v := range hits {
		results = append(results, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(results)))

	return results
}

func recurse(s *bufio.Scanner, hits chan<- int) int {
	size := 0
	isRoot := false

	defer func() {
		hits <- size
		if isRoot {
			close(hits)
		}
	}()

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

	return size
}
