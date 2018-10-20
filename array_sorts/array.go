package array_sorts

type MyArray []int

func (s *MyArray) Add(value int) {
	*s = append(*s, value)
}