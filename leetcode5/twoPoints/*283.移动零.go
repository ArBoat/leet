func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	first, second := 0, 0
	for ; first < len(nums); first++ {
		if nums[first] != 0 {
			if first > second {
				nums[second] = nums[first]
				nums[first] = 0
			}
			second++
		}
	}
}

// 双指针