package main

import (
	"image"
	"sync"
)

type corners struct {
	ne, nw, se, sw image.Point
}

var (
	north image.Point = image.Pt(0, -1)
	south image.Point = image.Pt(0, 1)
	east  image.Point = image.Pt(1, 0)
	west  image.Point = image.Pt(-1, 0)
)

type Forest struct {
	trees  map[image.Point]int // don't miss
	corner corners
}

func NewForest(input []string) *Forest {
	f := &Forest{
		trees: make(map[image.Point]int),
		corner: corners{
			nw: image.Pt(0, 0),
			ne: image.Pt(len(input[0])-1, 0),
			sw: image.Pt(0, len(input)-1),
			se: image.Pt(len(input[0])-1, len(input)-1),
		},
	}

	for y, l := range input {
		for x, c := range l {
			pt := image.Pt(x, y)
			f.trees[pt] = int(c - '0')
		}
	}

	return f
}

func (f *Forest) VisibleTrees() int {
	hits := make(chan image.Point)

	seen := make(map[image.Point]bool)

	var wg sync.WaitGroup

	walk := func(start image.Point, next image.Point, hits chan<- image.Point) {
		max := -1
		for p := start; ; p = p.Add(next) {
			ht, ok := f.trees[p]
			if !ok {
				return
			}
			if ht > max {
				hits <- p
				max = ht
			}
		}
	}

	scan := func(q []image.Point, next image.Point, hits chan<- image.Point) {
		defer wg.Done()
		for _, p := range q {
			walk(p, next, hits)
		}
	}

	go func() {
		wg.Wait()
		close(hits)
	}()

	wg.Add(4)
	go scan(getLine(f.corner.nw, f.corner.ne, east), south, hits)
	go scan(getLine(f.corner.ne, f.corner.se, south), west, hits)
	go scan(getLine(f.corner.se, f.corner.sw, west), north, hits)
	go scan(getLine(f.corner.sw, f.corner.nw, north), east, hits)

	for p := range hits {
		seen[p] = true
	}

	return len(seen)
}

func getLine(start image.Point, end image.Point, v image.Point) []image.Point {
	l := make([]image.Point, 0)
	for pt := start; ; pt = pt.Add(v) {
		l = append(l, pt)
		if pt == end {
			break
		}
	}

	return l
}

func (f *Forest) ScenicScore() int {
	var wg sync.WaitGroup

	wg.Add(len(f.trees))

	scores := make(chan int)

	go func() {
		wg.Wait()
		close(scores)
	}()

	for p := range f.trees {
		go func(p image.Point) {
			defer wg.Done()
			scores <- f.scenicScore(p)
		}(p)
	}

	result := -1
	for v := range scores {
		if v > result {
			result = v
		}
	}

	return result
}

func (f *Forest) scenicScore(p image.Point) int {
	directions := []image.Point{north, east, south, west}

	ht := f.trees[p]

	acc := 1
	for _, d := range directions {
		seen := 0
		for n := p.Add(d); ; n = n.Add(d) {
			nHt, ok := f.trees[n]
			if !ok {
				break
			}

			seen++
			if nHt >= ht {
				break
			}
		}
		acc *= seen
	}
	return acc
}
