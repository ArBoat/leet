/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil { return nil}
    if root.Val == key {
        // IM: return
        if root.Left == nil { return root.Right }
        if root.Right == nil { return root.Left }
        min:=getMinNode(root.Right)
        root.Val = min.Val
        //IM: 左侧和传参一致,root.Right,  min.Val
        root.Right = deleteNode(root.Right, min.Val)
    } else if root.Val > key {
        // root.Left 
        root.Left = deleteNode(root.Left, key)
    } else {
        // root.Right
        root.Right = deleteNode(root.Right, key)
    }
    return root
}

func getMinNode(node *TreeNode) * TreeNode {
    // IM: node.Left
    for node.Left != nil {
        node = node.Left
    }
    return node
}