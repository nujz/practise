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

func isSymmetric2(root *TreeNode) bool {
	l, r := root, root
	var arr []*TreeNode
	arr = append(arr, l)
	arr = append(arr, r)
	for len(arr) > 0 {
		l, r := arr[0], arr[1]
		arr = arr[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		arr = append(arr, l.Left)
		arr = append(arr, r.Right)
		arr = append(arr, l.Right)
		arr = append(arr, r.Left)
	}
	return true
}

func main() {
	isSymmetric(nil)
}
