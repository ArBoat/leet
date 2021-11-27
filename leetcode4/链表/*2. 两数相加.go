/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy:=&ListNode{0,nil}
    cur:=dummy
    carry:=0
    //IM :   这里是for
    for l1!=nil || l2 != nil {
        if l1 != nil {
            carry+=l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            carry+=l2.Val
            l2 = l2.Next
        }
        cur.Next = &ListNode{carry%10,nil}
        carry = carry/10
        cur = cur.Next
    }
    // 注意
    if carry == 1 {
        cur.Next= &ListNode{1,nil}
    }
    return dummy.Next
}

// dummy 指针
// 注意carry