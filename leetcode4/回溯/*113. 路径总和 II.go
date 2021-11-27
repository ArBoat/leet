/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil { return nil}
    res:=[][]int{}
    path:=[]int{}
    var dfs func(node *TreeNode, sum int)
    dfs = func(node* TreeNode, sum int) {
        // IM: nil
        if node == nil {return}
        path = append(path, node.Val)
        defer func() { path = path[:len(path)-1] }()
        // IM : 终止条件
        if node.Left == nil && node.Right == nil && node.Val == sum{
            temp:=make([]int, len(path))
            copy(temp, path)
            res = append(res, temp)
            return
        }
        dfs(node.Left, sum-node.Val)
        dfs(node.Right, sum-node.Val)
    }
    dfs(root, targetSum)
    return res
}

/*
O(n2)/O(n)
回溯+dfs
defer妙用
*/