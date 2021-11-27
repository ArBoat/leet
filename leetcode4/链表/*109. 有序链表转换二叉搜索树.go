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
// 1.[left,tight), 参数
func buildBST(left, right *ListNode) *TreeNode {
    // 2. IM: 边界
    if left == right {
        return nil
    }
    mid:=getMid(left, right)
    root:=&TreeNode{mid.Val,nil,nil}
    root.Left = buildBST(left, mid)
    root.Right = buildBST(mid.Next, right)
    return root
}
// 参数
func getMid(left, right *ListNode) *ListNode {
    fast, slow := left, left
    //3. IM: != right,  != right
    for fast != right  && fast.Next != right {
        slow = slow.Next
        // fast.Next.Next
        fast = fast.Next.Next
    }
    return slow
}