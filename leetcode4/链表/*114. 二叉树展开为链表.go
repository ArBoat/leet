/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
     // 重要
    if root == nil {
        return
    }
    // no return, no :=
    flatten(root.Left)
    flatten(root.Right)
    l:=root.Left
    r:=root.Right
    // do not forget
    root.Left = nil
    root.Right = l
    p:=root
    for p.Right != nil {
        p = p.Right
    }
    p.Right = r
    return
}