/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 新值和原值都不同
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		// im 赋值
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		// im 赋值
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
