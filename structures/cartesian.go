package structures

type CartesianTree struct {
	root CartesianTreeNode
}

type CartesianTreeNode struct {
	x int
	y int
	left *CartesianTreeNode
	right *CartesianTreeNode
}

func Merge(left,right *CartesianTreeNode) *CartesianTreeNode  {
	if left == nil {
		return right;
	}
	if right == nil {
		return left;
	}

	if left.y > right.y {
		newR := Merge(left.right, right);
		return &CartesianTreeNode{left.x, left.y, left.left, newR}
	} else {
		newL := Merge(left, right.left);
		return &CartesianTreeNode{right.x, right.y, newL, right.right}
	}
}
