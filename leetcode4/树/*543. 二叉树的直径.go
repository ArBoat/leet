/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *TreeNode) int {
    res:=0
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil { return 0 }
        l:=dfs(root.Left)
        r:=dfs(root.Right)
        // IM: 
        res = max(res, l+r)
        // 层数返回时+1
        return max(l, r) + 1
    }
    dfs(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
/*
O(n)/O(height)
*/