package algorithms

import (
	"structures"
)

func (s *graphw) SearchDijkstra(x,y int) (int, []int, bool) {
	res := make([]int, 0)

	visited := make(map[int]bool, 0)
	visited[x] = true
	paths := make(map[int]*structures.IntSinglelinkedList, 0)

	paths[x] = &structures.IntSinglelinkedList{}
	paths[x].AddLast(x)

	vertexWeigths := map[int]int{}
	for k,_ := range s.vertex {
		vertexWeigths[k] = -1
	}
	vertexWeigths[x] = 0

	minHeap := structures.BinaryHeap{make(map[int]*structures.Unit, 0), 0}
	minHeap.Add(&structures.Unit{x, 0})
	for !minHeap.IsEmpty() {
		// проверяем, проходили ли мы по текущей вершине
		curVertexKey := minHeap.Extract().Value
		visited[curVertexKey] = true

		for k, v := range s.vertex[curVertexKey] {

			// проверяем, что вершину уже посещали
			_, ok := visited[k]
			if ok {
				continue
			}

			// вычисляем новый вес
			newWeight := vertexWeigths[curVertexKey] + v

			// добавляем вершину в кучу
			// приоритет - суммарный вес от начальной точки до текущей
			minHeap.Add(&structures.Unit{k, newWeight})

			if vertexWeigths[k] == -1 || newWeight < vertexWeigths[k] {
				vertexWeigths[k] = newWeight
				// путь к вершине отсутствует
				// либо нашли путь короче
				node := paths[curVertexKey].GetHead()
				nextList := &structures.IntSinglelinkedList{}
				nextList.AddLast(node.Value)
				for node.HasNext() {
					node = node.GetNext()
					nextList.AddLast(node.Value)
				}
				nextList.AddLast(k)
				paths[k] = nextList
			}
		}
	}

	list, ok := paths[y]
	if !ok {
		return -1 ,nil, false
	}
	node := list.GetHead()
	res = append(res, node.Value)
	for node.HasNext() {
		node = node.GetNext()
		res = append(res, node.Value)
	}

	return vertexWeigths[y], res, true
}
