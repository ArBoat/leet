/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return nil
	}
	var res *ListNode
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head // 不足结束
		}
		b = b.Next
	}
	res = reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return res
}

// [a,b)
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	for a != b {
		pre, a, a.Next = a, a.Next, pre
	}
	return pre
}