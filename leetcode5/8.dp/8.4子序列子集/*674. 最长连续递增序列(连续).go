func findLengthOfLCIS(nums []int) int {
	res := 1
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			dp[i+1] = dp[i] + 1
		}
		res = max(res, dp[i+1])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i+1] = dp[i]+1 : num[i]<nums[i+1]
      = 0
res = max(res, dp[i+1])
*/