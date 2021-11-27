/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func recoverTree(root *TreeNode)  {
    nodes:=[]*TreeNode{}
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        dfs(root.Left)
        // IM: nodes *TreeNode
        nodes = append(nodes, root)
        dfs(root.Right)
    }
    dfs(root)
    x, y:=-1,-1
    // IM: <
    for i:=0;i<len(nodes)-1;i++{
        if nodes[i].Val > nodes[i+1].Val {
            y = i+1    // 第二个位置
            if x == -1 {
                x = i  // 第一个
            }
        }
    }
    nodes[x].Val, nodes[y].Val = nodes[y].Val, nodes[x].Val
}

/*
先中序遍历到数组
再找到错误位置
交换
*/