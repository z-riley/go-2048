package util

// iter is a genetic iterator which can iterate forwards or backwards.
type iter struct {
	length  int
	reverse bool
	idx     int
}

// NewIter constructs a new iterator.
func NewIter(length int, reverse bool) *iter {
	if reverse {
		return &iter{length: length, reverse: true, idx: length - 1}
	} else {
		return &iter{length: length, reverse: false, idx: 0}
	}
}

// HasNext returns true if another iteration is available.
func (i *iter) HasNext() bool {
	if i.reverse {
		return i.idx >= 0
	} else {
		return i.idx < i.length
	}
}

// Next returns the index of the Next index.
func (i *iter) Next() int {
	if !i.HasNext() {
		panic("no more elements")
	}
	if i.reverse {
		out := i.idx
		i.idx--
		return out
	} else {
		out := i.idx
		i.idx++
		return out
	}
}

// Len returns the length of the iterator.
func (i *iter) Len() int {
	return i.length
}
