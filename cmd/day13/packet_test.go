package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePacket(t *testing.T) {
	input := "[[5],[1,[[0]]],[],[3,[[9,1],[3,4,10],8,3],6]]"
	p := PacketFromString(input)
	assert.Equal(t, 5, p.items[0].items[0].value)
	assert.Equal(t, input, p.String())
}
