package structures

import (
	"reflect"
)

const t = 2

// реализовано только добавление элемента > существующих в дереве
type BTree struct {
	root *BTreeNode
}

type BTreeNode struct {
	parent *BTreeNode
	keys map[int]*Entry
	childNodes map[int]*BTreeNode
}

type Entry struct {
	key int
	value string
}

func (s *BTree) getNewBTreeNode() *BTreeNode {
	return &BTreeNode{childNodes:make(map[int]*BTreeNode, 0), keys: map[int]*Entry{}}
}

func (s *BTree) getEntries(node *BTreeNode) []Entry {
	entries := []Entry{}
	pos := 0
	for {
		entry, ok := node.keys[pos]
		if !ok {
			break
		}
		entries = append(entries, *entry)
		pos++
	}
	return entries
}

func (s *BTree) addEntry(node *BTreeNode, pos int) {
	newMap := map[int]*Entry{}
	curPos := 0
	setPos := 0
	for {
		entry, ok := node.keys[curPos]
		if !ok {
			if pos + 1 > len(node.keys) {
				newMap[pos] = &Entry{}
			}
			break
		}
		if pos == curPos {
			newMap[setPos] = &Entry{}
			setPos++
		}
		newMap[setPos] = entry
		curPos++
		setPos++
	}
	node.keys = newMap
}

func (s *BTree) addElement(key int, value string) {
	if s.root == nil {
		node := s.getNewBTreeNode()
		node.keys[0] = &Entry{key: key, value: value}
		s.root = node
		return
	}

	node := s.findNode(key)
	if len(node.keys) == 2*t-1 {
		s.splitNode(node)
		node = s.findNode(key)
	}
	pos := 0
	for k, v := range node.keys {
		if key == v.key {
			node.keys[k].key = key
			node.keys[k].value = value
			return
		}
		if key < v.key {
			break
		}
		pos++
	}

	s.addEntry(node, pos)
	node.keys[pos].key = key
	node.keys[pos].value = value
}

func (s *BTree) splitNode(node *BTreeNode) {
	entry, _ := node.keys[t-1]

	leftChild := s.getNewBTreeNode()
	rightChild := s.getNewBTreeNode()

	// сплитим и добавляем ключ в родительскую ноду
	if node.parent != nil {
		parentNode := node.parent
		if reflect.DeepEqual(s.root, node.parent) {
			s.root = parentNode
		}
		pos := 0
		for _, v := range parentNode.keys {
			if entry.key < v.key {
				break
			}
			pos++
		}

		s.addEntry(parentNode, pos)
		parentNode.keys[pos].key = entry.key
		parentNode.keys[pos].value = entry.value

		leftChild.parent = parentNode
		leftChild.keys[0] = node.keys[t-2]

		rightChild.parent = parentNode
		rightChild.keys[0] = node.keys[t]

		parentNode.childNodes[pos] = leftChild
		parentNode.childNodes[pos+1] = rightChild

	} else {
		newNode := s.getNewBTreeNode()
		newNode.keys[0] = entry
		newNode.childNodes[0] = leftChild
		newNode.childNodes[1] = rightChild

		leftChild.parent = newNode
		leftChild.keys[0] = node.keys[t-2]

		rightChild.parent = newNode
		rightChild.keys[0] = node.keys[t]

		*node = *newNode
	}
}

func (s *BTree) findNode(key int) *BTreeNode {
	node := s.root
	pos := 0
	for {
		entry, ok := node.keys[pos]

		// вышли за границу 2t-1
		// проверяем его правый дочерний элемент
		if pos == 2*t {
			childNode, childOk := node.childNodes[pos]
			if !childOk {
				return node
			} else {
				node = childNode
				pos = 0
				continue
			}
		}
		if !ok {
			return node
		}

		// следующий элемент отсутствует в мапе
		// вставляем на эту позицию
		if !ok {
			return node
		}

		// ключ уже существует в структуре, переопределим значение на новое
		if key == entry.key {
			return node
		}

		// новый ключ меньше текущего
		if key < entry.key {
			childNode, ok := node.childNodes[pos]
			if !ok {
				return node
			} else {
				node = childNode
				pos = 0
				continue
			}
		}

		// новый ключ больше текущего
		if key > entry.key {
			_, ok := node.keys[pos+1]
			childNode, okChildNode := node.childNodes[pos+1]
			if !ok && okChildNode {
				node = childNode
				pos = 0
				continue
			}
			pos++
		}
	}
}
