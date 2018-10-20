package algorithms

import (
	"structures"
)

func (s *graph) SearchDST(x,y int) ([]int, bool) {
	visited := make(map[int]bool, 0)
	paths := make(map[int]*structures.IntSinglelinkedList, 0)

	paths[x] = &structures.IntSinglelinkedList{}
	paths[x].AddLast(x)

	res := make([]int, 0)
	stack := &structures.IntStackContainer{}
	stack.Add(x)
	for !stack.IsEmpty() {
		curVertexKey := stack.Get()
		for _,v := range s.vertex[curVertexKey] {
			_, ok := visited[v]
			if ok {
				continue
			}
			stack.Add(v)
			visited[v] = true

			node := paths[curVertexKey].GetHead()
			nextList := &structures.IntSinglelinkedList{}
			nextList.AddLast(node.Value)
			for node.HasNext() {
				node = node.GetNext()
				nextList.AddLast(node.Value)
			}
			nextList.AddLast(v)
			paths[v] = nextList
		}
	}

	list, ok := paths[y]
	if !ok {
		return nil, false
	}
	node := list.GetHead()
	res = append(res, node.Value)
	for node.HasNext() {
		node = node.GetNext()
		res = append(res, node.Value)
	}

	return res, true
}

