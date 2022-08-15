func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxCur, minCur, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxTemp, minTemp := maxCur, minCur
		maxCur = max(nums[i], max(nums[i]*maxTemp, nums[i]*minTemp))
		minCur = min(nums[i], min(nums[i]*maxTemp, nums[i]*minTemp))
		res = max(res, maxCur)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
因为有负数，所以负数可能变成最大， 所以要记录max，min

max , min
不操作， 乘最大，乘最小

dp[i]
*/