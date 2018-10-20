package array_sorts

func (s *MyArray) GnomeSort() {
	slice := *s
	index := 1
	for index < len(slice) {
		if slice[index] > slice[index - 1] {
			index++
		} else {
			slice[index], slice[index - 1] = slice[index - 1], slice[index]
			index--
		}
	}
}
