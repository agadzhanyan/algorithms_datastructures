package algorithms

import (
	"math/rand"
	"structures"
)

// эвристическая функция, которая определяет условный приоритет между текущем вершиной и искомой
// возвращает int значение приоритета
// эвристику я не придумал, поэтому просто возвращаю рандомное int значение
// конечно же, от такой эвристической функции толку мало, но для демонстрации алгоритма BFS сойдет
func (s *graphw) a_star_heuristic(from, to *structures.Unit) int {
	return rand.Int()
}

func (s *graphw) SearchAStar(x,y int) (int, []int, bool) {
	res := make([]int, 0)

	visited := make(map[int]bool, 0)
	paths := make(map[int]*structures.IntSinglelinkedList, 0)

	paths[x] = &structures.IntSinglelinkedList{}
	paths[x].AddLast(x)

	vertexWeigths := map[int]int{}
	for k,_ := range s.vertex {
		vertexWeigths[k] = -1
	}
	vertexWeigths[x] = 0

	mainUnit := &structures.Unit{x, 0}
	minHeap := structures.BinaryHeap{make(map[int]*structures.Unit, 0), 0}
	minHeap.Add(mainUnit)

	for !minHeap.IsEmpty() {
		// проверяем, проходили ли мы по текущей вершине
		currentUnit := minHeap.Extract()
		curVertexKey := currentUnit.Value
		_, ok := visited[curVertexKey]
		if ok {
			continue
		}
		visited[curVertexKey] = true

		for k, v := range s.vertex[curVertexKey] {
			// вычисляем новый вес
			newWeight := vertexWeigths[curVertexKey] + v

			_, ok := paths[y]
			if ok {
				// путь до искомого значения найден
				// завершение выполнения
				break
			}

			if vertexWeigths[k] == -1 || newWeight < vertexWeigths[k] {
				// добавляем вершину в кучу
				// приоритет - результат вычисления по эвристической функции
				newUnit := &structures.Unit{k, newWeight + s.a_star_heuristic(currentUnit, mainUnit)}
				minHeap.Add(newUnit)

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
