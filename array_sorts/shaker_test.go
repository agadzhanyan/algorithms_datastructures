package array_sorts

import (
	"reflect"
	"testing"
)

func TestMyArray_ShakerSort(t *testing.T) {

	array := MyArray{}
	array.Add(1)
	array.Add(10)
	array.Add(3)
	array.Add(25)
	array.Add(2)
	array.Add(5)
	array.Add(9)
	array.Add(7)
	array.Add(6)

	array.ShakerSort()

	expected := MyArray{1,2,3,5,6,7,9,10,25}
	if !reflect.DeepEqual(array, expected) {
		t.Errorf("\n Expected: %v \n Got: %v \n", expected, array)
	}
}
