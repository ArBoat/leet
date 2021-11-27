/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        // IM: nil insert
        return &TreeNode{Val:val}
    }
    if root.Val > val {
        root.Left = insertIntoBST(root.Left, val)
    } 
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
/*
O(n)/O(1)
查找是logn, 插入是O(1)
*/