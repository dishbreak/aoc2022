package main

type CharRegister struct {
	window []byte
	hits   [26]int
	size   int
}

type CharRegisterOption func(*CharRegister)

func WithWindowSize(size int) CharRegisterOption {
	return func(cr *CharRegister) {
		cr.size = size
	}
}

func NewCharRegister(opts ...CharRegisterOption) *CharRegister {
	c := &CharRegister{
		window: make([]byte, 0),
		size:   4,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *CharRegister) Add(b byte) {
	c.window = append(c.window, b)
	c.hits[b-'a']++

	if len(c.window) > c.size {
		p := c.window[0]
		c.window = c.window[1:]
		c.hits[p-'a']--
	}
}

func (c *CharRegister) Match() bool {
	if len(c.window) < c.size {
		return false
	}

	for _, ct := range c.hits {
		if ct > 1 {
			return false
		}
	}
	return true
}
