/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
     // IM: 无需画蛇添足
    // if head == nil { return nil}
    dummy:=&ListNode{0,head}
    cur := dummy
    for cur != nil && cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {
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