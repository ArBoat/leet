func sortArray(nums []int) []int {
    qucikSelect(nums, 0, len(nums)-1)
    return nums
}

func qucikSelect(nums []int, left, right int) {
    if left < right {
        p := partition(nums, left, right)
        qucikSelect(nums, left, p-1)
        qucikSelect(nums, p+1, right)
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