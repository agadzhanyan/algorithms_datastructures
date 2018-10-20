package structures

import "fmt"

type BinarySearchTree struct {
	head *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	left, right *BinarySearchTreeNode
	key int
	value string
}

// префиксный обход бинарного дерева
func (s *BinarySearchTree) PrefixTraversal() []string {

	res := make([]string, 0)
	s.__prefixTraversal(s.head, &res)
	return res
}

func (s *BinarySearchTree) __prefixTraversal(node *BinarySearchTreeNode, res *[]string) {

	*res = append(*res, node.value)
	if node.left != nil {
		s.__prefixTraversal(node.left, res)
	}
	if node.right != nil {
		s.__prefixTraversal(node.right, res)
	}
}

// инфиксный обход дерева
func (s *BinarySearchTree) InfixTraversal() []string {

	res := make([]string, 0)
	s.__infixTraversal(s.head, &res)
	return res
}

func (s *BinarySearchTree) __infixTraversal(node *BinarySearchTreeNode, res *[]string) {

	if node.left != nil {
		s.__infixTraversal(node.left, res)
	}
	*res = append(*res, node.value)
	if node.right != nil {
		s.__infixTraversal(node.right, res)
	}
}

func (s *BinarySearchTree) Find(key int) (string, error) {

	if s.head == nil {
		return "", fmt.Errorf("empty tree")
	}

	node := s.head
	for {
		if node.key == key {
			return node.value, nil
		}
		if key > node.key {
			if node.right == nil {
				return "", fmt.Errorf("not found")
			} else {
				node = node.right
			}
		} else {
			if node.left == nil {
				return "", fmt.Errorf("not found")
			} else {
				node = node.left
			}
		}
	}
}

func (s *BinarySearchTree) AddElement(key int, value string) {

	var node *BinarySearchTreeNode
	if s.head == nil {
		s.head = &BinarySearchTreeNode{}
		node = s.head
	} else {
		node = s.head
		for {
			if key >= node.key {
				if node.right == nil {
					newNode := &BinarySearchTreeNode{}
					node.right = newNode
					node = newNode
					break
				}
				node = node.right
			} else {
				if node.left == nil {
					newNode := &BinarySearchTreeNode{}
					node.left = newNode
					node = newNode
					break
				}
				node = node.left
			}
		}
	}

	node.key = key
	node.value = value
}

func (s *BinarySearchTree) Delete(key int) {

	if s.head == nil {
		return
	}

	var prev *BinarySearchTreeNode
	node := s.head
	for {
		if node.key == key {
			var newNode *BinarySearchTreeNode
			if node.right != nil {
				if  node.right.left != nil {
					newNode = node.right.left
					var prevNewNode *BinarySearchTreeNode
					for {
						if newNode.left != nil {
							prevNewNode = newNode
							newNode = newNode.left
						} else {
							if prevNewNode != nil {
								prevNewNode.left = nil
							}
							break
						}
					}
				} else {
					newNode = node.right
					node.right = nil
				}
			} else if node.left != nil {
				newNode = node.left
				node.left = nil
			}
			if newNode == nil {
				node = nil
			} else {
				node.key = newNode.key
				node.value = newNode.value
			}
			if prev != nil {}
			return
		}
		if key > node.key {
			if node.right == nil {
				return
			} else {
				prev = node
				node = node.right
			}
		} else {
			if node.left == nil {
				return
			} else {
				prev = node
				node = node.left
			}
		}
	}
}
