/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func convertBST(root *TreeNode) *TreeNode {
    sum:=0
    // IM: 无需返回
    var dfs func(root *TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        dfs(root.Right)
        // IM
        sum += root.Val
        root.Val = sum
        dfs(root.Left)
    }
    dfs(root)
    return root
}

/* 
O(n)/O(n)
反中序遍历
*/