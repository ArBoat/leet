func findKthLargest(nums []int, k int) int {
    return qucikSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func qucikSelect(nums []int, left, right, index int) int {
    p := partition(nums, left, right)
    if p == index {
        return nums[p]
    } else if p > index {
        return qucikSelect(nums, left, p-1, index)
    } else {
        return qucikSelect(nums, p+1, right, index)
    }
}

func partition(nums []int, left, right int) int {
    m:= left + (right -left)/2
    nums[m], nums[right] = nums[right], nums[m]
    i, j := left, left
    for ;i<right;i++{
        if nums[i]<nums[right]{
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    nums[j], nums[right] = nums[right], nums[j]
    return j
}
/*
O(n)/O（logn）
基于快排
*/