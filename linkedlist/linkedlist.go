package linkedlist

type Element struct {
	val  int
	next *Element // We're done when next is nil
}

func (e *Element) Get() int {
	return e.val
}

func (e *Element) Next() *Element {
	return e.next
}

func LinkedListFromSlice(s []int) *Element {
	var head *Element
	var prev *Element

	for _, x := range s {
		e := &Element{val: x}
		if prev == nil {
			head, prev = e, e
			continue
		}
		prev.next = e
		prev = e
	}

	return head
}

func Reverse(e *Element) *Element {
	// This will be nil for the first element of the new list, thus setting its
	// next pointer to nil, marking it as the end.
	var prev *Element

	for {
		if e.Next() == nil {
			// This is the last element. Thus its our new first element.
			return &Element{val: e.Get(), next: prev}
		}

		// This is either the first element or a boring old middle element
		prev = &Element{val: e.Get(), next: prev}

		e = e.Next()
	}
}

func RecurseReverse(e *Element) *Element {
	return recurseReverse(e, nil)
}

func recurseReverse(e *Element, prev *Element) *Element {
	if e.Next() == nil {
		// This is the last element. Thus its our new first element.
		return &Element{val: e.Get(), next: prev}
	}

	// This is either the first element or a boring old middle element
	return recurseReverse(e.Next(), &Element{val: e.Get(), next: prev})
}

func set(e *Element) map[int]bool {
	s := make(map[int]bool)
	if e == nil {
		return s
	}
	for e != nil {
		s[e.Get()] = true
		e = e.Next()
	}
	return s
}

func Intersection(i, j *Element) *Element {
	// This feels like cheating... Try merge sort the lists instead?
	seen := set(i)

	var prev *Element

	for j != nil {
		if seen[j.Get()] {
			prev = &Element{val: j.Get(), next: prev}
		}
		j = j.Next()
	}

	return prev
}
