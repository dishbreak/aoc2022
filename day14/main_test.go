package main

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

/*

  4     5  5
  9     0  0
  4     0  3
0 ......+...
1 ..........
2 ..........
3 ..........
4 ....#...##
5 ....#...#.
6 ..###...#.
7 ........#.
8 ........#.
9 #########.

*/

func TestParseCave(t *testing.T) {
	c := parseCave(input)
	assert.True(t, c[image.Pt(498, 5)])
	assert.False(t, c[image.Pt(500, 5)])
	assert.True(t, c[image.Pt(501, 9)])
}
