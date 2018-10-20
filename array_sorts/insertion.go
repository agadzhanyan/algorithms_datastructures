package array_sorts

func (s *MyArray) InsertionSort() {
	slice := *s
	currentIndex := 0
	for i := 0; i < len(slice); i++ {
		for j := 0; j <= currentIndex; j++ {
			if slice[j] > slice[i] {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
		currentIndex++
	}
}
