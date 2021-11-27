/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) []int {
    if root == nil { return nil }
    res:=[]int{}
    q:=[]*TreeNode{root}
    for len(q)>0 {
        node:=q[0]
        q=q[1:]
        res = append(res, node.Val)
        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right !=nil {
            q = append(q, node.Right)
        }
    }
    return res
}