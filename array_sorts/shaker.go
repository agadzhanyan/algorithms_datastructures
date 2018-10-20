package array_sorts

func (s *MyArray) ShakerSort() {
	slice := *s
	left := 0
	right := len(slice) - 1
	for {
		for i := left; i < right; i++ {
			if slice[i] > slice[i + 1] {
				slice[i], slice[i + 1] = slice[i + 1], slice[i]
			}
		}
		right--

		for i := right; i > left; i-- {
			if slice[i] < slice[i - 1] {
				slice[i], slice[i - 1] = slice[i - 1], slice[i]
			}
		}
		left++

		if left == right {
			break
		}
	}
}
