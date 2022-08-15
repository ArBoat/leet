/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// IM: 64 not 32
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	// IM : && not ||
	return dfs(root.Left, min, root.Val) &&
		dfs(root.Right, root.Val, max)
}

//树是否在范围内， 递归到左右子树