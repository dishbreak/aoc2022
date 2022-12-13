package main

type CharRegister struct {
	window []byte
	hits   [26]int
}

func NewCharRegister() *CharRegister {
	c := &CharRegister{
		window: make([]byte, 0),
	}

	return c
}

func (c *CharRegister) Add(b byte) {
	c.window = append(c.window, b)
	c.hits[b-'a']++

	if len(c.window) > 4 {
		p := c.window[0]
		c.window = c.window[1:]
		c.hits[p-'a']--
	}
}

func (c *CharRegister) Match() bool {
	if len(c.window) < 4 {
		return false
	}

	for _, ct := range c.hits {
		if ct > 1 {
			return false
		}
	}
	return true
}
