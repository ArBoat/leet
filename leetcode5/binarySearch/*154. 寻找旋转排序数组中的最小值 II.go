func findMin(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        m := l+(r-l)/2 
        if nums[m] < nums[r] {
            r = m
        } else if nums[m] > nums[r] {
            l = m + 1
        } else {
			r--
		}
    }
    return nums[l]
}
// 有重复， 和右边相等直接r--