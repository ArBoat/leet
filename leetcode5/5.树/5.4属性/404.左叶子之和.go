/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
 难点在于判断左叶子节点
 通过父节点判断：左子节点存在，且左子节点无左右子节点
*/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	v := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		v += root.Left.Val
	}
	return v + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
