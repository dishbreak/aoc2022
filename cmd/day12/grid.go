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

	// define a frame type to help the BFS
	type frame struct {
		pt   image.Point
		dist int
	}

	// make a queue of frames
	q := make([]frame, 1)

	// seed the queue with the end point -- we'll walk backwards to the start.
	q[0] = frame{
		pt:   g.end,
		dist: 0,
	}

	for len(q) > 0 {
		// pull an item off the queue
		f := q[0]
		q = q[1:]

		// if our callback says we're done, we're done!
		if reached(f.pt) {
			return f.dist
		}

		// for each neighboring point...
		for _, n := range neighbors {
			p := f.pt.Add(n)
			z := g.space[f.pt]

			// attempt to find the neighbor. if it doesn't exist, move on to the next one
			// this happens at our edges often.
			z0, ok := g.space[p]
			if !ok {
				continue
			}

			// if the neighbor doesn't flow into this point, move on to the next one
			// flow means the neighbor's height is at least 1 less than the current ht
			if z0-z < -1 {
				continue
			}

			// if we've already visited, move along.
			if _, ok := dists[p]; ok {
				continue
			}

			// note the distance for the neighbor, then enqueue
			dists[p] = f.dist + 1
			q = append(q, frame{p, dists[p]})
		}
	}

	// if we've exhausted our queue and we didn't reach, the constraint is impossible.
	// -1 signifies there's no solution (but this is AoC so we know this shouldn't happen.)
	return -1
}
