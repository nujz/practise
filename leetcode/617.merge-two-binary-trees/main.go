package main

import (
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 注意 nil 与 *TreeNode (typed nil) 的区别
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var node TreeNode

	if root1 != nil && root2 != nil {
		node.Val = root1.Val + root2.Val
		node.Left = mergeTrees(root1.Left, root2.Left)
		node.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 != nil {
		node.Val = root1.Val
		node.Left = mergeTrees(root1.Left, nil)
		node.Right = mergeTrees(root1.Right, nil)
	} else if root2 != nil {
		node.Val = root2.Val
		node.Left = mergeTrees(nil, root2.Left)
		node.Right = mergeTrees(nil, root2.Right)
	} else {
		return nil
	}

	return &node
}

func main() {
	root1 := TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	root2 := TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}},
	}
	root := mergeTrees(&root1, &root2)
	fmt.Println(root)
}

// TODO
func echo(root *TreeNode) []string {
	return nil
}
