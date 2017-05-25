package list

type Node struct {
	Prev  *Node
	Next  *Node
	Value uint32
}

type List struct {
	Head *Node
	Tail *Node
}

func NewFromSlice(s []uint32) *List {
	l := &List{}
	for _, i := range s {
		l.Append(i)
	}
	return l
}

func (l *List) Append(v uint32) *Node {
	n := &Node{Value: v}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return n
	}
	n.Prev = l.Tail
	l.Tail.Next = n
	l.Tail = n
	return n
}

func (l *List) MoveToTail(n *Node) {
	if n == l.Tail {
		return
	}
	if n == l.Head {
		l.Head = n.Next
	}
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	n.Prev = l.Tail
	l.Tail.Next = n
	l.Tail = n
	l.Tail.Next = nil
}

func (l *List) Delete(n *Node) {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	if n == l.Head {
		l.Head = n.Next
	}
	if n == l.Tail {
		l.Tail = n.Prev
	}
}

func (l *List) TrimHead() *Node {
	if l.Head == l.Tail {
		h := l.Head
		l.Head = nil
		l.Tail = nil
		return h
	}
	h := l.Head
	l.Head = l.Head.Next
	l.Head.Prev = nil
	return h
}

func (l *List) ToSlice() []uint32 {
	s := []uint32{}
	n := l.Head
	for n != nil {
		s = append(s, n.Value)
		n = n.Next
	}
	return s
}
