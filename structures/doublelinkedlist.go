package structures

import "fmt"

type IntDoublelinkedListNode struct {
	Value int
	Next,Prev  *IntDoublelinkedListNode
}

type IntDoublelinkedList struct {
	head *IntDoublelinkedListNode
	last *IntDoublelinkedListNode
}

func (s *IntDoublelinkedList) addFirst(element int) {
	if s.head == nil {
		s.head = &IntDoublelinkedListNode{Value:element}
		s.last = s.head
	} else {
		newNode := &IntDoublelinkedListNode{Value:element, Next:s.head}
		s.head.Prev = newNode
		s.head = newNode
	}
}

func (s *IntDoublelinkedList) addLast(element int) {
	if s.last == nil {
		s.addFirst(element)
	} else {
		newNode := &IntDoublelinkedListNode{Value:element, Prev:s.last}
		s.last.Next = newNode
		s.last = newNode
	}
}

func (s *IntDoublelinkedList) addAfter(element,n int) error {
	if n < 0 {
		panic("wrong N")
	}
	i := 1
	curr := s.getHead()
	for {
		if n == i {
			if !curr.hasNext() {
				s.addLast(element)
				return nil
			}
			newNode := &IntDoublelinkedListNode{Value:element, Prev:curr}
			curr.Next = newNode
			newNode.Next = curr.getNext()
			curr.getNext().Prev = newNode
			return nil
		}
		i++
		if !curr.hasNext() {
			return fmt.Errorf("out of list")
		}
		curr = curr.getNext()
	}
}

func (s *IntDoublelinkedList) getHead() *IntDoublelinkedListNode {
	return s.head
}

func (s *IntDoublelinkedList) getLast() *IntDoublelinkedListNode {
	return s.last
}

func (s *IntDoublelinkedList) delAll() {
	s.head = nil
	s.last = nil
}

func (s *IntDoublelinkedListNode) hasNext() bool {
	if s.Next == nil {
		return false
	}
	return true
}

func (s *IntDoublelinkedListNode) hasPrev() bool {
	if s.Prev == nil {
		return false
	}
	return true
}

func (s *IntDoublelinkedListNode) getNext() *IntDoublelinkedListNode {
	if !s.hasNext() {
		panic("use hasNext before using getNext")
	}
	return s.Next
}

func (s *IntDoublelinkedListNode) getPrev() *IntDoublelinkedListNode {
	if !s.hasPrev() {
		panic("use hasPrev before using getPrev")
	}
	return s.Prev
}
