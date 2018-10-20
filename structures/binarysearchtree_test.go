package structures

import (
	"testing"
	"reflect"
)

func TestBinarySearchTree(t *testing.T) {

	testStruct := &BinarySearchTree{}
	testStruct.AddElement(15, "15")
	testStruct.AddElement(20, "20")
	testStruct.AddElement(10, "10")
	testStruct.AddElement(11, "11")
	testStruct.AddElement(14, "14")
	testStruct.AddElement(18, "18")
	testStruct.AddElement(16, "16")

	expectedFirst := []string{"10", "11", "14", "15", "16", "18", "20"}
	gotFirst := testStruct.InfixTraversal()

	if !reflect.DeepEqual(expectedFirst, gotFirst) {
		t.Errorf("wrong result #1")
	}

	val, err := testStruct.Find(20)
	if err != nil || val != "20" {
		t.Errorf("wrong result #2")
	}

	val, err = testStruct.Find(202)
	if err == nil {
		t.Errorf("wrong result #3")
	}

	testStruct.Delete(15)
	expectedThree := []string{"16", "10", "11", "14", "20", "18"}
	gotThree := testStruct.PrefixTraversal()
	if !reflect.DeepEqual(expectedThree, gotThree) {
		t.Errorf("wrong  result #3")
	}
}