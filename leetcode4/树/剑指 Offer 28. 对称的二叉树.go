/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    var dfs func(p, q *TreeNode) bool
    dfs = func(p, q *TreeNode) bool {
        if p == nil && q == nil { 
            return true
        } 
        if p == nil || q == nil {
            return false
        }
        // 1, 2, 3
        return p.Val == q.Val && dfs(p.Left, q.Right) && 
        dfs(p.Right, q.Left)
    }
    return dfs(root, root)
}