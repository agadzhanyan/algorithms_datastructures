package structures

import "fmt"

type Trie struct {
	headNone *TrieNode
}

type TrieNode struct {
	ChildNodes map[rune]*TrieNode
	IsCompleteString bool
	NodeValue string
}

func (s *Trie) addElement(key, value string) {

	node := s.headNone
	i := 0
	length := len(key)
	for _, symbol := range key {
		i++
		currentNode, ok := node.ChildNodes[symbol]
		if !ok {
			currentNode = &TrieNode{ChildNodes: map[rune]*TrieNode{}, NodeValue:"", IsCompleteString:false}
			node.ChildNodes[symbol] = currentNode
		}
		if i == length {
			currentNode.IsCompleteString = true
			currentNode.NodeValue = value
			return
		}
		node = currentNode
	}
}

func (s *Trie) removeElement(key string) {

	node := s.headNone
	prevNode := node
	ok := true
	i := 0
	length := len(key)
	for _, symbol := range key {
		i++
		prevNode = node
		node, ok = node.ChildNodes[symbol]
		if !ok {
		}
		if i == length {
			if len(node.ChildNodes) == 0 {
				delete(prevNode.ChildNodes, rune(key[i-1]))
				return
			}
			node.IsCompleteString = false
		}
	}
}

func (s *Trie) getValue(key string) (string, error) {
	node := s.headNone
	ok := true
	i := 0
	length := len(key)
	for _, symbol := range key {
		i++
		node, ok = node.ChildNodes[symbol]
		if !ok {
			return "", fmt.Errorf("not exist")
		}
		if i == length {
			if node.IsCompleteString {
				return node.NodeValue, nil
			}
			return "", fmt.Errorf("not exist")
		}
	}

	return "", fmt.Errorf("not exist")
}

func (s *Trie) getAvailableKeys(key string) []string {

	availablePrefixes := []string{}
	var found *TrieNode
	node := s.headNone
	ok := true
	i := 0
	length := len(key)
	for _, symbol := range key {
		i++
		node, ok = node.ChildNodes[symbol]
		if !ok {
			return availablePrefixes
		}
		if i == length {
			found = node
		}
	}
	if found == nil {
		return availablePrefixes
	}
	if found.IsCompleteString {
		availablePrefixes = append(availablePrefixes, key)
	}
	s.getKeysRecursive(key, found, &availablePrefixes)
	return availablePrefixes
}

func (s *Trie) getKeysRecursive(preKey string, node *TrieNode, prefixes *[]string) {
	for symbol, curNode := range node.ChildNodes {
		if curNode.IsCompleteString {
			*prefixes = append(*prefixes, preKey + string(symbol))
		}
		s.getKeysRecursive(preKey + string(symbol), curNode, prefixes)
	}
}


