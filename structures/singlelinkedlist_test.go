package structures

import (
	"testing"
	"reflect"
	)

func TestIntSinglelinkedList(t *testing.T) {
	expected := []int{1,5,2,4,3}
	got := []int{}

	testStruct := &IntSinglelinkedList{}
	testStruct.AddFirst(5)
	testStruct.AddLast(2)
	testStruct.AddLast(4)
	testStruct.AddFirst(1)
	testStruct.AddLast(3)

	head := testStruct.GetHead()
	got = append(got, head.Value)
	for head.HasNext() {
		head = head.GetNext()
		got = append(got, head.Value)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}

	testStruct.delAll()

	expected = []int{3,4,5,6}
	got = []int{}

	testStruct.AddFirst(3)
	testStruct.AddLast(4)
	testStruct.AddLast(6)
	testStruct.AddAfter(5, 2)

	head = testStruct.GetHead()
	got = append(got, head.Value)
	for head.HasNext() {
		head = head.GetNext()
		got = append(got, head.Value)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}
}
