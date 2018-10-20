package structures

import (
	"math"
)

type Unit struct {
	Value int
	Priority int
}

type BinaryMaxHeap struct {
	Elements map[int]*Unit
	HeapSize int
}

func (s *BinaryMaxHeap) NewUnit() *Unit {
	return &Unit{}
}

// построение кучи за сложность O(n)
func (s *BinaryMaxHeap) BuildHeapFast(elements []*Unit) {
	if len(elements) < 1 {
		panic("empty slice")
	}
	if len(elements) == 1 {
		s.Elements[0] = elements[0]
		s.HeapSize = 1
		return
	}
	for k, v := range elements {
		s.Elements[k] = v
	}
	s.HeapSize = len(elements)
	index := int(math.Round(float64(s.HeapSize-1/2)))
	for index >= 0 {
		s.Heapify(index)
		index--
	}
}

func (s  *BinaryMaxHeap) IsEmpty() bool {
	return s.HeapSize == 0
}

func (s *BinaryMaxHeap) Extract() *Unit {
	if s.HeapSize == 0 {
		panic("error")
	}

	element := s.Elements[0]
	s.Elements[0] = s.Elements[s.HeapSize - 1]
	delete(s.Elements, s.HeapSize - 1)

	s.HeapSize--
	s.Heapify(0)

	return element
}

func (s *BinaryMaxHeap) Add(unit *Unit) {
	s.HeapSize++

	// позиция только что вставленного элемента
	place := s.HeapSize-1
	s.Elements[place] = unit

	// находим позицию родительского элемента
	// так как нумерация с 0, то вычитаем 1 и округляем в меньшую сторону
	parentPlace := s.getParentIndex(place)
	for parentPlace >= 0 {
		if s.Elements[parentPlace].Priority < s.Elements[place].Priority {
			s.swap(parentPlace, place)

			place = parentPlace
			parentPlace = s.getParentIndex(place)
		} else {
			break
		}
	}
}

func (s *BinaryMaxHeap) getParentIndex(index int) int {
	return int(math.RoundToEven(float64(index - 1)/2))
}

func (s *BinaryMaxHeap) swap(first,second int) {
	temp := s.Elements[second]
	s.Elements[second] = s.Elements[first]
	s.Elements[first] = temp
}

func (s *BinaryMaxHeap) Heapify(index int) {
	for {
		left := 2*index+1
		right := 2*index+2
		least := index
		if left < s.HeapSize && s.Elements[least].Priority < s.Elements[left].Priority {
			least = left
		}
		if right < s.HeapSize && s.Elements[least].Priority < s.Elements[right].Priority {
			least = right
		}

		if least == index {
			break
		}

		s.swap(least, index)
		index = least
	}
}