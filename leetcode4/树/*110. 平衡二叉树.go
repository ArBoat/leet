/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    res:=true
    if root == nil { return true }
    var dfs func(root *TreeNode) int 
    dfs = func(root *TreeNode) int {
        // 重要
        if root == nil { return 0}
        l:=dfs(root.Left)
        // 剪枝
        if l == -1 { return -1 }
        r:=dfs(root.Right)
        // 剪枝
        if r == - 1 { return -1}
        if math.Abs(float64(l-r)) > 1{
            res = false
            return -1
        }
        // 返回上一层
        return max(l, r)+1
    }
    // 别忘调用
    dfs(root)
    return res
}

func max(a, b int ) int {
    if a>b{
        return a
    }
    return b
}
/*
O(n)/O(n/)
剪枝
*/