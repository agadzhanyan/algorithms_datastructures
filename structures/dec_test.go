package structures

import (
	"testing"
	"reflect"
)

func TestIntDecContainer(t *testing.T) {
	expected := []int{2, 3, 6, 5, 3}
	got := []int{}

	// 2,6,3,5,3
	// 6,3
	testStruct := &IntDecContainer{make([]int, 0)}
	testStruct.addFirst(6)
	testStruct.addLast(3)
	testStruct.addFirst(2)
	testStruct.addLast(5)
	testStruct.addLast(3)

	got = append(got, testStruct.getFirst(), testStruct.getLast(), testStruct.getFirst(), testStruct.getLast(), testStruct.getFirst())

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}
}
