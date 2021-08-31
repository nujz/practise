package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	node := head.Next
	prev.Next = nil
	for node.Next != nil {
		tmp := node
		node = node.Next

		tmp.Next = prev
		prev = tmp
	}
	node.Next = prev
	return node
}

func main() {

}
