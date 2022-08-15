/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//BFS
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	// 传入
	q := []*TreeNode{root}
	for level := 0; len(q) > 0; level++ {
		next := []*TreeNode{}
		// IM: 初始
		res = append([][]int{[]int{}}, res...)
		for i := 0; i < len(q); i++ {
			res[0] = append(res[0], q[i].Val)
			if q[i].Left != nil {
				next = append(next, q[i].Left)
			}
			if q[i].Right != nil {
				next = append(next, q[i].Right)
			}
		}
		q = next
	}
	return res
}

// DFS
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(res) {
			res = append([][]int{[]int{}}, res...)
		}
		// 没层level不同
		index := len(res) - 1 - level
		res[index] = append(res[index], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}