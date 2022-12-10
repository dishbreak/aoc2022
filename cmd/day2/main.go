package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, io.SeekStart)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	// The winner of the whole tournament is the player with the highest score.
	// Your total score is the sum of your scores for each round. The score for
	// a single round is the score for the shape you selected (1 for Rock, 2 for
	// Paper, and 3 for Scissors) plus the score for the outcome of the round (0
	// if you lost, 3 if the round was a draw, and 6 if you won).

	g := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 6,
			"Z": 0,
		},
		"B": {
			"X": 0,
			"Y": 3,
			"Z": 6,
		},
		"C": {
			"X": 6,
			"Y": 0,
			"Z": 3,
		},
	}

	pts := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	s := bufio.NewScanner(r)
	acc := 0
	for s.Scan() {
		l := s.Text()
		parts := strings.Fields(l)
		acc += pts[parts[1]]
		acc += g[parts[0]][parts[1]]
	}

	return acc
}

func part2(r io.Reader) int {
	score := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	shapeBonus := map[string]map[string]int{
		"A": {
			"X": 3,
			"Y": 1,
			"Z": 2,
		},
		"B": {
			"X": 1,
			"Y": 2,
			"Z": 3,
		},
		"C": {
			"X": 2,
			"Y": 3,
			"Z": 1,
		},
	}

	acc := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		parts := strings.Fields(l)
		acc += score[parts[1]]
		acc += shapeBonus[parts[0]][parts[1]]
	}

	return acc
}
