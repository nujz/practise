package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}
	if l < n {
		return head
	}
	if l > 0 && l == n {
		next := head.Next
		head.Next = nil
		return next
	}

	f := l - n
	i := 0
	cur = head
	for cur != nil {
		i++
		if i == f {
			next := cur.Next
			cur.Next = next.Next
			next.Next = nil
			return head
		}
		cur = cur.Next
	}

	return head
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	cur := head
	for cur != nil && n > 0 {
		cur = cur.Next
		n--
	}
	fast := cur

	if fast == nil {
		next := head.Next
		head.Next = nil
		return next
	}

	cur = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		cur = cur.Next
	}

	next := cur.Next
	cur.Next = next.Next
	next.Next = nil

	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	prev, slow, fast := dummy, head, head
	for fast != nil {
		if n <= 0 {
			prev = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	prev.Next = slow.Next

	return dummy.Next
}

func main() {
}
