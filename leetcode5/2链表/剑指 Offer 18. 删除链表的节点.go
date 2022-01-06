/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	pre, cur := dummy, head
	for cur.Val != val && cur != nil && cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = cur.Next
	return dummy.Next
}

/*
dummy
双指针
O(n)/O(1)
*/