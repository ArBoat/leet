func searchRange(nums []int, target int) []int {
	start := binarySearch(nums, target)
	// start >= len(nums) start返回的值在数组范围内， 边界问题，数组访问要判断
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := binarySearch(nums, target+1) - 1
	return []int{start, end}
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}