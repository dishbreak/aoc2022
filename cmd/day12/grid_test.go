package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFile = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestGetShortestPath(t *testing.T) {
	g := GridFromInput(strings.Split(testFile, "\n"))
	assert.Equal(t, 31, g.ShortestPathTo(g.ReachStartPoint))
	assert.Equal(t, 29, g.ShortestPathTo(g.ReachZeroElevation))
}
