package array_sorts

import (
	"reflect"
	"testing"
)

func TestMyArray_BubbleSort(t *testing.T) {

	array := MyArray{}
	array.Add(1)
	array.Add(10)
	array.Add(3)
	array.Add(25)
	array.Add(2)

	array.BubbleSort()

	expected := MyArray{1,2,3,10,25}
	if !reflect.DeepEqual(array, expected) {
		t.Errorf("\n Expected: %v \n Got: %v \n", expected, array)
	}
}
