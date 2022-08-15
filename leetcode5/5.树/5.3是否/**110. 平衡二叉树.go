/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//递归剪枝
// dfs匿名函数3注意：1.声明 2.实现 3.调用
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		if l == -1 {
			return -1
		}
		r := dfs(root.Right)
		if r == -1 {
			return -1
		}
		if abs(l, r) > 1 {
			res = false
			return -1
		}
		return max(l, r) + 1
	}
	//IM : 调用
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
