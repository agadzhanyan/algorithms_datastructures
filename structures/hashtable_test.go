package structures

import (
	"testing"
			"sync"
)

func TestHashTable(t *testing.T) {

	testStruct := &HashTable{make(map[int]string, 0), sync.RWMutex{}}
	testStruct.add("first", "test1")
	testStruct.add("second", "test2")
	testStruct.add("first", "test3")

	if testStruct.Get("first") != "test3" {
		t.Errorf("Wrong result")
	}

	testStruct.Delete("first")

	if testStruct.Get("first") != "" {
		t.Errorf("Wrong result")
	}

	if testStruct.Size() != 1 {
		t.Errorf("Wrong result \n Expected: 1 \n Got: %v", testStruct.Size())
	}
}