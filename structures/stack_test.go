package structures

import (
	"testing"
	"reflect"
)

func TestStructIntStackContainer(t *testing.T) {
	expected := []int{5,6,2,1}
	got := []int{}

	testStruct := &IntStackContainer{make([]int, 0)}
	testStruct.Add(1)
	testStruct.Add(2)
	testStruct.Add(6)
	testStruct.Add(5)
	for !testStruct.IsEmpty() {
		got = append(got, testStruct.Get())
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}
}
