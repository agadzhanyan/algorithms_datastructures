package structures

import (
	"testing"
	"reflect"
	)

func TestBinaryHeap_BuildHeap(t *testing.T) {

	expected := []*Unit{
		&Unit{4,4},
		&Unit{6,6},
		&Unit{10, 10},
		&Unit{17, 17},
		&Unit{20, 20},
		&Unit{53, 53},
	}
	got := []*Unit{}

	elements := make([]*Unit, 0)
	elements = append(
		elements,
		&Unit{17, 17},
		&Unit{20, 20},
		&Unit{6, 6},
		&Unit{4, 4},
		&Unit{53,53},
		&Unit{10,10},
	)

	testStruct := &BinaryHeap{make(map[int]*Unit, 0), 0}
	testStruct.BuildHeap(elements)

	for !testStruct.IsEmpty() {
		got = append(got, testStruct.Extract())
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}

	testStructSecond := &BinaryHeap{make(map[int]*Unit, 0), 0}
	testStructSecond.BuildHeapFast(elements)

	gotSecond := []*Unit{}
	for !testStructSecond.IsEmpty() {
		gotSecond = append(gotSecond, testStructSecond.Extract())
	}

	if !reflect.DeepEqual(expected, gotSecond) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, gotSecond)
	}
}
