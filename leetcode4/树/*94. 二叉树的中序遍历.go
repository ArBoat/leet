/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    res:=[]int{}
    stack:=[]*TreeNode{}
    node := root
    for node != nil || len(stack)>0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, node.Val)
        node = node.Right
    }
    return res
}