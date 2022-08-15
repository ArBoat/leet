/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	// 找到长度时p要为最后一个元素， 方便成环
	p := head
	size := 1
	for p.Next != nil {
		size++
		p = p.Next
	}
	offset := size - k%size
	if offset == size {
		return head
	}
	p.Next = head
	for offset > 0 {
		p = p.Next
		offset--
	}
	res := p.Next
	p.Next = nil
	return res
}

/*
1.求出链表长度
2.成环
3.移动 n-k&n
4.断开
*/