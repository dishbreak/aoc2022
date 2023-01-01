package main

import "image"

type Grid struct {
	start, end image.Point
	space      map[image.Point]int
}

func GridFromInput(input []string) *Grid {
	g := &Grid{
		space: make(map[image.Point]int),
	}

	for y, l := range input {
		for x, c := range l {
			pt := image.Pt(x, y)
			elev := 0
			switch c {
			case 'E':
				elev = 25
				g.end = pt
			case 'S':
				elev = 0
				g.start = pt
			default:
				elev = int(c - 'a')
			}

			g.space[pt] = elev
		}
	}

	return g
}

var neighbors = []image.Point{
	image.Pt(1, 0),
	image.Pt(-1, 0),
	image.Pt(0, -1),
	image.Pt(0, 1),
}

func (g *Grid) ReachStartPoint(p image.Point) bool {
	return p == g.start
}

func (g *Grid) ReachZeroElevation(p image.Point) bool {
	return g.space[p] == 0
}

func (g *Grid) ShortestPathTo(reached func(image.Point) bool) int {
	//borrowed from https://observablehq.com/@jwolondon/advent-of-code-2022-day-12
	dists := make(map[image.Point]int)

	type frame struct {
		pt   image.Point
		dist int
	}

	q := make([]frame, 1)

	q[0] = frame{
		pt:   g.end,
		dist: 0,
	}

	for len(q) > 0 {
		f := q[0]
		q = q[1:]

		if reached(f.pt) {
			return f.dist
		}

		for _, n := range neighbors {
			p := f.pt.Add(n)
			z := g.space[f.pt]

			z0, ok := g.space[p]
			if !ok {
				continue
			}

			if z0-z < -1 {
				continue
			}

			if _, ok := dists[p]; ok {
				continue
			}

			dists[p] = f.dist + 1
			q = append(q, frame{p, dists[p]})
		}
	}
	return -1
}
