/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}
	var min func(node *TreeNode) int
	min = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		left := min(node.Left)
		right := min(node.Right)

		if left == 2 && right == 2 { // (2,2) //左右都覆盖
			return 0
		} else if left == 0 || right == 0 { //(0,1),(0,2),(1,0),(2,0)(0,0)左右存在无覆盖
			res++
			return 1
		} else { //(1,1),(1,2),(2,1) //左右至少一个摄像头，且无无覆盖
			return 2
		}
	}
	if min(root) == 0 {
		res++
	}
	return res
}

/*
0：该节点无覆盖
1：本节点有摄像头
2：本节点有覆盖, 空节点也为2防止在叶子放
*/