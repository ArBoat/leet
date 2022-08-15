/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		// 左子树也都小于low
		//IM: 是赋值节点
		root = trimBST(root.Right, low, high)
	} else if root.Val > high {
		// 右子树也都大于hight
		//IM: 是赋值节点
		root = trimBST(root.Left, low, high)
	} else {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
	}
	return root
}
