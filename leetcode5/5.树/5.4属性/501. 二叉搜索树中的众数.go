/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	count, max := 0, 0
	var pre *TreeNode
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		// if pre == nil {
		// 	count = 1
		// 	max = 1
		// 	res = []int{root.Val}
		// } else if pre.Val == root.Val {
		// 	count++
		// 	if count == max {
		// 		res = append(res, root.Val)
		// 	} else if count > max {
		// 		max = count
		// 		res = []int{root.Val}
		// 	}
		// } else {
		// 	count = 1
		// 	if count == max {
		// 		res = append(res, root.Val)
		// 	} else if count > max {
		// 		max = count
		// 		res = []int{root.Val}
		// 	}
		// }
		// 先处理count， 再根据count和max关系处理res和max，然后赋值pre
		if pre != nil && pre.Val == root.Val {
			count++
		} else {
			count = 1
		}
		if count > max {
			res = []int{root.Val}
			max = count
		} else if count == max {
			res = append(res, root.Val)
		}
		pre = root
		travel(root.Right)
	}
	travel(root)
	return res
}
