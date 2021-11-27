/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
    // MinInt32
    res:=math.MinInt32
    // dfs
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int{
        if root == nil { return 0}
        l:=max(0, dfs(root.Left))
        r:=max(0, dfs(root.Right))
        res=max(res, l+r+root.Val)
        return max(l, r) + root.Val
    }
    dfs(root)
    return res
}
func max(a, b int) int {
    if a>b {
         return a
    }
    return b
}