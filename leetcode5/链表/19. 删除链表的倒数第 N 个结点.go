/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
     // Head
    dummy:=&ListNode{0,head}
    slow, fast:= dummy, head
    for i:=0;i<n;i++ {
        if fast == nil {
            return nil
        }
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

//添加dummy节点， 无需额外处理头节点
//双指针