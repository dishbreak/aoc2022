package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day7.txt")
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

	dirSizes := getDirSizes(s)

	acc := 0
	for _, v := range dirSizes {
		if v > 100000 {
			continue
		}
		acc += v
	}

	return acc
}

const totalSize int = 70000000
const minUnused int = 30000000

func part2(input io.Reader) int {
	s := bufio.NewScanner(input)

	sizes := getDirSizes(s)

	inUse := sizes[0]

	unused := totalSize - inUse
	target := minUnused - unused

	for i, size := range sizes {
		if i == 0 {
			continue
		}

		if size < target {
			return sizes[i-1]
		}
	}

	return -1
}
