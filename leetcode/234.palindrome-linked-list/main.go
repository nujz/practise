package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 快慢指针 + 反转链表*2
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	end := reverseList(slow)

	left := head
	right := end
	for right != nil {
		if right.Val != left.Val {
			reverseList(end)
			return false
		}
		right = right.Next
		left = left.Next
	}

	reverseList(end)

	return true
}

func isPalindrome0(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

func isPalindrome1(head *ListNode) bool {
	var cached []int
	for node := head; node != nil; node = node.Next {
		cached = append(cached, node.Val)
	}

	l := len(cached)
	for i := 0; i < (l+1)/2; i++ {
		if cached[i] != cached[l-1-i] {
			return false
		}
	}

	return true
}

func main() {

}
