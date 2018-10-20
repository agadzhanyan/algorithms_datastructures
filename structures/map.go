package structures

type MyMap struct {
	elements map[int]string
}

func (s *MyMap) Set(key int, value string) {
	s.elements[key] = value
}

func (s *MyMap) Delete(key int) {
	if !s.Has(key) {
		return
	}

	delete(s.elements, key)
}

func (s *MyMap) Has(key int) bool {
	_, ok := s.elements[key]
	return ok
}

func (s *MyMap) Get(key int) string {
	if !s.Has(key) {
		return ""
	}

	return s.elements[key]
}

func (s *MyMap) Keys() []int {

	keys := make([]int, 0, len(s.elements))
	for k, _ := range s.elements {
		keys = append(keys, k)
	}

	return keys
}

func (s *MyMap) Clear() {
	s.elements = make(map[int]string, 0)
}
