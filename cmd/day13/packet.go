package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Packet struct {
	full  bool
	value int
	items []*Packet
}

func PacketFromString(s string) *Packet {
	r := strings.NewReader(s)
	r.ReadByte() // read the leading brace.
	return parsePacket(r)
}

func parsePacket(r io.ByteReader) *Packet {
	p := &Packet{
		items: make([]*Packet, 0),
	}

	var buf *Packet

	for b, err := r.ReadByte(); err == nil; b, err = r.ReadByte() {
		switch b {
		case ']':
			if buf != nil {
				p.items = append(p.items, buf)
				p.full = true
			}
			return p
		case ',':
			p.items = append(p.items, buf)
			p.full = true
			buf = nil
		case '[':
			buf = parsePacket(r)
		default: // numeric
			if buf == nil {
				buf = &Packet{
					full: true,
				}
			}
			digit := int(b - '0')
			buf.value = (buf.value * 10) + digit
		}
	}
	panic(errors.New("unexpected EOF"))
}

func (p *Packet) String() string {
	if !p.full {
		return "[]"
	}

	if len(p.items) == 0 {
		return strconv.Itoa(p.value)
	}

	pts := make([]string, len(p.items))
	for i, item := range p.items {
		pts[i] = item.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(pts, ","))
}

func (p *Packet) IsInt() bool {
	return p.full && len(p.items) == 0
}

func (p *Packet) ToListPacket() *Packet {
	return &Packet{
		items: []*Packet{p},
	}
}

func LessThan(left, right *Packet) bool {
	leftInt, rightInt := left.IsInt(), right.IsInt()

	if leftInt && rightInt {
		return left.value < right.value
	}

	if leftInt {
		left = left.ToListPacket()
	}

	if rightInt {
		right = right.ToListPacket()
	}

	length := len(left.items)
	if rLen := len(right.items); rLen < len(left.items) {
		length = rLen
	}

	for i := 0; i < length; i++ {
		if LessThan(left.items[i], right.items[i]) {
			return true
		}
	}

	return len(left.items) < len(right.items)

}
