func search(nums []int, target int) int {
	start := bianrySearch(nums, target)
	if start >= len(nums) || nums[start] != target {
		return 0
	}
	end := bianrySearch(nums, target+1) - 1
	return end - start + 1
}

func bianrySearch(nums []int, target int) int {
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