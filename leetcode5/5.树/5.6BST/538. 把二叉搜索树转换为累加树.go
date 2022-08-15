/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//反中序遍历， 右中左
func convertBST(root *TreeNode) *TreeNode {
	//IM: sum
	sum := 0
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Right)
		sum += root.Val
		root.Val = sum
		travel(root.Left)
	}
	travel(root)
	return root
}
