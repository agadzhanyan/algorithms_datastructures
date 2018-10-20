package structures

type IntStackContainer struct {
	intSlice []int
}

func (s *IntStackContainer) IsEmpty() bool {
	if len(s.intSlice) == 0 {
		return true
	}

	return false
}

func (s *IntStackContainer) Get() int {
	if s.IsEmpty() {
		panic("use isEmpty before using Get")
	}

	elem := s.intSlice[len(s.intSlice)-1]
	s.intSlice = s.intSlice[0:len(s.intSlice)-1]
	return elem
}

func (s *IntStackContainer) Add(element int) {
	s.intSlice = append(s.intSlice, element)
}

