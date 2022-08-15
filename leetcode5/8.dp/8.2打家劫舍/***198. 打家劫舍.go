func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, nums[0] // IM: init
	for i := 2; i <= n; i++ { //IM: <=
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i]  下表为i的房屋最大金额
dp[i] = max(dp[i-2]+v, dp[i-1]) // 当前i，偷/不偷
dp[0] = 0, dp[1] = nums[0]
*/