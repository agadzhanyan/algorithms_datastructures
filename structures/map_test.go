package structures

import (
	"testing"
	"reflect"
)

func TestMyMap_Clear(t *testing.T) {
	testStruct := &MyMap{make(map[int]string)}
	testStruct.Set(1, "test")
	testStruct.Set(9999, "00131")
	testStruct.Clear()

	keys := testStruct.Keys()
	if len(keys) > 0 {
		t.Errorf("Wrong result")
	}
}

func TestMyMap_Delete(t *testing.T) {
	testStruct := &MyMap{make(map[int]string)}
	testStruct.Set(1, "test")
	testStruct.Set(9999, "00131")

	testStruct.Delete(9999)

	val := testStruct.Get(9999)
	if len(val) > 0 {
		t.Errorf("wrong result")
	}
}

func TestMyMap_Keys(t *testing.T) {
	expected := []int{1, 9999}

	testStruct := &MyMap{make(map[int]string)}
	testStruct.Set(1, "test")
	testStruct.Set(9999, "00131")

	keys := testStruct.Keys()
	if !reflect.DeepEqual(expected, keys) {
		t.Errorf("wrong result")
	}
}

func TestMyMap_Get(t *testing.T) {
	testStruct := &MyMap{make(map[int]string)}
	testStruct.Set(1, "test")
	testStruct.Set(9999, "00131")

	val := testStruct.Get(1)
	if val != "test" {
		t.Errorf("wrong result")
	}
}

func TestMyMap_Has(t *testing.T) {
	testStruct := &MyMap{make(map[int]string)}
	testStruct.Set(1, "test")
	testStruct.Set(9999, "00131")

	if testStruct.Has(15) {
		t.Errorf("wrong result")
	}

	if !testStruct.Has(9999) {
		t.Errorf("wrong result")
	}
}
