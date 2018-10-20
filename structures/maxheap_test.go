package structures

import (
	"reflect"
	"testing"
)

func TestBinaryMaxHeap_BuildHeap(t *testing.T) {

	expected := []*Unit{
		&Unit{53, 53},
		&Unit{20, 20},
		&Unit{17, 17},
		&Unit{10, 10},
		&Unit{6,6},
		&Unit{4,4},
	}

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

	testStructSecond := &BinaryMaxHeap{make(map[int]*Unit, 0), 0}
	testStructSecond.BuildHeapFast(elements)

	gotSecond := []*Unit{}
	for !testStructSecond.IsEmpty() {
		gotSecond = append(gotSecond, testStructSecond.Extract())
	}

	if !reflect.DeepEqual(expected, gotSecond) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, gotSecond)
	}
}
