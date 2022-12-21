package main

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewForest(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	f := NewForest(input)
	assert.Equal(t, 5, f.trees[image.Pt(1, 1)])
}

func TestVisibleTrees(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	f := NewForest(input)
	assert.Equal(t, 21, f.VisibleTrees())
}

func TestScenicScore(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	f := NewForest(input)
	assert.Equal(t, 8, f.ScenicScore())
}

func TestScenicScorePoint(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	f := NewForest(input)
	assert.Equal(t, 8, f.scenicScore(image.Pt(2, 3)))
}
