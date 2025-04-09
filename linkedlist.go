package linkedlist

import (
	"iter"
)

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	len  int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0}
}

func NewWithSlice[T any](slice []T) *LinkedList[T] {
	l := New[T]()

	for i := 0; i < len(slice); i++ {
		l.PushBack(&slice[i])
	}

	return l
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) PushBack(value *T) {
	n := newNode(value)

	if l.tail != nil {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	} else {
		l.head = n
		l.tail = l.head
	}

	l.len++
}

func (l *LinkedList[T]) PushFront(value *T) {
	n := newNode(value)

	if l.head != nil {
		n.next = l.head
		l.head.prev = n
		l.head = n
	} else {
		l.head = n
		l.tail = l.head
	}

	l.len++
}

func (l *LinkedList[T]) PopBack() (*T, error) {
	if err := checkEmpty(l.len); err != nil {
		return nil, err
	}

	n := l.tail
	if l.tail.prev != nil {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	} else {
		l.head = nil
		l.tail = nil
	}

	l.len--
	return n.value, nil
}

func (l *LinkedList[T]) PopFront() (*T, error) {
	if err := checkEmpty(l.len); err != nil {
		return nil, err
	}

	n := l.head
	if l.head.next != nil {
		l.head.next.prev = nil
		l.head = l.head.next
	} else {
		l.head = nil
		l.tail = nil
	}

	l.len--
	return n.value, nil
}

func (l *LinkedList[T]) Insert(index int, value *T) error {
	if index == 0 {
		l.PushFront(value)
		return nil
	}
	if index == l.len {
		l.PushBack(value)
		return nil
	}
	if err := checkIndex(index, l.len); err != nil {
		return err
	}

	p := l.head
	for _ = range index {
		p = p.next
	}

	n := newNode(value)
	p.prev.next = n
	n.prev = p.prev
	n.next = p
	p.prev = n

	l.len++
	return nil
}

func (l *LinkedList[T]) Get(index int) (*T, error) {
	if err := checkIndex(index, l.len); err != nil {
		return nil, err
	}

	p := l.head
	for _ = range index {
		p = p.next
	}

	return p.value, nil
}

func (l *LinkedList[T]) Remove(index int) (*T, error) {
	if index == 0 {
		value, err := l.PopFront()
		return value, err
	}
	if index == l.len-1 {
		value, err := l.PopBack()
		return value, err
	}
	if err := checkIndex(index, l.len); err != nil {
		return nil, err
	}

	p := l.head
	for _ = range index {
		p = p.next
	}

	p.prev.next = p.next
	p.next.prev = p.prev

	l.len--
	return p.value, nil
}

func (l *LinkedList[T]) Iter() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		p := l.head
		for p != nil {
			if !yield(p.value) {
				return
			}
			p = p.next
		}
	}
}

type node[T any] struct {
	value *T
	next  *node[T]
	prev  *node[T]
}

func newNode[T any](value *T) *node[T] {
	return &node[T]{
		value,
		nil,
		nil,
	}
}
