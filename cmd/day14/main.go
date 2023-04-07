package main

import (
	"errors"
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day14.txt")
	if err != nil {
		panic(err)
	}

	input = input[:len(input)-1]

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func (c *Cave) SandComesToRest() bool {
	o := image.Pt(500, 0)
	for {
		if c.pts[o] {
			return false
		}
		if o.Y >= c.maxY {
			if c.hasFloor {
				c.pts[o] = true
				return true
			}
			return false
		}
		switch {
		case !c.pts[o.Add(south)]:
			o = o.Add(south)
		case !c.pts[o.Add(southwest)]:
			o = o.Add(southwest)
		case !c.pts[o.Add(southeast)]:
			o = o.Add(southeast)
		default:
			c.pts[o] = true
			return true
		}
	}
}

func (c *Cave) FillWithSand() int {
	i := 0
	for ; c.SandComesToRest(); i++ {
	}
	return i
}

func part1(input []string) int {
	c := parseCave(input)
	return c.FillWithSand()
}

func part2(input []string) int {
	c := parseCave(input, WithInfiniteFloor)
	return c.FillWithSand()
}

type Cave struct {
	pts      map[image.Point]bool
	maxY     int
	hasFloor bool
}

type CaveOption func(*Cave)

func WithInfiniteFloor(c *Cave) {
	c.maxY++
	c.hasFloor = true
}

func parseCave(lines []string, opts ...CaveOption) *Cave {
	cave := &Cave{
		pts:  make(map[image.Point]bool),
		maxY: -1,
	}

	vecs := make([][]image.Point, len(lines))

	for n, line := range lines {
		parts := strings.Split(line, " -> ")
		pts := make([]image.Point, len(parts))

		for i, part := range parts {
			strs := strings.Split(part, ",")
			y, _ := strconv.Atoi(strs[1])
			if y > cave.maxY {
				cave.maxY = y
			}
			x, _ := strconv.Atoi(strs[0])
			pts[i] = image.Pt(x, y)
		}

		vecs[n] = pts
	}

	for _, v := range vecs {
		if len(v) < 2 {
			continue
		}

		for i := 1; i < len(v); i++ {
			d := unitVec(v[i].Sub(v[i-1]))
			for c := v[i-1]; c != v[i]; c = c.Add(d) {
				cave.pts[c] = true
			}
		}

		cave.pts[v[len(v)-1]] = true
	}

	for _, opt := range opts {
		opt(cave)
	}

	return cave
}

var (
	north     image.Point = image.Pt(0, -1)
	south                 = image.Pt(0, 1)
	east                  = image.Pt(1, 0)
	west                  = image.Pt(-1, 0)
	southwest             = south.Add(west)
	southeast             = south.Add(east)
)

func unitVec(v image.Point) image.Point {
	if v.X == 0 && v.Y == 0 {
		panic(errors.New("invalid zero vector"))
	}

	if v.X == 0 {
		if v.Y < 0 {
			return north
		}
		return south
	}
	if v.X < 0 {
		return west
	} else {
		return east
	}
}
