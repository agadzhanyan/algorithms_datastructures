package structures

import "fmt"

type IntSinglelinkedListNode struct {
	Value int
	Next  *IntSinglelinkedListNode
}

type IntSinglelinkedList struct {
	head *IntSinglelinkedListNode
	last *IntSinglelinkedListNode
}

func (s *IntSinglelinkedList) AddFirst(element int) {
	if s.head == nil {
		s.head = &IntSinglelinkedListNode{Value:element}
		s.last = s.head
	} else {
		s.head = &IntSinglelinkedListNode{Value:element, Next:s.head}
	}
}

func (s *IntSinglelinkedList) AddLast(element int) {
	if s.last == nil {
		s.AddFirst(element)
	} else {
		newNode := &IntSinglelinkedListNode{Value:element}
		s.last.Next = newNode
		s.last = newNode
	}
}

func (s *IntSinglelinkedList) AddAfter(element,n int) error {
	if n < 0 {
		panic("wrong N")
	}
	i := 1
	curr := s.GetHead()
	for {
		if n == i {
			newNode := &IntSinglelinkedListNode{Value:element}
			if curr.HasNext() {
				newNode.Next = curr.Next
			}
			curr.Next = newNode
			return nil
		}
		i++
		if !curr.HasNext() {
			return fmt.Errorf("out of list")
		}
		curr = curr.GetNext()
	}
}

func (s *IntSinglelinkedList) GetHead() *IntSinglelinkedListNode {
	return s.head
}

func (s *IntSinglelinkedList) GetLast() *IntSinglelinkedListNode {
	return s.last
}

func (s *IntSinglelinkedList) delAll() {
	s.head = nil
	s.last = nil
}

func (s *IntSinglelinkedListNode) HasNext() bool {
	if s.Next == nil {
		return false
	}
	return true
}

func (s *IntSinglelinkedListNode) GetNext() *IntSinglelinkedListNode {
	if !s.HasNext() {
		panic("use hasNext befor using getNext")
	}
	return s.Next
}
