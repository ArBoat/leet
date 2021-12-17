/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{math.MaxInt32, head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Next != nil && cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

/*
测试用例: [0,0,0,1,2,2,3,4,4,4]
链表Next后有取值， 判断Next ！= nil
*/