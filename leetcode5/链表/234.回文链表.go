/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	res := true
	mid := midNode(head)
	p1 := head
	r := reverseList(mid.Next)
	p2 := r
	for p2 != nil {
		//    fmt.Println(p1.Val, p2.Val)
		if p1.Val != p2.Val {
			res = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	mid.Next = reverseList(r)
	return res
}

func midNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil { // 第一个中间点
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