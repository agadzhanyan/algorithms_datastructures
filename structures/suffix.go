package structures

import (
	"math"
	"strings"
)

type SuffixArray struct {
	String string
	Suffixes map[int]string
}

func (s *SuffixArray) GetSuffixArray() map[int]string {
	return s.Suffixes
}

func (s *SuffixArray) GetAllSuffixes() []string {

	res := []string{}
	for _, v := range s.Suffixes {
		res = append(res, v)
	}

	return res
}

// для более быстрого поиска необходимо собирать LCP массив, сложность также O(n)
func (s *SuffixArray) Search(str string) []string {

	suffixes := make([]string, 0)

	l := 0
	r := len(s.String) - 1
	for {
		m := math.Round(float64(l + r) / 2)
		if m < 0 {
			return suffixes
		}
		if strings.HasPrefix(s.Suffixes[int(m)], str) {
			for i := l; i <= r; i ++ {
				if strings.HasPrefix(s.Suffixes[int(i)], str) {
					suffixes = append(suffixes, s.Suffixes[int(i)])
				}
			}
			break
		}
		if s.compareStrings(s.Suffixes[int(m)], str) {
			l = int(m)
		} else {
			r = int(m)
		}
	}

	return suffixes
}

// построение суффиксного массива за O(n*n*n)
// если сортируем массив после сбора суффиксов
//
// сортировка после, то O(n*n)
// если заюзать поразрядную сортировку O(n) - ABC сортировка, то сложность сокращается до O(n)
func (s *SuffixArray) SetString(str string) {

	s.Suffixes = make(map[int]string, 0)
	s.String = ""

	if len(str) < 1 {
		panic("error")
	}

	s.String = str
	for i := 0; i < len(s.String); i++ {
		s.Suffixes[i] = s.String[i:]
		s.sort()
	}
}

func (s *SuffixArray) sort() {

	for i := len(s.Suffixes) - 1; i > 0 ; i-- {
		if !s.compareStrings(s.Suffixes[i - 1], s.Suffixes[i]) {
			s.swap(i, i - 1)
		}
	}
}

// if first > seconds - true
// otherwise - false
func (s *SuffixArray) compareStrings(first, second string) bool {

	lengthFirst := len(first)
	lengthSecond := len(second)
	maxStringLength := lengthFirst
	if first == second {
		return true
	}
	if lengthSecond > lengthFirst {
		maxStringLength = lengthSecond
	}

	for k := 0; k < maxStringLength - 1; k++ {
		if k >= lengthFirst  {
			return true
		}
		if k >= lengthSecond {
			return false
		}
		if first[k] < second[k] {
			return true
		}
		if first[k] > second[k] {
			return false
		}
	}

	return false
}

func (s *SuffixArray) swap(i, j int) {
	s.Suffixes[i], s.Suffixes[j] = s.Suffixes[j], s.Suffixes[i]
}