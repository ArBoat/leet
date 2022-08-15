func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, l, r int) int {
	n := r - l + 1
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[l]
	for i := 2; i <= n; i++ { // IM:  i:=2;i<=n;
		dp[i] = max(dp[i-2]+nums[i+l-1], dp[i-1]) // IM: nums[i+l-1]
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
环状
dp[i]
dp[i] = max(dp[i-2]+value, dp[i-1]) ; value:= nums[i+l-1]
dp[0] = 0,  dp[1] = nums[0]
*/