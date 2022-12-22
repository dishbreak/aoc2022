package main

import "image"

type Rope struct {
	head, tail image.Point
	tailSpots  map[image.Point]bool
}

func (r *Rope) Move(v image.Point) {
	r.head = r.head.Add(v)
	r.moveTail(v)
	r.tailSpots[r.tail] = true
}

func (r *Rope) moveTail(v image.Point) {
	if touching(r.head, r.tail) {
		return
	}

	if colinear(r.head, r.tail) {
		r.tail = r.tail.Add(v)
		return
	}

	d := r.head.Sub(r.tail)
	if aX, aY := abs(d.X), abs(d.Y); aX > 1 {
		d.X = d.X / aX
	} else {
		d.Y = d.Y / aY
	}
	r.tail = r.tail.Add(d)
}

func (r *Rope) TailHits() int {
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
