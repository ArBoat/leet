/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) < 1 { return nil }
    return build(nums, 0, len(nums)-1)
}

func build(nums []int, low, high int) *TreeNode {
    //终止条件
    if low > high {
        return nil
    }
    mid:= (low+high)/2
    node:= &TreeNode{Val:nums[mid]}
    node.Left = build(nums, low, mid-1)
    node.Right = build(nums, mid+1, high)
    return node
}

/*
O(n)/O(logn)
*/