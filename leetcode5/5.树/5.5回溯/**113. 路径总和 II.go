/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	path := []int{}
	var backtrack func(root *TreeNode, sum int)
	backtrack = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		//IM:defer
		defer func() {
			path = path[:len(path)-1]
		}()
		if root.Left == nil && root.Right == nil && root.Val == sum {
			//IM:copy slice
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		backtrack(root.Left, sum-root.Val)
		backtrack(root.Right, sum-root.Val)
	}
	backtrack(root, targetSum)
	return res
}

/*
O(n2)/O(n)

1.defer
2.copy
*/

