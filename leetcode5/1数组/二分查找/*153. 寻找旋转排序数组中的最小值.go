// l < r
func findMin(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        m := l+(r-l)/2 
        //  == 无所谓
        if nums[m] < nums[r] {
            r = m
        } else {
            l = m + 1
        }
    }
    return nums[l]
}

// 问清 nums[i] != nums[i + 1]