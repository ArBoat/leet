	/**
	* Definition for a binary tree node.
	* type TreeNode struct {
	*     Val int
	*     Left *TreeNode
	*     Right *TreeNode
	* }
	 */
	// BFS
	func rightSideView(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		res := []int{}
		q := []*TreeNode{root}
		for len(q) > 0 {
			next := []*TreeNode{}
			res = append(res, q[len(q)-1].Val)
			for i := 0; i < len(q); i++ {
				node := q[i]
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

	// DFS
	// 中右左， 相等加入
	func rightSideView(root *TreeNode) []int {
		res := []int{}
		var dfs func(root *TreeNode, level int)
		dfs = func(root *TreeNode, level int) {
			if root == nil {
				return
			}
			if level == len(res) {
				res = append(res, root.Val)
			}
			// IM: 秒啊，先右再左
			dfs(root.Right, level+1)
			dfs(root.Left, level+1)
		}
		dfs(root, 0)
		return res
	}