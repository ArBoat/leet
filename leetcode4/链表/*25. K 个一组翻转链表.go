/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    var res *ListNode
    if head == nil || k < 1 { return nil }s
    // IM: two points
    a, b := head, head
    for i:=0; i<k; i++ {
        if b == nil {
            // 不足结束
            return head
        }
        b=b.Next
    }
    res = resverse(a, b)
    //重要
    a.Next = reverseKGroup(b, k)
    return res
}

// reverse [a,b)
func resverse (a, b *ListNode) *ListNode {
    var pre *ListNode
    for a != b {
        //  params
        pre, a ,a.Next = a, a.Next, pre
    }
    return pre
}