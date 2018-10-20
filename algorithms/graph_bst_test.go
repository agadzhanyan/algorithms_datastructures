package algorithms

import (
	"reflect"
	"testing"
)

func TestGraph_SearchBST(t *testing.T) {

	graph := graph{make(map[int][]int, 0)}
	graph.vertex[1] = []int{4,7,2}
	graph.vertex[2] = []int{1,5,3}
	graph.vertex[3] = []int{2,6}
	graph.vertex[4] = []int{1,7}
	graph.vertex[5] = []int{2,6,7}
	graph.vertex[6] = []int{3,5,7}
	graph.vertex[7] = []int{4,5,6}

	got, ok := graph.SearchBST(1,7)
	if !ok {
		t.Errorf("error")
	}

	expected := []int{1,7}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected: %v \n Got: %v \n", expected, got)
	}
}