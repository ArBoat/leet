/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	evenHead := head.Next
	odd, even := head, evenHead
	for even != nil && even.Next != nil {
		/*
				odd.Next = even.Next
		        odd = odd.Next
		        even.Next = odd.Next
		        even = even.Next
		*/
		odd.Next, even.Next = odd.Next.Next, even.Next.Next // 多赋值
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}