func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(a, b int) bool {
		return math.Abs(float64(nums[a])) > math.Abs(float64(nums[b]))
	})

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

/*
1.按绝对值从大到小排序
2.按顺序翻转负数
3. 奇数个是否反转最后一个
*/