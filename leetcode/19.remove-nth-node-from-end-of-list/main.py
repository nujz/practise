
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        dummy = ListNode(0, head)
        slow, fast = dummy, head
        for i in range(n):
            fast = fast.next

        while fast:
            fast = fast.next
            slow = slow.next

        slow.next = slow.next.next
        return dummy.next

    def removeNthFromEnd0(self, head: ListNode, n: int) -> ListNode:
        node = head
        count = 0
        while node:
            count += 1
            node = node.next

        if count == n:
            if head is None:
                return None
            return head.next

        node = head
        i = count - n
        while i > 0 and node:
            i -= 1
            if i == 0:
                node.next = node.next.next
                break
            node = node.next

        return head
