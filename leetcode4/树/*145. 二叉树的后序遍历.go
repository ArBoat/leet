/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    res:=[]int{}
    stack:=[]*TreeNode{}
    var pre *TreeNode
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // IM: 条件
        if node.Right == nil || node.Right == pre {
            res = append(res, node.Val)
            pre = node
            node = nil
        } else { // else
            stack = append(stack, node)
            node = node.Right
        }
    }
    return res
}