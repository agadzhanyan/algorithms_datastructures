package array_sorts

import (
	"structures"
)

func (s *MyArray) HeapSort() {
	slice := *s

	heap := structures.BinaryMaxHeap{make(map[int]*structures.Unit, 0), 0}
	for _, v := range slice {
		heap.Add(&structures.Unit{v, v})
	}

	stack := structures.IntStackContainer{}
	for !heap.IsEmpty() {
		stack.Add(heap.Extract().Value)
	}
	newSlice := []int{}
	for !stack.IsEmpty() {
		newSlice = append(newSlice, stack.Get())
	}
	*s = newSlice
}
