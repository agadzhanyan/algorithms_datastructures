package array_sorts

import (
	"math"
)

func (s *MyArray) ShellSort() {
	slice := *s
	d := (len(slice) + 1) / 2
	for i := d; i > 0; i = int(math.Round(float64(i / 2))) {
		for j := 0; j < d; j++ {
			s.shellInsertionSort(j, i)
		}
	}
}

func (s *MyArray) shellInsertionSort(startIndex, step int) {
	slice := *s
	currentIndex := startIndex
	for i := startIndex; i < len(slice); i = i + step {
		for j := startIndex; j <= currentIndex; j = j + step {
			if slice[j] > slice[i] {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
		currentIndex += step
	}
}