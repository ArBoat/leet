/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		//cur.Next.Next = pre.Next 插到头部
		cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, pre.Next, cur.Next
	}
	return dummy.Next
}

/*
头插法
*/