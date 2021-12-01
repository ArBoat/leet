func sortArrayByParityII(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]&1 == 1 {
			for nums[j]&1 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

// 喵喵喵啊