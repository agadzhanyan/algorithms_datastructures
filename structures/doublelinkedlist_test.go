package structures

import (
	"testing"
	"reflect"
	)

func TestIntDoublelinkedList(t *testing.T) {
	expected := []int{1,5,2,4,3}
	got := []int{}

	testStruct := &IntDoublelinkedList{}
	testStruct.addFirst(5)
	testStruct.addLast(2)
	testStruct.addLast(4)
	testStruct.addFirst(1)
	testStruct.addLast(3)

	head := testStruct.getHead()
	got = append(got, head.Value)
	for head.hasNext() {
		head = head.getNext()
		got = append(got, head.Value)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}

	expected = []int{3,4,2,5,1}
	got = []int{}
	last := testStruct.getLast()
	got = append(got, last.Value)
	for last.hasPrev() {
		last = last.getPrev()
		got = append(got, last.Value)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}

	expected = []int{1,5,2,4,3,15}
	testStruct.addAfter(15, 5)

	got = []int{}
	head = testStruct.getHead()
	got = append(got, head.Value)
	for head.hasNext() {
		head = head.getNext()
		got = append(got, head.Value)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}
}
