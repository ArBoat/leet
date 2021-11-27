/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy:=&ListNode{0,nil}
    p,a,b:=dummy,l1,l2
    // IM: 先判断都非nil,  a , b
    for a != nil && b != nil {
        if a.Val < b.Val {
            p.Next = a
            a=a.Next
        } else {
            p.Next = b
            b=b.Next
        }
        p = p.Next
   }
   // IM: 不等于
   if a != nil {
       p.Next = a
   }
   if b != nil {
       p.Next = b
   }
   return dummy.Next
}