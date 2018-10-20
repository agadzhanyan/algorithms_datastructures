package structures

import (
	"testing"
	"reflect"
)

func TestIntQueueContainer(t *testing.T) {

	expected := []int{1,6,5,4}
	got := make([]int, 0, 0)

	testStruct := IntQueueContainer{}
	testStruct.Add(1)
	testStruct.Add(6)
	testStruct.Add(5)
	testStruct.Add(4)

	for !testStruct.IsEmpty() {
		got = append(got, testStruct.Get())
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\n Expected: %v \v Got: %v \n", expected, got)
	}
}
