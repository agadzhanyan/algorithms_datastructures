package structures

import (
	"testing"
	"reflect"
)

func TestIntSet(t *testing.T) {

	expected := []int{15,4,0}
	got := []int{}

	testStruct := &IntSet{}
	testStruct.add(15)
	testStruct.add(4)
	testStruct.add(15)
	testStruct.add(0)
	testStruct.add(15)

	for _, v := range testStruct.getElements() {
		got = append(got, v)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Wrong result \n Expected: %v \n Got: %v", expected, got)
	}
}
