package algorithms

import (
	"structures"
)

func (s *graph) SearchBST(x,y int) ([]int, bool) {
	visited := make(map[int]bool, 0)
	paths := make(map[int]*structures.IntSinglelinkedList, 0)

	paths[x] = &structures.IntSinglelinkedList{}
	paths[x].AddLast(x)

	res := make([]int, 0)
	fifo := &structures.IntQueueContainer{}
	fifo.Add(x)
	for !fifo.IsEmpty() {
		curVertexKey := fifo.Get()
		for _,v := range s.vertex[curVertexKey] {
			_, ok := visited[v]
			if ok {
				continue
			}
			fifo.Add(v)
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

func main() {
	graph := graph{make(map[int][]int, 0)}
	graph.vertex[1] = []int{4,7,2}
	graph.vertex[2] = []int{1,5,3}
	graph.vertex[3] = []int{2,6}
	graph.vertex[4] = []int{1,7}
	graph.vertex[5] = []int{2,6,7}
	graph.vertex[6] = []int{3,5,7}
	graph.vertex[7] = []int{4,5,6}

	graph.SearchBST(1,7)
}
