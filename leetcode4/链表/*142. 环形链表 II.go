/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    slow, fast := head, head
    slow = slow.Next
    fast = fast.Next.Next
    for slow != fast {
        if fast == nil || fast.Next == nil {
            return nil
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    res := head
    for res != slow {
        res = res.Next
        slow = slow.Next
    }
    return res
}