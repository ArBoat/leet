/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// inorder
// 节点小于2个的时候需要讨论
func getMinimumDifference(root *TreeNode) int {
	// 求最小赋最大
	res := math.MaxInt64
	var pre *TreeNode
	var travel func(cur *TreeNode)
	travel = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		travel(cur.Left)
		if pre != nil {
			res = min(res, cur.Val-pre.Val)
		}
		pre = cur
		travel(cur.Right)
	}
	travel(root)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
