package array_sorts

func (s *MyArray) SelectionSort() {
	slice := *s
	for i := 0; i < len(slice); i++ {
		minIndex := i
		for j := minIndex; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		slice[i], slice[minIndex] = slice[minIndex], slice[j]
	}
}
