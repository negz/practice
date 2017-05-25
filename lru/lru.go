package lru

import "github.com/negz/practice/lru/list"

type Cache struct {
	l  *list.List
	kn map[string]*list.Node
	nk map[*list.Node]string
	m  int
}

func New(max int) *Cache {
	return &Cache{
		&list.List{},
		make(map[string]*list.Node),
		make(map[*list.Node]string),
		max,
	}
}

func (c *Cache) Get(k string) (uint32, bool) {
	n := c.kn[k]
	if n == nil {
		return 0, false
	}
	c.l.MoveToTail(n)
	return n.Value, true
}

func (c *Cache) Insert(k string, v uint32) {
	n := c.kn[k]
	if n != nil {
		c.l.MoveToTail(n)
		c.l.Tail.Value = v
		return
	}

	n = c.l.Append(v)
	c.kn[k] = n
	c.nk[n] = k

	if len(c.kn) > c.m {
		h := c.l.TrimHead()
		delete(c.kn, c.nk[h])
		delete(c.nk, h)
	}
}

func (c *Cache) ToSlice() []uint32 {
	return c.l.ToSlice()
}
