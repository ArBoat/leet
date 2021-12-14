/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归 O(n)/O(n)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 迭代 O(n)/O(1)
func swapPairs(head *ListNode) *ListNode {
	dump := &ListNode{0, head}
	p := dump
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := p.Next.Next
		p.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		p = node1
	}
	return dump.Next
}
