/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthLargest(root *TreeNode, k int) int {
    res:=0
    count:=0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode){
        if root == nil || count > k {
            return
        }
        dfs(root.Right)
        count++
        if count == k {
            res = root.Val
            return
        }
        dfs(root.Left)
    }
    dfs(root)
    return res
}