/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    if head.Val == val {
        return head.Next
    }
    cur:=head
    // 退出条件是下个节点为nil（尾部没找到），或者下个节点为目标值（中间找到了）
    for (cur.Next!=nil && cur.Next.Val != val){
        cur = cur.Next
    }
    // 判断
    if cur.Next != nil {
        cur.Next = cur.Next.Next
    }
    return head
}
// O（n）/O(1)