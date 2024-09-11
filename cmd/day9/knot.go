package main

import "image"

type KnotLink struct {
	head, tail image.Point
	tailSpots  map[image.Point]bool
}

func (r *KnotLink) Move(v image.Point) image.Point {
	r.head = r.head.Add(v)
	n := r.moveTail(v)
	r.tailSpots[r.tail] = true
	return n
}

func (r *KnotLink) moveTail(v image.Point) image.Point {
	if touching(r.head, r.tail) {
		return image.Point{}
	}

	if colinear(r.head, r.tail) {
		r.tail = r.tail.Add(v)
		return v
	}

	d := r.head.Sub(r.tail)
	if aX, aY := abs(d.X), abs(d.Y); aX > 1 {
		d.X = d.X / aX
	} else {
		d.Y = d.Y / aY
	}
	r.tail = r.tail.Add(d)

	return d
}

func (r *KnotLink) TailHits() int {
	return len(r.tailSpots)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func touching(one, other image.Point) bool {
	if d := one.Sub(other); abs(d.X) <= 1 && abs(d.Y) <= 1 {
		return true
	}
	return false
}

func colinear(one, other image.Point) bool {
	if d := one.Sub(other); d.X == 0 || d.Y == 0 {
		return true
	}
	return false
}

type Rope struct {
	knots [10]*KnotLink
}

func NewRope() *Rope {
	r := &Rope{}

	for i := range r.knots {
		r.knots[i] = &KnotLink{
			tailSpots: make(map[image.Point]bool),
		}
	}

	return r
}

func (r *Rope) Move(v image.Point) {
	for _, k := range r.knots {
		v = k.Move(v)
		if v == image.Pt(0, 0) {
			return
		}
	}
}

func (r *Rope) TailHits() int {
	return len(r.knots[9].tailSpots)
}
