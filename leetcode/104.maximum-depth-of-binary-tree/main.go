package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	if l > r {
		return l
	}

	return r
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for ; size > 0; size-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

func main() {
	maxDepth(nil)
}
