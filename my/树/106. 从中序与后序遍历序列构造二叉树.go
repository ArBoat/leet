/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    n:= len(inorder)
    if n == 0 {
        return nil
    }
    root := &TreeNode{postorder[n-1],nil,nil}
    i:=0
    for ;i<n;i++{
        if inorder[i] == root.Val {
            break
        }
    }
    root.Left = buildTree(inorder[:i],postorder[:i])
    root.Right = buildTree(inorder[i+1:],postorder[i:n-1])
    return root
}