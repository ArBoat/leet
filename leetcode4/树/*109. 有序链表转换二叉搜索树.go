/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedListToBST(head *ListNode) *TreeNode {
    return buildBST(head, nil)
}
// [left,right)
func buildBST(left, right *ListNode) *TreeNode {
    if left == right {
        return nil
    }
    mid:=getMid(left, right)
    root:=&TreeNode{mid.Val,nil,nil}
    root.Left = buildBST(left, mid)
    root.Right = buildBST(mid.Next,right)
    return root
}

func getMid(left, right *ListNode) *ListNode {
    slow, fast:= left, left
    for fast!= right && fast.Next!=right{
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}