// l<r
func searchRange(nums []int, target int) []int {
	start := bianySearch(nums, target)
	// IM： 判断出界
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := bianySearch(nums, target+1) - 1
	return []int{start, end}
}

func bianySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

/*
O(logn)/O(1)
二分查找
*/