package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]struct{})

	node := head
	for node.Next != nil {
		if _, ok := m[node]; ok {
			return true
		} else {
			m[node] = struct{}{}
		}
		node = node.Next
	}

	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
	}

	return false
}

func main() {
	// 0 -> 1 -> 2 -> 3
	//      ^----------

	head3 := &ListNode{
		Val: 3,
	}
	head2 := &ListNode{
		Val:  2,
		Next: head3,
	}
	head1 := &ListNode{
		Val:  1,
		Next: head2,
	}
	head := &ListNode{
		Val:  0,
		Next: head1,
	}
	head3.Next = head1

	b := hasCycle(head)
	println(b)
}
