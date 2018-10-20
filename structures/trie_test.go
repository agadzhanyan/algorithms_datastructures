package structures

import (
	"testing"
)

func TestTrie(t *testing.T) {

	trie := &Trie{headNone:&TrieNode{make(map[rune]*TrieNode, 0), false, ""}}
	trie.addElement("test", "check1")
	trie.addElement("testing", "check2")
	trie.addElement("tef", "op?")

	val, err := trie.getValue("test")
	if err != nil {
		t.Errorf("wrong result")
	}
	if val != "check1" {
		t.Errorf("wrong result")
	}

	val, err = trie.getValue("testing")
	if val != "check2" {
		t.Errorf("wrong result")
	}

	val, err = trie.getValue("")
	if err == nil {
		t.Errorf("wrong result")
	}

	availableKeys := trie.getAvailableKeys("te")
	if len(availableKeys) != 3 {
		t.Errorf("wrong result")
	}

	trie.removeElement("tef")
	availableKeys = trie.getAvailableKeys("te")
	if len(availableKeys) != 2 {
		t.Errorf("wrong result")
	}
}
