package structures

import (
	"testing"
	"reflect"
)

func TestRadixTree(t *testing.T) {

	tree := RadixTree{head:&RadixTreeNode{childNodes: map[rune]*RadixTreeNode{}}}
	tree.addElement(tree.head,"David", 150)
	tree.addElement(tree.head,"Davlan", 250)
	tree.addElement(tree.head,"Davidos", 250)
	tree.addElement(tree.head,"Dav", 650)
	tree.addElement(tree.head,"Davy", 250)
	tree.addElement(tree.head,"David" +
		"", 250)

	got := tree.getKeys()
	if len(got) != 5 {
		t.Errorf("wrong result \n Expecte elements: %v \n Got: %v", 5, got)
	}

	got = tree.getKeysByPrefix("Davi")
	expected := []string{"David", "Davidos"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("wrong result \n Expected: %v \n Got: %v", got, expected)
	}
}
