package array_sorts

func (s *MyArray) QuickSort() {
	slice := *s
	s.quickSort(0, len(slice) - 1)
}

func (s *MyArray) quickSort(left, right int) {
	// для учебных целях опорный элемент - первый элемент слайса
	// это не самый лучший выбор, плохо отработает в массиве, отсортированном в обратном порядке
	// худший случай O(n^2), если при каждом рекурсивном вызове quickSort, опорный элемент
	// будет делить исходный массив на два таких, что в первом 1 элемент, во втором n-1 элементов

	// берем значение указателя
	slice := *s

	pivot := left
	left = pivot + 1

	l := left
	r := right

	for {
		if left >= right {
			break
		}

		if slice[right] < slice[pivot] {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			continue
		}

		if slice[right] < slice[left] {
			slice[left], slice[right] = slice[right], slice[left]
		}

		right--
	}

	if left - 1 != pivot {
		pivotValue := slice[pivot]
		newSlice := []int{}
		for k, v := range slice {
			if k == pivot {
				continue
			}
			newSlice = append(newSlice, v)
			if k == left {
				newSlice = append(newSlice, pivotValue)
			}
		}
		*s = newSlice
	}

	if right > l {
		s.quickSort(l, right)
	}
	if left < r {
		s.quickSort(left, r)
	}
}
