package main

import (
	"errors"
	"image"
	"strconv"
	"strings"
)

func main() {

}

type Cave struct {
	pts  map[image.Point]bool
	maxY int
}

func parseCave(lines []string) Cave {
	cave := Cave{
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

	return cave
}

var (
	north     image.Point = image.Pt(0, 1)
	south                 = image.Pt(0, -1)
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
			return south
		}
		return north
	}
	if v.X < 0 {
		return west
	} else {
		return east
	}
}
