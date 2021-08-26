package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func check(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func main() {
	isSymmetric(nil)
}
