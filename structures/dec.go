package structures

type IntDecContainer struct {
	intSlice []int
}

func (s *IntDecContainer) isEmpty() bool {
	if len(s.intSlice) == 0 {
		return true
	}

	return false
}

func (s *IntDecContainer) addLast(element int) {
	s.intSlice = append(s.intSlice, element)
}

func (s *IntDecContainer) addFirst(element int) {
	newSlice := make([]int, 0, len(s.intSlice) + 1)
	newSlice = append(newSlice, element)
	for _, v := range s.intSlice {
		newSlice = append(newSlice, v)
	}
	s.intSlice = newSlice
}

func (s *IntDecContainer) getFirst() int {
	if s.isEmpty() {
		panic("use isEmpty before using getFirst")
	}

	elem := s.intSlice[0]
	s.intSlice = s.intSlice[1:len(s.intSlice)]
	return elem
}

func (s *IntDecContainer) getLast() int {
	if s.isEmpty() {
		panic("use isEmpty before using getLast")
	}

	elem := s.intSlice[len(s.intSlice)-1]
	s.intSlice = s.intSlice[0:len(s.intSlice)-1]
	return elem
}
