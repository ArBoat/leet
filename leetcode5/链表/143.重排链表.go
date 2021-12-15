/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	mid := midNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil // IM 断开链表
	r2 := reverseList(l2)
	mergerLists(l1, r2)
}

func midNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		pre, head, head.Next = head, head.Next, pre
	}
	return pre
}

func mergerLists(l1, l2 *ListNode) {
	p1, p2 := l1, l2
	for p2 != nil {
		temp1, temp2 := p1, p2 //merge
		p1 = p1.Next
		p2 = p2.Next
		temp1.Next = temp2
		temp2.Next = p1
	}
}