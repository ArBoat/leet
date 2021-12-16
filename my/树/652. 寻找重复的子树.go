/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    res := []*TreeNode{}
    m:=make(map[string]int)
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string{
        if root == nil {
            return "#"
        }
        subStr:= strconv.Itoa(root.Val) + "," + dfs(root.Left) +","+dfs(root.Right)
        m[subStr]++
        if m[subStr] == 2 {
            res = append(res, root)
        }
        return subStr
    }
    dfs(root)
    return res
}

/*
O(n2)/O(n2)
*/