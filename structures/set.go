package structures

type IntSet struct {
	elements map[int]bool
}

func (s *IntSet) add(element int) {

	if s.elements == nil {
		s.elements = make(map[int]bool)
	}
	if !s.hasElement(element) {
		s.elements[element] = true
	}
}

func (s *IntSet) hasElement(element int) bool {

	_,ok := s.elements[element]
	return ok
}

func (s *IntSet) getElements() []int {

	slice := make([]int, 0, len(s.elements))
	for k,_ := range s.elements {
		slice = append(slice, k)
	}

	return slice
}
