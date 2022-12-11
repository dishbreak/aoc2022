package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

func main() {
	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, io.SeekStart)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)

	acc := 0
	for s.Scan() {
		l := s.Text()
		acc += getRucksackPriority(l)
	}

	return acc
}

func toPriority(r rune) int {
	if r <= 'Z' {
		return int(r-'A') + 27
	}
	return int(r-'a') + 1
}

func getRucksackPriority(input string) int {
	compartment := []string{
		input[:len(input)/2],
		input[len(input)/2:],
	}

	hits := make(map[rune]bool)
	for _, b := range compartment[0] {
		hits[b] = true
	}

	for _, b := range compartment[1] {
		if hits[b] {
			return toPriority(b)
		}
	}
	panic(fmt.Errorf("found no overlap in compartments for %s", input))
}

func part2(r io.Reader) int {
	triplets := toTriplets(r)
	results := findBadge(triplets)

	acc := 0
	for n := range results {
		acc += n
	}

	return acc
}

func toTriplets(r io.Reader) <-chan [3]string {
	output := make(chan [3]string)

	go func() {
		defer close(output)

		var buf [3]string
		s := bufio.NewScanner(r)

		for i := 0; s.Scan(); i++ {
			l := s.Text()
			n := i % 3
			buf[n] = l
			if n == 2 {
				output <- buf
				buf = [3]string{}
			}
		}
	}()

	return output
}

func findBadge(input <-chan [3]string) chan int {
	output := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU()-1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for triplet := range input {
				output <- commonItem(triplet)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func commonItem(triplet [3]string) int {
	// encode each string as a bitvector
	var bv [3]int64
	for i, s := range triplet {
		bv[i] = toBitvector(s)
	}

	// get the intersection of the three using logical and
	r := bv[0] & bv[1] & bv[2]

	// we'll find the index of the least significant high bit.
	// we assume there's only 1 common element, so we'll stop when we find a bit that's 1.
	// declaring i outside the loop lets us break out of it.
	i := 0
	for ; r > 0; i, r = i+1, r>>1 {
		if r&1 == 1 {
			break
		}
	}

	// if r is 0, we screwed up. return 1.
	if r == 0 {
		return -1
	}

	// the first 26 bits are used to flag A-Z, so add 27 to the bit index to get the priority.
	if i < 26 {
		return i + 27
	}
	// otherwise it's a-z, reduce by 25
	return i - 25
}

func toBitvector(s string) (m int64) {
	// return a bitvector that shows which characters are present in the string
	for _, c := range s {
		m = m | 1<<toIdx(c)
	}
	return
}

func toIdx(r rune) int {
	//bits 0-25 are flags for A-Z
	if r <= 'Z' {
		return int(r - 'A')
	}
	// bit 26 onwards is for a-z
	return int(r-'a') + 26
}
