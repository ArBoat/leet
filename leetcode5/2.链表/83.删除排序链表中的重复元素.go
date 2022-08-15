/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{math.MaxInt64, head} // 用dump 比较值要注意 有没有 0
	cur := dummy
	for cur != nil && cur.Next != nil { // IM : 条件
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

/*
测试用例: [0,0,0,1,2,2,3,4,4,4]
用dummy的话要问值范围
*/