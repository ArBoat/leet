func exchange(nums []int) []int {
    i, j:=0, 0
    for i:=0;i<len(nums);i++{
        if nums[i]%2 == 1 {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    return nums
}