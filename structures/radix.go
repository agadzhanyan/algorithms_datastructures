package structures

import (
	"strings"
)

type RadixTree struct {
	head *RadixTreeNode
}

type RadixTreeNode struct {
	// префикс ключа узла
	prefix string
	childNodes map[rune]*RadixTreeNode
	isCompletedNode bool
	value int
}

func (s *RadixTree) addElement(node *RadixTreeNode, key string, value int) {

	firstSymbol := rune(key[0])
	nextNode, ok := node.childNodes[firstSymbol]
	if !ok {
		nextNode = &RadixTreeNode{prefix: key, childNodes: map[rune]*RadixTreeNode{}, value:value}
		node.childNodes[firstSymbol] = nextNode
		nextNode.isCompletedNode = true
		return
	}
	if strings.EqualFold(key, nextNode.prefix) {
		nextNode.value = value
		nextNode.isCompletedNode = true
		return
	}

	newPrefix := ""
	for _, symbol := range key {
		newPrefix = newPrefix + string(symbol)
		if strings.Index(nextNode.prefix, newPrefix) == -1 {
			newPrefix = newPrefix[:len(newPrefix)-1]
			break
		}
	}

	nextPrefix := strings.Replace(nextNode.prefix, newPrefix, "", -1)
	remPrefix := strings.Replace(key, newPrefix, "", -1)
	if len(newPrefix) == len(nextNode.prefix) {
		s.addElement(nextNode, remPrefix, value)
		return
	}
	nextNode.isCompletedNode = false
	nextNode.prefix = newPrefix
	nextNode.value = 0

	node = &RadixTreeNode{childNodes: map[rune]*RadixTreeNode{}, prefix:nextPrefix, value:nextNode.value}
	node.isCompletedNode = true
	nextNode.childNodes[rune(nextPrefix[0])] = node

	s.addElement(nextNode, remPrefix, value)
}

func (s *RadixTree) getKeysByPrefix(prefix string) []string {
	allKeys := []string{}
	node := s.head
	for _, v := range node.childNodes {
		allKeys = append(allKeys, s.getRecursiveNodeKeys("", prefix, v)...)
	}
	return allKeys
}

func (s *RadixTree) getKeys() []string {
	allKeys := []string{}
	node := s.head
	for _, v := range node.childNodes {
		allKeys = append(allKeys, s.getRecursiveNodeKeys("", "", v)...)
	}
	return allKeys
}

func (s *RadixTree) getRecursiveNodeKeys(prefix, need string, node *RadixTreeNode) []string {
	keys := []string{}
	if node.isCompletedNode {
		if len(need) > 0 {
			if strings.Index(prefix + node.prefix, need) == 0 {
				keys = append(keys, prefix + node.prefix)
			}
		} else {
			keys = append(keys, prefix + node.prefix)
		}
	}
	for _, v := range node.childNodes {
		keys = append(keys, s.getRecursiveNodeKeys(prefix + node.prefix, need, v)...)
	}
	return keys
}