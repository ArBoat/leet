/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findBottomLeftValue(root *TreeNode) int {
    res:=[]int{}
    var dfs func(root *TreeNode, level int) 
    dfs = func(root *TreeNode, level int) {
        if root == nil { 
            return 
        }
        if len(res) == level {
            res = append(res, root.Val)
        }
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
    }
    dfs(root, 0)
    return res[len(res)-1]
}