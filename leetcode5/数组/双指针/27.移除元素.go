func removeElement(nums []int, val int) int {
	right := len(nums) - 1
	left := 0
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}

//左右双指针， 优化

func removeElement(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ {
		// IM: 注意值
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

