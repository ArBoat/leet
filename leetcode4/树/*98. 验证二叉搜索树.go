/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt64, math.MaxInt64)
}

     // max, min
func dfs(root *TreeNode, min int, max int) bool {
    // IM: 
    if root == nil {
        return true
    }
    // 注意是<=
    if root.Val <= min || root.Val >= max {
        return false
    }
    return dfs(root.Left, min, root.Val) && 
    dfs(root.Right, root.Val, max)
}

/*
O(n)/O(n)
公式
*/