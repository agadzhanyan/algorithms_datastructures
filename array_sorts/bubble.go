package array_sorts

func (s *MyArray) BubbleSort() {
	for i := 0; i < len(*s); i++ {
		for j := i + 1; j < len(*s); j++ {
			if (*s)[i] > (*s)[j] {
				(*s)[j], (*s)[i] = (*s)[i], (*s)[j]
			}
		}
	}
}
