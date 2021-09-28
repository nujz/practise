package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func symmetric(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

func main() {
}
