package structures

type IntQueueContainer struct {
	qSlice []int
}

func (s *IntQueueContainer) Add(element int) {
	s.qSlice = append(s.qSlice, element)
}

func (s *IntQueueContainer) IsEmpty() bool {
	if len(s.qSlice) == 0 {
		return true
	}

	return false
}

func (s *IntQueueContainer) Get() int {

	if s.IsEmpty() {
		panic("you must use isEmpty to check before use Get")
	}
	elem := s.qSlice[0]
	s.qSlice = s.qSlice[1:len(s.qSlice)]
	return elem
}
