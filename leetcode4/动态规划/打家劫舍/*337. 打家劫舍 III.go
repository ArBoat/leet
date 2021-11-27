/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rob(root *TreeNode) int {
    res:=dp(root)
    return max(res[0], res[1])
}

func dp(root *TreeNode)[]int {
    if root == nil {
        // IM: 返回两个0
        return []int{0,0}
    }
    l:=dp(root.Left)
    r:=dp(root.Right)
    // IM: 公式, 左右子树相加
    no:= max(l[0], l[1]) + max(r[0], r[1])
    yes:= root.Val + l[0] + r[0]
    return []int{no, yes}
} 

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
res[] 两个元素， res[0]表示不枪root， res[1]表示强root
no = max(left[0],left[1]) + max(right[0], right[1])
yes = root.Val + left[0] + right[0]
*/