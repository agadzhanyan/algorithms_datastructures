package algorithms

import (
	"reflect"
	"testing"
)

func TestGraphw_SearchBFS(t *testing.T) {

	graphw := graphw{make(map[int]map[int]int, 0)}
	graphw.vertex[1] = map[int]int{2 : 30, 4 : 50}
	graphw.vertex[2] = map[int]int{3 : 20, 5 : 40}
	graphw.vertex[3] = map[int]int{}
	graphw.vertex[4] = map[int]int{7 : 0}
	graphw.vertex[5] = map[int]int{6 : 10, 7 : 15}
	graphw.vertex[6] = map[int]int{3 : 30}
	graphw.vertex[7] = map[int]int{6 : 20}

	got, ok := graphw.SearchBFS(1,6)
	if !ok {
		t.Errorf("error")
	}

	expected := []int{1,4,7,6}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected: %v \n Got: %v \n", expected, got)
	}
}
