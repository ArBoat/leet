/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dump := &ListNode{0, nil}
	dump.Next = head
	p := dump
	for p != nil && p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next //IM: 是else才指向下一个
		}
	}
	return dump.Next
}

/*
测试用例：
[6,2,6,6，3,4,5,6]
头尾连续
*/