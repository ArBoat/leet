/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	//IM: 开始传人
	q := []*TreeNode{root}
	// 层数，q长度
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		next := []*TreeNode{}
		// 遍历q
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)

			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return res
}

//DFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{})
		}
		// IM: 别忘加入
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	// IM: 别忘记调用
	dfs(root, 0)
	return res
}