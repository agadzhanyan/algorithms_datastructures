package structures

import (
	"testing"
)

func TestBTree(t *testing.T) {
	testStruct := &BTree{}
	testStruct.addElement(10, "test10")
	testStruct.addElement(15, "test15")
	testStruct.addElement(7, "test7")
	testStruct.addElement(12, "test12")
	testStruct.addElement(17, "test17")
	testStruct.addElement(20, "test20")
	testStruct.addElement(44, "test44")
	testStruct.addElement(56, "test56")
	testStruct.addElement(96, "test96")
}