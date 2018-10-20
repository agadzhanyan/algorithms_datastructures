package array_sorts

import (
	"math"
)

func (s *MyArray) MergeSort() {
	*s = s.merge(*s)
}

func (s* MyArray) merge(slice MyArray) MyArray {
	if len(slice) > 1 {
		resultSlice := MyArray{}

		n := int(math.Round(float64(len(slice) / 2)))
		firstSlice := s.merge(slice[:n])
		secondSlice := s.merge(slice[n:])

		firstIndex := 0
		secondIndex := 0

		firstIndexOk := firstIndex < len(firstSlice)
		secondIndexOk := secondIndex < len(secondSlice)

		for firstIndexOk || secondIndexOk {
			if firstIndexOk && secondIndexOk {
				if firstSlice[firstIndex] < secondSlice[secondIndex] {
					resultSlice = append(resultSlice, firstSlice[firstIndex])
					firstIndex++
				} else {
					resultSlice = append(resultSlice, secondSlice[secondIndex])
					secondIndex++
				}
			} else if firstIndexOk {
				resultSlice = append(resultSlice, firstSlice[firstIndex])
				firstIndex++
			} else if secondIndexOk {
				resultSlice = append(resultSlice, secondSlice[secondIndex])
				secondIndex++
			}

			firstIndexOk = firstIndex < len(firstSlice)
			secondIndexOk = secondIndex < len(secondSlice)
		}

		slice = resultSlice
	}

	return slice
}