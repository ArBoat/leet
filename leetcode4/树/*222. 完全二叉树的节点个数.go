/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func countNodes(root *TreeNode) int {
    if root == nil { return 0}
    l:=countHeight(root.Left)
    r:=countHeight(root.Right)
    // IM: 1<<
    if l == r {
        return ( 1<<l ) + countNodes(root.Right)
    } else {
        return ( 1<<r ) + countNodes(root.Left)
    }
}

func countHeight(root *TreeNode) int {
    res:=0
    // 重要
    for root != nil {
        res++
        root = root.Left
    }
    return res
}